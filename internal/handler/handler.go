package handler

import (
	"net/http"

	"github.com/Manikeshmk/silent-meeting-summarizer/internal/logger"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/service"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	authService *service.AuthService
	logger      *logger.Logger
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *service.AuthService, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration details"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.authService.Register(c.Request.Context(), req.Email, req.Name, req.Password)
	if err != nil {
		h.logger.Error("registration failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate a user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Error("login failed", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// MeetingHandler handles meeting endpoints
type MeetingHandler struct {
	meetingService *service.MeetingService
	logger         *logger.Logger
}

// NewMeetingHandler creates a new meeting handler
func NewMeetingHandler(meetingService *service.MeetingService, logger *logger.Logger) *MeetingHandler {
	return &MeetingHandler{
		meetingService: meetingService,
		logger:         logger,
	}
}

// CreateMeeting godoc
// @Summary Create a new meeting
// @Description Start a new meeting session
// @Tags meetings
// @Accept json
// @Produce json
// @Param request body CreateMeetingRequest true "Meeting details"
// @Success 201 {object} MeetingResponse
// @Failure 400 {object} ErrorResponse
// @Router /meetings [post]
func (h *MeetingHandler) CreateMeeting(c *gin.Context) {
	var req CreateMeetingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	meeting, err := h.meetingService.CreateMeeting(c.Request.Context(), req.Title, req.Description, userID.(string))
	if err != nil {
		h.logger.Error("failed to create meeting", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        meeting.ID,
		"title":     meeting.Title,
		"status":    meeting.Status,
		"startTime": meeting.StartTime,
	})
}

// GetMeeting godoc
// @Summary Get meeting details
// @Description Retrieve details of a specific meeting
// @Tags meetings
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {object} MeetingResponse
// @Failure 404 {object} ErrorResponse
// @Router /meetings/{id} [get]
func (h *MeetingHandler) GetMeeting(c *gin.Context) {
	id := c.Param("id")

	meeting, err := h.meetingService.GetMeeting(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("failed to get meeting", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "meeting not found"})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

// ListMeetings godoc
// @Summary List user meetings
// @Description Get all meetings for the authenticated user
// @Tags meetings
// @Produce json
// @Success 200 {array} MeetingResponse
// @Router /meetings [get]
func (h *MeetingHandler) ListMeetings(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	meetings, err := h.meetingService.ListMeetings(c.Request.Context(), userID.(string), 50, 0)
	if err != nil {
		h.logger.Error("failed to list meetings", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meetings)
}

// EndMeeting godoc
// @Summary End a meeting
// @Description Mark a meeting as completed
// @Tags meetings
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {object} MeetingResponse
// @Router /meetings/{id}/end [post]
func (h *MeetingHandler) EndMeeting(c *gin.Context) {
	id := c.Param("id")

	err := h.meetingService.EndMeeting(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("failed to end meeting", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "meeting ended"})
}

// TaskHandler handles task endpoints
type TaskHandler struct {
	taskService *service.TaskService
	logger      *logger.Logger
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(taskService *service.TaskService, logger *logger.Logger) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		logger:      logger,
	}
}

// GetMeetingTasks godoc
// @Summary Get meeting tasks
// @Description Get all tasks extracted from a meeting
// @Tags tasks
// @Produce json
// @Param meetingId path string true "Meeting ID"
// @Success 200 {array} TaskResponse
// @Router /meetings/{meetingId}/tasks [get]
func (h *TaskHandler) GetMeetingTasks(c *gin.Context) {
	meetingID := c.Param("meetingId")

	tasks, err := h.taskService.GetMeetingTasks(c.Request.Context(), meetingID)
	if err != nil {
		h.logger.Error("failed to get tasks", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Request/Response DTOs

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateMeetingRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type MeetingResponse struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	StartTime       interface{} `json:"startTime"`
	EndTime         interface{} `json:"endTime"`
	ParticipantCount int    `json:"participantCount"`
	Duration        int    `json:"duration"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Owner       string `json:"owner"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
