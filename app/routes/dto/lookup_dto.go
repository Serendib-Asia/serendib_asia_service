package dto

// LookupResponse represents a generic lookup table response
type LookupResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// PurposeTypeResponse represents a purpose type response
type PurposeTypeResponse struct {
	LookupResponse
}

// PropertyTypeResponse represents a property type response
type PropertyTypeResponse struct {
	LookupResponse
}

// FurnitureTypeResponse represents a furniture type response
type FurnitureTypeResponse struct {
	LookupResponse
}

// PropertyConditionResponse represents a property condition response
type PropertyConditionResponse struct {
	LookupResponse
}

// UtilityResponse represents a utility response
type UtilityResponse struct {
	LookupResponse
}

// AmenityResponse represents an amenity response
type AmenityResponse struct {
	LookupResponse
}

// LookupListResponse represents a list of lookup items
type LookupListResponse struct {
	Items []LookupResponse `json:"items"`
}
