package http

import (
	"log"
	"net/http"

	"github.com/CezarGarrido/airbnb-go/driver"
	"github.com/CezarGarrido/airbnb-go/lib/render"
	apartmentRepo "github.com/CezarGarrido/airbnb-go/repositories/apartment"
)

// NewApartment :
func NewApartment(db *driver.DB) *Apartment {
	return &Apartment{
		ApartmentRepo: apartmentRepo.NewPgApartmentRepo(db.SQL, nil),
	}
}

// Apartment :
type Apartment struct {
	ApartmentRepo apartmentRepo.ApartmentRepo
}

// FindAll :
func (ap *Apartment) FindAll(w http.ResponseWriter, r *http.Request) {
	apartments, err := ap.ApartmentRepo.FindAll(r.Context())
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	render.HTTPWriteJSON(w, http.StatusOK, apartments)
}
