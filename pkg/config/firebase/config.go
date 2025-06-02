package firebase

import (
	"encoding/base64"
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	// Firebase constants
	ProjectID            = "FIREBASE_PROJECT_ID"
	ServiceAccountBase64 = "FIREBASE_SERVICE_ACCOUNT_BASE64"
)

// Config holds the Firebase configuration
type Config struct {
	_                  struct{}
	ProjectID          string
	ServiceAccountJSON string
}

// GetServiceAccountJSON returns the Firebase service account JSON
func GetServiceAccountJSON() []byte {
	// Get the base64 encoded service account JSON from environment
	serviceAccountBase64 := viper.GetString(ServiceAccountBase64)
	if serviceAccountBase64 == "" {
		return []byte("{}")
	}

	// Decode base64
	decoded, err := base64.StdEncoding.DecodeString(serviceAccountBase64)
	if err != nil {
		return []byte("{}")
	}

	// Validate that it's proper JSON
	if !json.Valid(decoded) {
		return []byte("{}")
	}

	return decoded
}

// SetDefaultConfig sets the default Firebase configuration
func SetDefaultConfig() {
	viper.SetDefault(ProjectID, "")
	viper.SetDefault(ServiceAccountBase64, "")
}

// GetConfig returns the Firebase application configuration
func GetConfig() Config {
	return Config{
		ProjectID:          viper.GetString(ProjectID),
		ServiceAccountJSON: viper.GetString(ServiceAccountBase64),
	}
}
