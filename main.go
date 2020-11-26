package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CezarGarrido/airbnb-go/config"
	"github.com/CezarGarrido/airbnb-go/driver"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	portDB   = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "airbnb"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portDB, user, password, dbname)

	db, err := driver.NewConnectionPostgres(psqlInfo)
	if err != nil {
		log.Println("Error connect to database:", err.Error())
		return
	}

	config.InitRoutes(router, db)

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headersOk, methodsOk, originsOk)(router)))

}
