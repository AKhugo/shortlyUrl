package postgres

import "shortlyurl/internal/domain"

type ClickRepository interface {
	CreateClick(click *domain.Click) error
	GetClickByID(id uint) (*domain.Click, error)
	GetClicksByLinkID(linkID uint) ([]domain.Click, error)
	DeleteClick(id uint) error

	// advanced queries
	GetAllLinkClicks(linkID uint) ([]domain.Click, error)
	GetClicksByFilter(filter *domain.ClickFilter) ([]domain.Click, error)

}