package dto

// UploadImageRequest represents the request to upload a property image
type UploadImageRequest struct {
	URL       string `json:"url" validate:"required,url"`
	IsPrimary bool   `json:"is_primary"`
}

// ImageResponse represents a property image
type ImageResponse struct {
	ID         uint   `json:"id"`
	PropertyID uint   `json:"property_id"`
	URL        string `json:"url"`
	IsPrimary  bool   `json:"is_primary"`
}

// ImageListResponse represents a list of property images
type ImageListResponse struct {
	Items []ImageResponse `json:"items"`
}
