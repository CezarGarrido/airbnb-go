package entities

// ApartmentImage :
type ApartmentImage struct {
	Base
	ApartmentID int64  `json:"apartment_id"`
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size"`
	URL         string `json:"url"`
}
