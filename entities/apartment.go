package entities

// Apartment :
type Apartment struct {
	Base
	UserID       int64             `json:"user_id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	RoomType     string            `json:"room_type"`
	PropertyType string            `json:"property_type"`
	Street       string            `json:"street"`
	City         string            `json:"city"`
	State        string            `json:"state"`
	Country      string            `json:"country"`
	Guests       int               `json:"guests"`
	Bedrooms     int               `json:"bedrooms"`
	Beds         int               `json:"beds"`
	Baths        int               `json:"baths"`
	Likes        int               `json:"likes"`
	Price        float64           `json:"price"`
	Status       string            `json:"status"`
	Images       []*ApartmentImage `json:"images"`
	User         *User             `json:"user"`
}
