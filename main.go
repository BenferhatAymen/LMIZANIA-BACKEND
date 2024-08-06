package main

import (
	"lmizania/config"
	"lmizania/database"
	"lmizania/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadENV()
	database.Client = database.DBInstance()

	router := mux.NewRouter()

	routes.AuthRoutes(router)

	http.ListenAndServe(":8080", router)

}
