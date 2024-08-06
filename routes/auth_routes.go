package routes

import (
	"lmizania/config"
	"lmizania/controllers"
	"lmizania/database"

	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {

	coll := database.Client.Database(config.DB_NAME).Collection("users")
	AuthService := controllers.AuthService{MongoCollection: coll}

	router.HandleFunc("/login", AuthService.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", AuthService.Register).Methods(http.MethodPost)

}
