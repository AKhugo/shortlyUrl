package postgres

import "shortlyurl/internal/domain"

type LinkRepository interface {
	CreateLink(link *domain.Link) error
	GetLinkByID(id uint) (*domain.Link, error)
	GetAllUserLinks(userID uint) ([]domain.Link, error)
	GetLinkByShortCode(shortCode string) (*domain.Link, error)
	UpdateLink(link *domain.Link) error
	DeleteLink(id uint) error
}