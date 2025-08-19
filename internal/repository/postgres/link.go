package postgres

import (
	"shortlyurl/internal/domain"
	"time"

	"gorm.io/gorm"
)

type LinkRepository struct {
	db        *gorm.DB
	clickRepo *ClickRepository
}

// NewLinkRepository creates a new link repository
func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db:        db,
		clickRepo: NewClickRepository(db),
	}
}

// CreateLink creates a new link
func (r *LinkRepository) CreateLink(link *domain.Link) error {
	return r.db.Create(link).Error
}

// GetLinkByID retrieves a link by its ID
func (r *LinkRepository) GetLinkByID(id uint) (*domain.Link, error) {
	var link domain.Link
	if err := r.db.First(&link, id).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// GetLinkByShortCode retrieves a link by its short code
func (r *LinkRepository) GetLinkByShortCode(shortCode string) (*domain.Link, error) {
	var link domain.Link
	if err := r.db.Where("short_code = ? AND is_active = ?", shortCode, true).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// GetLinksByUserID retrieves all links for a specific user
func (r *LinkRepository) GetLinksByUserID(userID uint) ([]*domain.Link, error) {
	var links []*domain.Link
	if err := r.db.Where("user_id = ?", userID).Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

// UpdateLink updates an existing link
func (r *LinkRepository) UpdateLink(link *domain.Link) error {
	return r.db.Save(link).Error
}

// DeleteLink soft deletes a link
func (r *LinkRepository) DeleteLink(id uint) error {
	return r.db.Delete(&domain.Link{}, id).Error
}

// IncrementClickCount creates a new click record for tracking
func (r *LinkRepository) IncrementClickCount(id uint) error {
	// Create a new click record instead of just updating timestamp
	click := &domain.Click{
		LinkID:    id,
		ClickedAt: time.Now(),
	}
	return r.clickRepo.CreateClick(click)
}

// GetLinkStats retrieves click statistics for a specific link
func (r *LinkRepository) GetLinkStats(id uint) ([]*domain.Click, error) {
	return r.clickRepo.GetClicksByLinkID(id)
}
