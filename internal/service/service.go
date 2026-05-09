package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/Manikeshmk/silent-meeting-summarizer/internal/domain"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/logger"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AuthService handles authentication and authorization
type AuthService struct {
	userRepo *repository.UserRepository
	logger   *logger.Logger
	jwtSecret string
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo *repository.UserRepository, logger *logger.Logger, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		logger:    logger,
		jwtSecret: jwtSecret,
	}
}

// Register creates a new user account
func (s *AuthService) Register(ctx context.Context, email, name, password string) (*domain.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword := s.hashPassword(password)

	user := &domain.User{
		ID:       uuid.New().String(),
		Email:    email,
		Name:     name,
		Password: hashedPassword,
		Role:     "user",
		Active:   true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.Error("failed to create user", err)
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !s.verifyPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken validates a JWT token
func (s *AuthService) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func (s *AuthService) hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (s *AuthService) verifyPassword(password, hashedPassword string) bool {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]) == hashedPassword
}

// MeetingService handles meeting operations
type MeetingService struct {
	meetingRepo   *repository.MeetingRepository
	speakerRepo   *repository.SpeakerRepository
	transcriptRepo *repository.TranscriptRepository
	logger        *logger.Logger
}

// NewMeetingService creates a new meeting service
func NewMeetingService(
	meetingRepo *repository.MeetingRepository,
	speakerRepo *repository.SpeakerRepository,
	transcriptRepo *repository.TranscriptRepository,
	logger *logger.Logger,
) *MeetingService {
	return &MeetingService{
		meetingRepo:    meetingRepo,
		speakerRepo:    speakerRepo,
		transcriptRepo: transcriptRepo,
		logger:         logger,
	}
}

// CreateMeeting creates a new meeting
func (s *MeetingService) CreateMeeting(ctx context.Context, title, description, userID string) (*domain.Meeting, error) {
	meeting := domain.NewMeeting(title, description, userID)
	
	if err := s.meetingRepo.Create(ctx, meeting); err != nil {
		s.logger.Error("failed to create meeting", err)
		return nil, err
	}

	return meeting, nil
}

// GetMeeting retrieves a meeting by ID
func (s *MeetingService) GetMeeting(ctx context.Context, id string) (*domain.Meeting, error) {
	return s.meetingRepo.GetByID(ctx, id)
}

// EndMeeting marks a meeting as completed
func (s *MeetingService) EndMeeting(ctx context.Context, id string) error {
	meeting, err := s.meetingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	meeting.Status = "completed"
	return s.meetingRepo.Update(ctx, meeting)
}

// ListMeetings retrieves all meetings for a user
func (s *MeetingService) ListMeetings(ctx context.Context, userID string, limit, offset int) ([]domain.Meeting, error) {
	return s.meetingRepo.List(ctx, userID, limit, offset)
}

// TaskService handles task operations
type TaskService struct {
	taskRepo *repository.TaskRepository
	logger   *logger.Logger
}

// NewTaskService creates a new task service
func NewTaskService(taskRepo *repository.TaskRepository, logger *logger.Logger) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
		logger:   logger,
	}
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(ctx context.Context, meetingID, title, description, owner, priority string) (*domain.Task, error) {
	task := &domain.Task{
		ID:          uuid.New().String(),
		MeetingID:   meetingID,
		Title:       title,
		Description: description,
		Owner:       owner,
		Status:      "pending",
		Priority:    priority,
	}

	if err := s.taskRepo.Create(ctx, task); err != nil {
		s.logger.Error("failed to create task", err)
		return nil, err
	}

	return task, nil
}

// GetMeetingTasks retrieves all tasks for a meeting
func (s *TaskService) GetMeetingTasks(ctx context.Context, meetingID string) ([]domain.Task, error) {
	return s.taskRepo.GetByMeetingID(ctx, meetingID)
}

// UpdateTask updates a task
func (s *TaskService) UpdateTask(ctx context.Context, task *domain.Task) error {
	return s.taskRepo.Update(ctx, task)
}
