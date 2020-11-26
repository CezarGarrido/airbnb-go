package config

import (
	"github.com/gorilla/mux"

	"github.com/CezarGarrido/airbnb-go/driver"
	handlerHttp "github.com/CezarGarrido/airbnb-go/handlers/http"
)

// InitRoutes :
func InitRoutes(router *mux.Router, db *driver.DB) {

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apartmentRoutes(apiV1, db)
}

func apartmentRoutes(router *mux.Router, db *driver.DB) {
	apartment := handlerHttp.NewApartment(db)

	router.HandleFunc("/apartments", apartment.FindAll).Methods("GET")
}
