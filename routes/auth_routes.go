package routes

import (
	"lmizania/config"
	"lmizania/controllers"
	"lmizania/database"

	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {

	userColl := database.MongoClient.Database(config.DB_NAME).Collection("users")
	AuthService := controllers.AuthService{MongoCollection: userColl}

	router.HandleFunc("/login", AuthService.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", AuthService.Register).Methods(http.MethodPost)
	router.HandleFunc("/verify/{id}", AuthService.VerifyUser).Methods(http.MethodPost)
	router.HandleFunc("/resetpassword", AuthService.ResetPassword).Methods(http.MethodPost)

}
