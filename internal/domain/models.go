package domain

import (
	"time"

	"github.com/google/uuid"
)

// Meeting represents a meeting session
type Meeting struct {
	ID            string    `db:"id" json:"id"`
	Title         string    `db:"title" json:"title"`
	Description   string    `db:"description" json:"description"`
	StartTime     time.Time `db:"start_time" json:"start_time"`
	EndTime       *time.Time `db:"end_time" json:"end_time,omitempty"`
	Status        string    `db:"status" json:"status"` // active, completed, cancelled
	Duration      int       `db:"duration" json:"duration"`
	ParticipantCount int    `db:"participant_count" json:"participant_count"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	UserID        string    `db:"user_id" json:"user_id"`
}

// Transcript represents the transcript of a meeting
type Transcript struct {
	ID           string    `db:"id" json:"id"`
	MeetingID    string    `db:"meeting_id" json:"meeting_id"`
	SpeakerID    string    `db:"speaker_id" json:"speaker_id"`
	SpeakerName  string    `db:"speaker_name" json:"speaker_name"`
	Text         string    `db:"text" json:"text"`
	StartTime    time.Time `db:"start_time" json:"start_time"`
	EndTime      time.Time `db:"end_time" json:"end_time"`
	Confidence   float64   `db:"confidence" json:"confidence"`
	Language     string    `db:"language" json:"language"`
	Sentiment    string    `db:"sentiment" json:"sentiment"` // positive, neutral, negative
	Emotions     string    `db:"emotions" json:"emotions"`   // JSON string of emotion scores
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

// Speaker represents a unique speaker in the meeting
type Speaker struct {
	ID          string    `db:"id" json:"id"`
	MeetingID   string    `db:"meeting_id" json:"meeting_id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	Embedding   string    `db:"embedding" json:"embedding"` // JSON array of embeddings
	SpeakTime   int       `db:"speak_time" json:"speak_time"` // in seconds
	TurnCount   int       `db:"turn_count" json:"turn_count"`
	Interruptions int     `db:"interruptions" json:"interruptions"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// Summary represents AI-generated summaries
type Summary struct {
	ID              string    `db:"id" json:"id"`
	MeetingID       string    `db:"meeting_id" json:"meeting_id"`
	SummaryType     string    `db:"summary_type" json:"summary_type"` // concise, detailed, executive
	Content         string    `db:"content" json:"content"`
	KeyPoints       string    `db:"key_points" json:"key_points"`      // JSON array
	GeneratedAt     time.Time `db:"generated_at" json:"generated_at"`
	Model           string    `db:"model" json:"model"`
	Tokens          int       `db:"tokens" json:"tokens"`
}

// Task represents extracted tasks/action items
type Task struct {
	ID          string    `db:"id" json:"id"`
	MeetingID   string    `db:"meeting_id" json:"meeting_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Owner       string    `db:"owner" json:"owner"`
	Status      string    `db:"status" json:"status"` // pending, in_progress, completed
	Priority    string    `db:"priority" json:"priority"` // low, medium, high
	Deadline    *time.Time `db:"deadline" json:"deadline,omitempty"`
	ExtractedAt time.Time `db:"extracted_at" json:"extracted_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// Decision represents meeting decisions
type Decision struct {
	ID            string    `db:"id" json:"id"`
	MeetingID     string    `db:"meeting_id" json:"meeting_id"`
	Title         string    `db:"title" json:"title"`
	Description   string    `db:"description" json:"description"`
	Confidence    float64   `db:"confidence" json:"confidence"`
	Resolved      bool      `db:"resolved" json:"resolved"`
	ParticipantCount int    `db:"participant_count" json:"participant_count"`
	ExtractedAt   time.Time `db:"extracted_at" json:"extracted_at"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

// Conflict represents detected conflicts/disagreements
type Conflict struct {
	ID              string    `db:"id" json:"id"`
	MeetingID       string    `db:"meeting_id" json:"meeting_id"`
	Description     string    `db:"description" json:"description"`
	ConflictScore   float64   `db:"conflict_score" json:"conflict_score"`
	UnresolvedTopics string   `db:"unresolved_topics" json:"unresolved_topics"` // JSON array
	Participants    string    `db:"participants" json:"participants"` // JSON array
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         *time.Time `db:"end_time" json:"end_time,omitempty"`
	Status          string    `db:"status" json:"status"` // unresolved, resolved, acknowledged
	DetectedAt      time.Time `db:"detected_at" json:"detected_at"`
}

// Confusion represents confusion detection data
type Confusion struct {
	ID              string    `db:"id" json:"id"`
	MeetingID       string    `db:"meeting_id" json:"meeting_id"`
	Participant     string    `db:"participant" json:"participant"`
	Description     string    `db:"description" json:"description"`
	ClarificationCount int    `db:"clarification_count" json:"clarification_count"`
	HesitationMarkers string  `db:"hesitation_markers" json:"hesitation_markers"` // JSON array
	Heatmap         string    `db:"heatmap" json:"heatmap"`  // JSON object
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         *time.Time `db:"end_time" json:"end_time,omitempty"`
	DetectedAt      time.Time `db:"detected_at" json:"detected_at"`
}

// RepeatedTopic represents repeated topics in meetings
type RepeatedTopic struct {
	ID              string    `db:"id" json:"id"`
	MeetingID       string    `db:"meeting_id" json:"meeting_id"`
	Topic           string    `db:"topic" json:"topic"`
	Frequency       int       `db:"frequency" json:"frequency"`
	FirstOccurrence time.Time `db:"first_occurrence" json:"first_occurrence"`
	LastOccurrence  time.Time `db:"last_occurrence" json:"last_occurrence"`
	WastedTime      int       `db:"wasted_time" json:"wasted_time"` // in seconds
	Instances       string    `db:"instances" json:"instances"` // JSON array of time ranges
	DetectedAt      time.Time `db:"detected_at" json:"detected_at"`
}

// AudioChunk represents a chunk of audio data
type AudioChunk struct {
	ID        string
	Data      []byte
	Timestamp time.Time
	Duration  int // milliseconds
	MeetingID string
}

// TranscriptChunk represents a chunk of transcribed text
type TranscriptChunk struct {
	ID         string
	MeetingID  string
	SpeakerID  string
	Text       string
	Confidence float64
	StartTime  time.Time
	EndTime    time.Time
	IsFinal    bool
}

// User represents a system user
type User struct {
	ID        string    `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"-"`
	Role      string    `db:"role" json:"role"` // admin, user, viewer
	Active    bool      `db:"active" json:"active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// NewMeeting creates a new meeting
func NewMeeting(title, description, userID string) *Meeting {
	return &Meeting{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Status:      "active",
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		StartTime:   time.Now(),
	}
}

// NewUser creates a new user
func NewUser(email, name, password string) *User {
	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		Name:      name,
		Password:  password,
		Role:      "user",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
