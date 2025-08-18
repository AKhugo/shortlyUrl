package repository

import "shortlyurl/internal/domain"

// UserRepository définit le contrat pour l'accès aux données utilisateur
type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	UpdateUserInfo(email string, userInfo *domain.User) error
	DeleteUser(id uint) error
	HardDeleteUser(id uint) error
}
