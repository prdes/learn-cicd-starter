package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey some_key")

	apiKey, _ := GetAPIKey(headers)

	expectedApiKey := "some_key"
	if apiKey != expectedApiKey {
		t.Errorf("expected %s, got %s", expectedApiKey, apiKey)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}
	// Intentionally leave out the "Authorization" header

	apiKey, _ := GetAPIKey(headers)

	// Define what you expect for a missing header, e.g., an empty string or an error.
	var expectedApiKey string // adjust to expected return type/value
	if apiKey != expectedApiKey {
		t.Errorf("expected %s, got %s", expectedApiKey, apiKey)
	}
}
