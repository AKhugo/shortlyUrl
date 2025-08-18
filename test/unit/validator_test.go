package unit

import (
	"testing"
	"shortlyurl/pkg/validator"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected error
	}{
		{"valid email", "test@example.com", nil},
		{"valid email with subdomain", "test@sub.example.com", nil},
		{"empty email", "", validator.ErrEmptyField},
		{"invalid format", "invalid-email", validator.ErrInvalidEmail},
		{"missing @", "testexample.com", validator.ErrInvalidEmail},
		{"missing domain", "test@", validator.ErrInvalidEmail},
		{"missing tld", "test@example", validator.ErrInvalidEmail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.IsValidEmail(tt.email)
			if err != tt.expected {
				t.Errorf("IsValidEmail(%s) = %v, want %v", tt.email, err, tt.expected)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected error
	}{
		{"valid password", "password123", nil},
		{"empty password", "", validator.ErrEmptyField},
		{"too short", "123", validator.ErrTooShort},
		{"exactly 8 chars", "12345678", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.IsValidPassword(tt.password)
			if err != tt.expected {
				t.Errorf("IsValidPassword(%s) = %v, want %v", tt.password, err, tt.expected)
			}
		})
	}
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		expected error
	}{
		{"valid username", "testuser", nil},
		{"valid with underscore", "test_user", nil},
		{"empty username", "", validator.ErrEmptyField},
		{"too short", "ab", validator.ErrTooShort},
		{"invalid chars", "test-user", validator.ErrInvalidFormat},
		{"with spaces", "test user", validator.ErrInvalidFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.IsValidUsername(tt.username)
			if err != tt.expected {
				t.Errorf("IsValidUsername(%s) = %v, want %v", tt.username, err, tt.expected)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected error
	}{
		{"valid http url", "http://example.com", nil},
		{"valid https url", "https://example.com", nil},
		{"empty url", "", validator.ErrEmptyField},
		{"invalid protocol", "ftp://example.com", validator.ErrInvalidURL},
		{"no protocol", "example.com", validator.ErrInvalidURL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.IsValidURL(tt.url)
			if err != tt.expected {
				t.Errorf("IsValidURL(%s) = %v, want %v", tt.url, err, tt.expected)
			}
		})
	}
}
