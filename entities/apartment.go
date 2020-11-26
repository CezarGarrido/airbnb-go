package entities

// Apartment :
type Apartment struct {
	Base
	UserID      int64   `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
	Guests      int     `json:"guests"`
	Bedrooms    int     `json:"bedrooms"`
	Beds        int     `json:"beds"`
	Baths       int     `json:"baths"`
	Likes       int     `json:"likes"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	User        *User   `json:"user"`
}
