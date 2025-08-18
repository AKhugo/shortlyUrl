package domain

import "time"

type Link struct {

	// Primary key
	ID        uint   `gorm:"primaryKey"`

	// Foreign key
	UserID uint `gorm:"foreignKey:ID"`

	// Link information
	ShortCode string `gorm:"unique"` // short code is the unique identifier for the link
	OriginalURL string // original url is the url that the link will redirect to
	AccessLimit int // access limit is the number of times the link can be accessed
	Domain string // domain is the domain that the link will be hosted on

	// Link status
	IsActive bool // is active is a boolean that indicates if the link is active
	IsPublic bool // is public is a boolean that indicates if the link is public

	// Metadata
	CreatedAt time.Time
	UpdatedAt time.Time

	// Deletion information
	DeletedAt time.Time
}