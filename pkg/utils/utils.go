package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID generates a random hex ID
func GenerateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Contains checks if a slice contains a string
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// MapStringToInt converts map[string]interface{} to int
func MapStringToInt(m map[string]interface{}, key string, defaultVal int) int {
	if val, ok := m[key]; ok {
		if intVal, ok := val.(float64); ok {
			return int(intVal)
		}
	}
	return defaultVal
}

// MapStringToString converts map[string]interface{} to string
func MapStringToString(m map[string]interface{}, key string, defaultVal string) string {
	if val, ok := m[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return defaultVal
}
