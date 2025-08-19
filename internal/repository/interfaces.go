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

// LinkRepository définit le contrat pour l'accès aux données des liens
type LinkRepository interface {
	CreateLink(link *domain.Link) error
	GetLinkByID(id uint) (*domain.Link, error)
	GetLinkByShortCode(shortCode string) (*domain.Link, error)
	GetLinksByUserID(userID uint) ([]*domain.Link, error)
	UpdateLink(link *domain.Link) error
	DeleteLink(id uint) error
	IncrementClickCount(id uint) error
	GetLinkStats(id uint) ([]*domain.Click, error)
}

// ClickRepository définit le contrat pour l'accès aux données des clics
type ClickRepository interface {
	CreateClick(click *domain.Click) error
	GetClicksByLinkID(linkID uint) ([]*domain.Click, error)
	GetClickStats(linkID uint, filter *domain.ClickFilter) ([]*domain.Click, error)
	DeleteClick(id uint) error
}
