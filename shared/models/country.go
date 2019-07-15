package models

// Country defines models for countries
type Country struct {
	ID             uint64 `gorm:"primary_key"`
	Name           string
	CountryCode    int
	IsoCode        string
	Longitude      float32
	Latitude       float32
	Currency       string
	CurrencySymbol string
	CurrencyCode   string
}
