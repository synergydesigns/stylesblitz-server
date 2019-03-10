package models

// Category defines the category models for graphql
// for getting a single category
type Category struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	Image       string
	ProviderID  uint64 `json:"provider_id"`
	ShopID      uint64 `json:"shop_id"`
}
