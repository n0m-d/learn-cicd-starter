package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "ApiKey 1234567890")
	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("GetAPIKey() error = %v", err)
	}
	if key != "1234567890" {
		t.Errorf("GetAPIKey() key = %v, want %v", key, "1234567890")
	}
}
