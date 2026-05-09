package ai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Manikeshmk/silent-meeting-summarizer/internal/domain"
	"github.com/Manikeshmk/silent-meeting-summarizer/internal/logger"
)

// SpeechProvider is an interface for different speech-to-text providers
type SpeechProvider interface {
	Transcribe(ctx context.Context, audioData []byte) (*domain.TranscriptChunk, error)
	Health(ctx context.Context) error
}

// TranscriptionService handles speech-to-text operations
type TranscriptionService struct {
	providers map[string]SpeechProvider
	primary   string
	fallback  string
	logger    *logger.Logger
}

// NewTranscriptionService creates a new transcription service
func NewTranscriptionService(logger *logger.Logger) *TranscriptionService {
	providers := make(map[string]SpeechProvider)
	
	return &TranscriptionService{
		providers: providers,
		logger:    logger,
	}
}

// RegisterProvider registers a speech provider
func (s *TranscriptionService) RegisterProvider(name string, provider SpeechProvider) {
	s.providers[name] = provider
	if s.primary == "" {
		s.primary = name
	} else if s.fallback == "" {
		s.fallback = name
	}
}

// Transcribe transcribes audio using the primary provider with fallback
func (s *TranscriptionService) Transcribe(ctx context.Context, audioData []byte) (*domain.TranscriptChunk, error) {
	provider, ok := s.providers[s.primary]
	if !ok {
		return nil, errors.New("no speech provider available")
	}

	transcript, err := provider.Transcribe(ctx, audioData)
	if err != nil {
		s.logger.Warn("primary provider failed, trying fallback", "error", err.Error())
		
		if s.fallback != "" {
			if fallbackProvider, ok := s.providers[s.fallback]; ok {
				return fallbackProvider.Transcribe(ctx, audioData)
			}
		}
		return nil, err
	}

	return transcript, nil
}

// SummarizationService handles text summarization
type SummarizationService struct {
	logger *logger.Logger
}

// NewSummarizationService creates a new summarization service
func NewSummarizationService(logger *logger.Logger) *SummarizationService {
	return &SummarizationService{logger: logger}
}

// SummarizeTranscript creates summaries from transcript
func (s *SummarizationService) SummarizeTranscript(ctx context.Context, transcripts []domain.Transcript) (*domain.Summary, error) {
	if len(transcripts) == 0 {
		return nil, errors.New("no transcripts to summarize")
	}

	// Build combined text from transcripts
	var combinedText strings.Builder
	for _, t := range transcripts {
		fmt.Fprintf(&combinedText, "%s: %s\n", t.SpeakerName, t.Text)
	}

	// For now, create a simple summary (in production, would use ML model)
	summary := &domain.Summary{
		Content: extractKeySentences(combinedText.String(), 5),
		KeyPoints: generateKeyPoints(transcripts),
		SummaryType: "concise",
	}

	return summary, nil
}

// TaskExtractionService extracts action items from transcripts
type TaskExtractionService struct {
	logger *logger.Logger
}

// NewTaskExtractionService creates a new task extraction service
func NewTaskExtractionService(logger *logger.Logger) *TaskExtractionService {
	return &TaskExtractionService{logger: logger}
}

// ExtractTasks extracts tasks from transcripts
func (s *TaskExtractionService) ExtractTasks(ctx context.Context, transcripts []domain.Transcript) ([]domain.Task, error) {
	var tasks []domain.Task
	
	// Keywords that indicate action items
	actionKeywords := []string{"todo", "action", "task", "need to", "should", "must", "will", "going to"}
	
	for _, t := range transcripts {
		lowerText := strings.ToLower(t.Text)
		for _, keyword := range actionKeywords {
			if strings.Contains(lowerText, keyword) {
				task := &domain.Task{
					Title:       extractTaskTitle(t.Text),
					Description: t.Text,
					Owner:       t.SpeakerName,
					Status:      "pending",
					Priority:    "medium",
				}
				tasks = append(tasks, task)
				break
			}
		}
	}
	
	return tasks, nil
}

// ConflictDetectionService detects conflicts in meetings
type ConflictDetectionService struct {
	logger *logger.Logger
}

// NewConflictDetectionService creates a new conflict detection service
func NewConflictDetectionService(logger *logger.Logger) *ConflictDetectionService {
	return &ConflictDetectionService{logger: logger}
}

// DetectConflicts detects conflicts in transcripts
func (s *ConflictDetectionService) DetectConflicts(ctx context.Context, transcripts []domain.Transcript) ([]domain.Conflict, error) {
	var conflicts []domain.Conflict
	
	// Keywords that indicate disagreement
	conflictKeywords := []string{"disagree", "don't think", "but", "however", "instead", "no", "wrong", "issue"}
	
	for i := 0; i < len(transcripts)-1; i++ {
		current := transcripts[i]
		next := transcripts[i+1]
		
		// Check for sentiment changes
		if current.Sentiment == "positive" && next.Sentiment == "negative" {
			conflict := &domain.Conflict{
				Description: fmt.Sprintf("Potential disagreement between %s and %s", current.SpeakerName, next.SpeakerName),
				ConflictScore: 0.7,
				Status: "unresolved",
			}
			conflicts = append(conflicts, conflict)
		}
		
		// Check for explicit conflict keywords
		for _, keyword := range conflictKeywords {
			if strings.Contains(strings.ToLower(current.Text), keyword) {
				conflict := &domain.Conflict{
					Description: fmt.Sprintf("Conflict detected: %s", extractSentence(current.Text, keyword)),
					ConflictScore: 0.6,
					Status: "unresolved",
				}
				conflicts = append(conflicts, conflict)
				break
			}
		}
	}
	
	return conflicts, nil
}

