package domain

import "time"

type Click struct {

	// Primary key
	ID        uint   `gorm:"primaryKey"`

	// Foreign key
	LinkID uint `gorm:"foreignKey:ID"`

	// Click information
	ClickedAt time.Time // clicked at is the time the link was clicked
	IPAddress string // ip address is the ip address of the user who clicked the link
	UserAgent string // user agent is the user agent of the user who clicked the link
	Referer string // referer is the referer of the user who clicked the link
	CountryCode string // country code is the country code of the user who clicked the link

	// Metadata
	CreatedAt time.Time
	UpdatedAt time.Time

	// Deletion information
	DeletedAt time.Time
}


type ClickFilter struct {
	LinkID uint
	ClickedAt time.Time
	CountryCode string
	IPAddress string
	UserAgent string
	Referer string
	Page int
	Limit int
}