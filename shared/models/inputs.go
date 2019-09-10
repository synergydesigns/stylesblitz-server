package models

type AssetInput struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Alt       string `json:"alt"`
	MediaType string `json:"mediaType"`
	MimeType  string `json:"mimeType"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename"`
	Size      int    `json:"size"`
}

type AssetUploadOutput struct {
	ID        string
	UploadURL string
	AssetURL  string
}

type VendorCategoryInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	VendorID    string  `json:"vendor_id"`
}
