package auth

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// TestGetAPIKey calls the auth.GetAPIKey with headers, checking
// for valid return values
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		expected    string
		expectedErr error
	}{
		{"Valid Token", "ApiKey abc123", "abc123", nil},
		{"No Header", "", "", ErrNoAuthHeaderIncluded},
		{"Invalid Format", "ApiKe abc123", "", ErrInvalidAuthHeader},
		{"No Token", "ApiKey ", "", ErrInvalidAuthHeader},
		{"Extra Spaces", "ApiKey  abc123", "", ErrInvalidAuthHeader},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.apiKey != "" {
				headers.Set("Authorization", tt.apiKey)
			}

			token, err := GetAPIKey(headers)
			assert.Equal(t, tt.expected, token)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
