package models

type VendorCategory struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
	Vendor      Vendor
}
