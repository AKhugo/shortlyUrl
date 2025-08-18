package service

import (
	"errors"
	"shortlyurl/internal/domain"
	"shortlyurl/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(formUserData *domain.User) error {
	// Validate user data using domain rules
	if err := formUserData.IsValid(); err != nil {
		return err
	}

	// Check if user already exists
	if _, err := s.userRepo.GetUserByEmail(formUserData.Email); err == nil {
		return errors.New("user already exists")
	}

	// Check if username already exists
	if _, err := s.userRepo.GetUserByUsername(formUserData.Username); err == nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formUserData.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	formUserData.Password = string(hashedPassword)

	// Generate verification token
	verificationToken := uuid.New().String()

	// Verification token expires in 3 hours
	verificationExpires := time.Now().Add(time.Hour * 3)

	formUserData.VerificationToken = verificationToken
	formUserData.VerificationExpires = verificationExpires

	userData := &domain.User{
		Username:            formUserData.Username,
		Email:               formUserData.Email,
		Password:            formUserData.Password,
		FirstName:           formUserData.FirstName,
		LastName:            formUserData.LastName,
		VerificationToken:   verificationToken,
		VerificationExpires: verificationExpires,
		IsActive:            true,
		IsVerified:          false,
		FailedLoginAttempts: 0,
		LastLoginAt:         time.Time{},
		FailedLoginAt:       time.Time{},
		LastFailedLoginIP:   "",
	}

	return s.userRepo.CreateUser(userData)
}

func (s *UserService) AuthenticateUser(email, password string) (*domain.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if account is locked
	if !user.CanLogin() {
		return nil, errors.New("account is locked")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		user.IncrementFailedLogin()
		s.userRepo.UpdateUser(user)
		return nil, errors.New("invalid credentials")
	}

	// Reset failed login attempts on successful login
	user.ResetFailedLoginAttempts()
	user.LastLoginAt = time.Now()

	return user, s.userRepo.UpdateUser(user)
}
