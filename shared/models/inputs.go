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

type ServiceCartInput struct {
	VendorID    string  `json:"vendor_id"`
	ServiceID   int     `json:"service_id"`
	Quantity    int     `json:"quantity"`
}

type ProductInput struct {
	Name        string `json:"name"`
	VendorID    string `json:"vendor_id"`
	Available   int    `json:"available"`
	CategoryID  string `json:"category_id"`
	BrandID     string `json:"brand_id"`
}

type ProductCartInput struct {
	VendorID    string  `json:"vendor_id"`
	ProductID string `json:"product_id"`
	ShopID string `json:"shop_id"`
	Quantity int `json:"quantity"`
}

type CartUpdateInput struct {
	CartID   string  `json:"cart_id"`
	Quantity int     `json:"quantity"`
	TypeID   string  `json:"type_id"`
	Type     string   `json:"type"`
}

type CartInput struct {
	VendorID    string      `json:"vendor_id"`
	TypeID      string      `json:"type_id"`
	Type     		string      `json:"type"`
	Quantity    int         `json:"quantity"`
}
