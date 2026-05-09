package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Manikeshmk/silent-meeting-summarizer/internal/ai"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/config"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/handler"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/logger"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/middleware"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/repository"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize logger
	appLogger := logger.NewLogger(cfg.Environment)
	defer appLogger.Sync()

	appLogger.Info("starting Silent Meeting Summarizer",
		"environment", cfg.Environment,
		"server_port", cfg.ServerPort,
	)

	// Initialize database
	db, err := initializeDatabase(cfg, appLogger)
	if err != nil {
		appLogger.Error("database initialization failed", err)
		os.Exit(1)
	}
	defer db.Close()

	// Initialize repositories
	meetingRepo := repository.NewMeetingRepository(db)
	speakerRepo := repository.NewSpeakerRepository(db)
	transcriptRepo := repository.NewTranscriptRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, appLogger, cfg.JWTSecret)
	meetingService := service.NewMeetingService(meetingRepo, speakerRepo, transcriptRepo, appLogger)
	taskService := service.NewTaskService(taskRepo, appLogger)

	// Initialize AI services
	summarizationService := ai.NewSummarizationService(appLogger)
	taskExtractionService := ai.NewTaskExtractionService(appLogger)
	conflictDetectionService := ai.NewConflictDetectionService(appLogger)
	confusionDetectionService := ai.NewConfusionDetectionService(appLogger)
	sentimentService := ai.NewSentimentService(appLogger)
	decisionService := ai.NewDecisionService(appLogger)
	_ = summarizationService
	_ = taskExtractionService
	_ = conflictDetectionService
	_ = confusionDetectionService
	_ = sentimentService
	_ = decisionService

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService, appLogger)
	meetingHandler := handler.NewMeetingHandler(meetingService, appLogger)
	taskHandler := handler.NewTaskHandler(taskService, appLogger)

	// Setup Gin router
	router := setupRouter(authHandler, meetingHandler, taskHandler, authService, appLogger)

	// Create HTTP server
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server in a goroutine
	go func() {
		appLogger.Info("server starting", "address", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Error("server error", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.Error("server forced to shutdown", err)
	}

	appLogger.Info("server stopped")
}

func initializeDatabase(cfg *config.Config, appLogger *logger.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxConnections)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	appLogger.Info("database connected successfully")

	// Run migrations
	if err := runMigrations(db, appLogger); err != nil {
		appLogger.Warn("migrations failed", "error", err.Error())
	}

	return db, nil
}

func runMigrations(db *sqlx.DB, appLogger *logger.Logger) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) DEFAULT 'user',
			active BOOLEAN DEFAULT true,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS meetings (
			id VARCHAR(36) PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			start_time TIMESTAMP NOT NULL,
			end_time TIMESTAMP,
			status VARCHAR(50) DEFAULT 'active',
			duration INTEGER,
			participant_count INTEGER DEFAULT 0,
			user_id VARCHAR(36) NOT NULL REFERENCES users(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS speakers (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255),
			embedding TEXT,
			speak_time INTEGER DEFAULT 0,
			turn_count INTEGER DEFAULT 0,
			interruptions INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS transcripts (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			speaker_id VARCHAR(36) REFERENCES speakers(id),
			speaker_name VARCHAR(255),
			text TEXT NOT NULL,
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			confidence DECIMAL(4,2),
			language VARCHAR(10) DEFAULT 'en',
			sentiment VARCHAR(50),
			emotions TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS tasks (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			title VARCHAR(255) NOT NULL,
			description TEXT,
			owner VARCHAR(255),
			status VARCHAR(50) DEFAULT 'pending',
			priority VARCHAR(50) DEFAULT 'medium',
			deadline TIMESTAMP,
			extracted_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS summaries (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			summary_type VARCHAR(50),
			content TEXT,
			key_points TEXT,
			generated_at TIMESTAMP,
			model VARCHAR(100),
			tokens INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS decisions (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			title VARCHAR(255),
			description TEXT,
			confidence DECIMAL(4,2),
			resolved BOOLEAN DEFAULT false,
			participant_count INTEGER,
			extracted_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS conflicts (
			id VARCHAR(36) PRIMARY KEY,
			meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
			description TEXT,
			conflict_score DECIMAL(4,2),
			unresolved_topics TEXT,
			participants TEXT,
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			status VARCHAR(50),
			detected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_meetings_user_id ON meetings(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_meetings_status ON meetings(status)`,
		`CREATE INDEX IF NOT EXISTS idx_transcripts_meeting_id ON transcripts(meeting_id)`,
		`CREATE INDEX IF NOT EXISTS idx_speakers_meeting_id ON speakers(meeting_id)`,
		`CREATE INDEX IF NOT EXISTS idx_tasks_meeting_id ON tasks(meeting_id)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			appLogger.Warn("migration failed", "error", err.Error())
		}
	}

	appLogger.Info("migrations completed")
	return nil
}

func setupRouter(
	authHandler *handler.AuthHandler,
	meetingHandler *handler.MeetingHandler,
	taskHandler *handler.TaskHandler,
	authService *service.AuthService,
	appLogger *logger.Logger,
) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(middleware.LoggingMiddleware(appLogger))
	router.Use(middleware.ErrorHandlerMiddleware(appLogger))
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Auth routes (public)
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected routes
	protectedRoutes := router.Group("")
	protectedRoutes.Use(middleware.AuthMiddleware(authService, appLogger))
	{
		// Meeting routes
		meetingsGroup := protectedRoutes.Group("/meetings")
		{
			meetingsGroup.POST("", meetingHandler.CreateMeeting)
			meetingsGroup.GET("", meetingHandler.ListMeetings)
			meetingsGroup.GET("/:id", meetingHandler.GetMeeting)
			meetingsGroup.POST("/:id/end", meetingHandler.EndMeeting)

			// Task routes
			meetingsGroup.GET("/:meetingId/tasks", taskHandler.GetMeetingTasks)
		}
	}

	return router
}
