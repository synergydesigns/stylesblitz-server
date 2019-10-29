package models

type Category struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	Image       string
	VendorID    uint64 `json:"Vendor_id"`
	ShopID      uint64 `json:"shop_id"`
}
