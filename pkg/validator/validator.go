 package validator

import (
	"errors"
	"regexp"
	"strings"
)

// Common validation errors
var (
	ErrEmptyField    = errors.New("field cannot be empty")
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrInvalidURL    = errors.New("invalid URL format")
	ErrTooShort      = errors.New("field is too short")
	ErrTooLong       = errors.New("field is too long")
	ErrInvalidFormat = errors.New("invalid format")
)

// Email validation
func IsValidEmail(email string) error {
	if email == "" {
		return ErrEmptyField
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

// Password validation
func IsValidPassword(password string) error {
	if password == "" {
		return ErrEmptyField
	}

	if len(password) < 8 {
		return ErrTooShort
	}

	if len(password) > 128 {
		return ErrTooLong
	}

	return nil
}

// Username validation
func IsValidUsername(username string) error {
	if username == "" {
		return ErrEmptyField
	}

	if len(username) < 3 {
		return ErrTooShort
	}

	if len(username) > 30 {
		return ErrTooLong
	}

	// Username should only contain alphanumeric characters and underscores
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return ErrInvalidFormat
	}

	return nil
}

// URL validation
func IsValidURL(url string) error {
	if url == "" {
		return ErrEmptyField
	}

	// Basic URL validation
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return ErrInvalidURL
	}

	return nil
}

// String length validation
func IsValidLength(field, value string, min, max int) error {
	if value == "" && min > 0 {
		return ErrEmptyField
	}

	if len(value) < min {
		return ErrTooShort
	}

	if len(value) > max {
		return ErrTooLong
	}

	return nil
}

// Required field validation
func IsRequired(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return ErrEmptyField
	}
	return nil
}
