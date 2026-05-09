package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Server
	ServerPort       string
	ServerHost       string
	Environment      string
	
	// Database
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	DBMaxConnections int
	
	// Redis
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDB          int
	
	// JWT
	JWTSecret        string
	JWTExpiration    int
	
	// AI/ML
	WhisperModel     string
	OllamaURL        string
	OpenAIKey        string
	OpenAIModel      string
	
	// NATS/Kafka
	NATSUrl          string
	KafkaURL         string
	
	// Observability
	JaegerURL        string
	PrometheusPort   string
	
	// Audio Processing
	AudioSampleRate  int
	AudioChannels    int
	AudioBitrate     int
}

// NewConfig creates a new configuration from environment variables
func NewConfig() *Config {
	// Load .env file
	_ = godotenv.Load()

	return &Config{
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		ServerHost:       getEnv("SERVER_HOST", "0.0.0.0"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBName:           getEnv("DB_NAME", "meeting_summarizer"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", "postgres"),
		DBMaxConnections: getEnvInt("DB_MAX_CONNECTIONS", 25),
		
		RedisHost:        getEnv("REDIS_HOST", "localhost"),
		RedisPort:        getEnv("REDIS_PORT", "6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:          getEnvInt("REDIS_DB", 0),
		
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration:    getEnvInt("JWT_EXPIRATION", 86400),
		
		WhisperModel:     getEnv("WHISPER_MODEL", "base"),
		OllamaURL:        getEnv("OLLAMA_URL", "http://localhost:11434"),
		OpenAIKey:        getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:      getEnv("OPENAI_MODEL", "gpt-4-turbo"),
		
		NATSUrl:          getEnv("NATS_URL", "nats://localhost:4222"),
		KafkaURL:         getEnv("KAFKA_URL", "localhost:9092"),
		
		JaegerURL:        getEnv("JAEGER_URL", "http://localhost:14268/api/traces"),
		PrometheusPort:   getEnv("PROMETHEUS_PORT", "9090"),
		
		AudioSampleRate:  getEnvInt("AUDIO_SAMPLE_RATE", 16000),
		AudioChannels:    getEnvInt("AUDIO_CHANNELS", 1),
		AudioBitrate:     getEnvInt("AUDIO_BITRATE", 128000),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