// ConfusionDetectionService detects confusion signals
type ConfusionDetectionService struct {
	logger *logger.Logger
}

// NewConfusionDetectionService creates a new confusion detection service
func NewConfusionDetectionService(logger *logger.Logger) *ConfusionDetectionService {
	return &ConfusionDetectionService{logger: logger}
}

// DetectConfusion detects confusion signals in transcripts
func (s *ConfusionDetectionService) DetectConfusion(ctx context.Context, transcripts []domain.Transcript) ([]domain.Confusion, error) {
	var confusions []domain.Confusion
	
	// Keywords that indicate confusion
	confusionKeywords := []string{"what", "huh", "what's that", "pardon", "can you repeat", "didn't catch", "sorry"}
	
	for _, t := range transcripts {
		lowerText := strings.ToLower(t.Text)
		for _, keyword := range confusionKeywords {
			if strings.Contains(lowerText, keyword) {
				confusion := &domain.Confusion{
					Participant: t.SpeakerName,
					Description: t.Text,
					ClarificationCount: 1,
				}
				confusions = append(confusions, confusion)
				break
			}
		}
	}
	
	return confusions, nil
}

// Helper functions

func extractKeySentences(text string, count int) string {
	sentences := strings.Split(text, ".")
	if len(sentences) > count {
		sentences = sentences[:count]
	}
	return strings.Join(sentences, ".")
}

func generateKeyPoints(transcripts []domain.Transcript) string {
	points := []string{}
	seen := make(map[string]bool)
	
	for _, t := range transcripts {
		words := strings.Fields(t.Text)
		if len(words) > 5 && !seen[t.Text] {
			points = append(points, t.Text)
			seen[t.Text] = true
		}
	}
	
	if len(points) > 3 {
		points = points[:3]
	}
	
	jsonBytes, _ := json.Marshal(points)
	return string(jsonBytes)
}

func extractTaskTitle(text string) string {
	words := strings.Fields(text)
	if len(words) > 5 {
		return strings.Join(words[:5], " ")
	}
	return text
}

func extractSentence(text, keyword string) string {
	idx := strings.Index(strings.ToLower(text), strings.ToLower(keyword))
	if idx == -1 {
		return text
	}
	
	start := idx
	if start > 50 {
		start = idx - 50
	}
	
	end := idx + len(keyword) + 50
	if end > len(text) {
		end = len(text)
	}
	
	return text[start:end]
}

// SentimentService analyzes sentiment
type SentimentService struct {
	logger *logger.Logger
}

// NewSentimentService creates a new sentiment service
func NewSentimentService(logger *logger.Logger) *SentimentService {
	return &SentimentService{logger: logger}
}

// AnalyzeSentiment analyzes text sentiment
func (s *SentimentService) AnalyzeSentiment(ctx context.Context, text string) string {
	lowerText := strings.ToLower(text)
	
	positiveWords := []string{"good", "great", "excellent", "amazing", "love", "happy", "awesome"}
	negativeWords := []string{"bad", "terrible", "hate", "awful", "angry", "sad", "horrible"}
	
	positiveScore := 0
	negativeScore := 0
	
	for _, word := range positiveWords {
		if strings.Contains(lowerText, word) {
			positiveScore++
		}
	}
	
	for _, word := range negativeWords {
		if strings.Contains(lowerText, word) {
			negativeScore++
		}
	}
	
	if positiveScore > negativeScore {
		return "positive"
	} else if negativeScore > positiveScore {
		return "negative"
	}
	
	return "neutral"
}

// DecisionService extracts and analyzes decisions
type DecisionService struct {
	logger *logger.Logger
}

// NewDecisionService creates a new decision service
func NewDecisionService(logger *logger.Logger) *DecisionService {
	return &DecisionService{logger: logger}
}

// ExtractDecisions extracts decisions from transcripts
func (s *DecisionService) ExtractDecisions(ctx context.Context, transcripts []domain.Transcript) ([]domain.Decision, error) {
	var decisions []domain.Decision
	
	// Keywords that indicate decisions
	decisionKeywords := []string{"decide", "decided", "let's", "we'll", "approved", "final", "launch"}
	
	for _, t := range transcripts {
		lowerText := strings.ToLower(t.Text)
		for _, keyword := range decisionKeywords {
			if strings.Contains(lowerText, keyword) {
				decision := &domain.Decision{
					Title: extractTaskTitle(t.Text),
					Description: t.Text,
					Confidence: 0.75,
					Resolved: true,
				}
				decisions = append(decisions, decision)
				break
			}
		}
	}
	
	return decisions, nil
}
