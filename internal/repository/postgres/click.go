package postgres

import (
	"shortlyurl/internal/domain"

	"gorm.io/gorm"
)

type ClickRepository struct {
	db *gorm.DB
}

// NewClickRepository creates a new click repository
func NewClickRepository(db *gorm.DB) *ClickRepository {
	return &ClickRepository{db: db}
}

// CreateClick creates a new click record
func (r *ClickRepository) CreateClick(click *domain.Click) error {
	return r.db.Create(click).Error
}

// GetClicksByLinkID retrieves all clicks for a specific link
func (r *ClickRepository) GetClicksByLinkID(linkID uint) ([]*domain.Click, error) {
	var clicks []*domain.Click
	if err := r.db.Where("link_id = ?", linkID).Order("clicked_at DESC").Find(&clicks).Error; err != nil {
		return nil, err
	}
	return clicks, nil
}

// GetClickStats retrieves click statistics with filtering
func (r *ClickRepository) GetClickStats(linkID uint, filter *domain.ClickFilter) ([]*domain.Click, error) {
	var clicks []*domain.Click
	query := r.db.Where("link_id = ?", linkID)

	// Apply filters if provided
	if filter != nil {
		if !filter.ClickedAt.IsZero() {
			query = query.Where("clicked_at >= ?", filter.ClickedAt)
		}
		if filter.CountryCode != "" {
			query = query.Where("country_code = ?", filter.CountryCode)
		}
		if filter.IPAddress != "" {
			query = query.Where("ip_address = ?", filter.IPAddress)
		}
		if filter.UserAgent != "" {
			query = query.Where("user_agent LIKE ?", "%"+filter.UserAgent+"%")
		}
		if filter.Referer != "" {
			query = query.Where("referer LIKE ?", "%"+filter.Referer+"%")
		}
	}

	// Apply pagination
	if filter != nil && filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	// Order by clicked_at descending
	query = query.Order("clicked_at DESC")

	if err := query.Find(&clicks).Error; err != nil {
		return nil, err
	}
	return clicks, nil
}

// DeleteClick soft deletes a click record
func (r *ClickRepository) DeleteClick(id uint) error {
	return r.db.Delete(&domain.Click{}, id).Error
}
