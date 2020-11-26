package entities

import "time"

// Booking :
type Booking struct {
	Base
	ApartmentID int64     `json:"apartment_id"`
	LocatorID   int64     `json:"locator_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Locator     *User     `json:"locator"`
}
