package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Manikeshmk/silent-meeting-summarizer/internal/domain"
	"github.com/jmoiron/sqlx"
)

// MeetingRepository handles meeting database operations
type MeetingRepository struct {
	db *sqlx.DB
}

// NewMeetingRepository creates a new meeting repository
func NewMeetingRepository(db *sqlx.DB) *MeetingRepository {
	return &MeetingRepository{db}
}

// Create inserts a new meeting
func (r *MeetingRepository) Create(ctx context.Context, meeting *domain.Meeting) error {
	query := `
		INSERT INTO meetings (id, title, description, start_time, status, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		meeting.ID, meeting.Title, meeting.Description, meeting.StartTime,
		meeting.Status, meeting.UserID, meeting.CreatedAt, meeting.UpdatedAt)
	return err
}

// GetByID retrieves a meeting by ID
func (r *MeetingRepository) GetByID(ctx context.Context, id string) (*domain.Meeting, error) {
	meeting := &domain.Meeting{}
	query := `SELECT * FROM meetings WHERE id = $1`
	err := r.db.GetContext(ctx, meeting, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("meeting not found")
	}
	return meeting, err
}

// Update updates a meeting
func (r *MeetingRepository) Update(ctx context.Context, meeting *domain.Meeting) error {
	query := `
		UPDATE meetings SET 
			title = $1, description = $2, status = $3, duration = $4,
			participant_count = $5, end_time = $6, updated_at = $7
		WHERE id = $8
	`
	meeting.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, query,
		meeting.Title, meeting.Description, meeting.Status, meeting.Duration,
		meeting.ParticipantCount, meeting.EndTime, meeting.UpdatedAt, meeting.ID)
	return err
}

// List retrieves all meetings for a user
func (r *MeetingRepository) List(ctx context.Context, userID string, limit, offset int) ([]domain.Meeting, error) {
	var meetings []domain.Meeting
	query := `
		SELECT * FROM meetings WHERE user_id = $1 
		ORDER BY start_time DESC LIMIT $2 OFFSET $3
	`
	err := r.db.SelectContext(ctx, &meetings, query, userID, limit, offset)
	return meetings, err
}

// Delete deletes a meeting
func (r *MeetingRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM meetings WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// TranscriptRepository handles transcript database operations
type TranscriptRepository struct {
	db *sqlx.DB
}

// NewTranscriptRepository creates a new transcript repository
func NewTranscriptRepository(db *sqlx.DB) *TranscriptRepository {
	return &TranscriptRepository{db}
}

// Create inserts a new transcript
func (r *TranscriptRepository) Create(ctx context.Context, transcript *domain.Transcript) error {
	query := `
		INSERT INTO transcripts (id, meeting_id, speaker_id, speaker_name, text, start_time, 
			end_time, confidence, language, sentiment, emotions, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := r.db.ExecContext(ctx, query,
		transcript.ID, transcript.MeetingID, transcript.SpeakerID, transcript.SpeakerName,
		transcript.Text, transcript.StartTime, transcript.EndTime, transcript.Confidence,
		transcript.Language, transcript.Sentiment, transcript.Emotions, transcript.CreatedAt)
	return err
}

// GetByMeetingID retrieves transcripts for a meeting
func (r *TranscriptRepository) GetByMeetingID(ctx context.Context, meetingID string) ([]domain.Transcript, error) {
	var transcripts []domain.Transcript
	query := `SELECT * FROM transcripts WHERE meeting_id = $1 ORDER BY start_time ASC`
	err := r.db.SelectContext(ctx, &transcripts, query, meetingID)
	return transcripts, err
}

// SpeakerRepository handles speaker database operations
type SpeakerRepository struct {
	db *sqlx.DB
}

// NewSpeakerRepository creates a new speaker repository
func NewSpeakerRepository(db *sqlx.DB) *SpeakerRepository {
	return &SpeakerRepository{db}
}

// Create inserts a new speaker
func (r *SpeakerRepository) Create(ctx context.Context, speaker *domain.Speaker) error {
	query := `
		INSERT INTO speakers (id, meeting_id, name, email, embedding, speak_time, turn_count, interruptions, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		speaker.ID, speaker.MeetingID, speaker.Name, speaker.Email, speaker.Embedding,
		speaker.SpeakTime, speaker.TurnCount, speaker.Interruptions, speaker.CreatedAt)
	return err
}

// GetByMeetingID retrieves speakers for a meeting
func (r *SpeakerRepository) GetByMeetingID(ctx context.Context, meetingID string) ([]domain.Speaker, error) {
	var speakers []domain.Speaker
	query := `SELECT * FROM speakers WHERE meeting_id = $1 ORDER BY speak_time DESC`
	err := r.db.SelectContext(ctx, &speakers, query, meetingID)
	return speakers, err
}

// Update updates a speaker
func (r *SpeakerRepository) Update(ctx context.Context, speaker *domain.Speaker) error {
	query := `
		UPDATE speakers SET 
			speak_time = $1, turn_count = $2, interruptions = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query,
		speaker.SpeakTime, speaker.TurnCount, speaker.Interruptions, speaker.ID)
	return err
}

// TaskRepository handles task database operations
type TaskRepository struct {
	db *sqlx.DB
}

// NewTaskRepository creates a new task repository
func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db}
}

// Create inserts a new task
func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := `
		INSERT INTO tasks (id, meeting_id, title, description, owner, status, priority, deadline, extracted_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		task.ID, task.MeetingID, task.Title, task.Description, task.Owner,
		task.Status, task.Priority, task.Deadline, task.ExtractedAt, task.CreatedAt)
	return err
}

// GetByMeetingID retrieves tasks for a meeting
func (r *TaskRepository) GetByMeetingID(ctx context.Context, meetingID string) ([]domain.Task, error) {
	var tasks []domain.Task
	query := `SELECT * FROM tasks WHERE meeting_id = $1 ORDER BY priority DESC, deadline ASC`
	err := r.db.SelectContext(ctx, &tasks, query, meetingID)
	return tasks, err
}

// Update updates a task
func (r *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	query := `
		UPDATE tasks SET 
			status = $1, priority = $2, deadline = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query,
		task.Status, task.Priority, task.Deadline, task.ID)
	return err
}

// UserRepository handles user database operations
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

// Create inserts a new user
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, email, name, password, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.Name, user.Password, user.Role, user.Active,
		user.CreatedAt, user.UpdatedAt)
	return err
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, user, query, email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, user, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}
