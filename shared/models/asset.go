package models

import "time"

// Asset defines the Asset models for graphql
// for getting a single Asset
type Asset struct {
	ID          string `gorm:"primary_key"`
	Title       string
	Description string
	Caption     string
	Alt         string
	MediaType   string
	MimeType    string
	Width       uint64
	Height      uint64
	Filename    string
	Size        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
