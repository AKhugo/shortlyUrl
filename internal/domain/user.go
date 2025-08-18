package domain

import (
	"errors"
	"shortlyurl/pkg/validator"
	"time"
)

type User struct {

	// Primary key
	ID uint `gorm:"primaryKey"`

	// Identity & Authentication information
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string

	// Account information
	FirstName   string
	LastName    string
	FullName    string
	AvatarURL   string
	PhoneNumber string

	// Confirmation d'email
	EmailVerified       bool
	VerificationToken   string
	VerificationExpires time.Time

	// Account status
	IsActive   bool
	IsVerified bool

	// Forgot password
	ForgotPasswordToken   string
	ForgotPasswordExpires time.Time

	// Remember me
	RememberMeToken   string
	RememberMeExpires time.Time

	// Connection information
	LastLoginAt         time.Time
	FailedLoginAttempts int
	FailedLoginAt       time.Time
	LastFailedLoginIP   string

	// Deletion information
	DeletedAt time.Time

	// Metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	LastIP    string
	UserAgent string
}

// Validation methods using pkg/validator
func (u *User) IsValid() error {
	if err := validator.IsValidUsername(u.Username); err != nil {
		return errors.New("username: " + err.Error())
	}

	if err := validator.IsValidEmail(u.Email); err != nil {
		return errors.New("email: " + err.Error())
	}

	if err := validator.IsValidPassword(u.Password); err != nil {
		return errors.New("password: " + err.Error())
	}

	return nil
}

// Business logic methods
func (u *User) IsAccountLocked() bool {
	return u.FailedLoginAttempts >= 5 && time.Since(u.FailedLoginAt) < time.Hour
}

func (u *User) CanLogin() bool {
	return u.IsActive && !u.IsAccountLocked()
}

func (u *User) IncrementFailedLogin() {
	u.FailedLoginAttempts++
	u.FailedLoginAt = time.Now()
}

func (u *User) ResetFailedLoginAttempts() {
	u.FailedLoginAttempts = 0
	u.FailedLoginAt = time.Time{}
}

func (u *User) IsVerificationTokenExpired() bool {
	return time.Now().After(u.VerificationExpires)
}

func (u *User) IsForgotPasswordTokenExpired() bool {
	return time.Now().After(u.ForgotPasswordExpires)
}

func (u *User) IsRememberMeTokenExpired() bool {
	return time.Now().After(u.RememberMeExpires)
}
