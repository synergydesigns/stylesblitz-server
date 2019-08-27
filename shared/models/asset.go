package models

import "time"

type Asset struct {
	ID          string `gorm:"primary_key"`
	Title       string
	Description string
	Caption     string
	Alt         string
	MediaType   string
	MimeType    string
	Width       int
	Height      int
	Filename    string
	Size        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
