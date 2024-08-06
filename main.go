package main

import (
	"lmizania/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.AuthRoutes(router)

	http.ListenAndServe(":8080", router)

}
