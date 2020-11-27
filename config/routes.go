package config

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/CezarGarrido/airbnb-go/driver"
	handlerHttp "github.com/CezarGarrido/airbnb-go/handlers/http"
)

// InitRoutes :
func InitRoutes(router *mux.Router, db *driver.DB) {
	// Create the route
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	router.
		PathPrefix("/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/"))))

	apartmentRoutes(apiV1, db)
	userRoutes(apiV1, db)

}

func apartmentRoutes(router *mux.Router, db *driver.DB) {
	apartment := handlerHttp.NewApartment(db)

	router.HandleFunc("/apartments", apartment.FindAll).Methods("GET")
	router.HandleFunc("/apartments/form", apartment.CreateFormData).Methods("POST")
}

func userRoutes(router *mux.Router, db *driver.DB) {
	user := handlerHttp.NewUser(db)
	router.HandleFunc("/users", user.Create).Methods("POST")
}
