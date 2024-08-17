package routes

import (
	"lmizania/config"
	"lmizania/controllers"
	"lmizania/database"
	"net/http"

	"lmizania/middlewares"

	"github.com/gorilla/mux"
)

func TransactionRoutes(router *mux.Router) {

	coll := database.MongoClient.Database(config.DB_NAME).Collection("transactions")
	TransactionService := controllers.TransactionService{MongoCollection: coll}

	router.HandleFunc("/transactions", middlewares.LoginRequired(TransactionService.AddTransaction)).Methods(http.MethodPost)
	router.HandleFunc("/transactions/{id}", middlewares.LoginRequired(TransactionService.UpdateTransaction)).Methods(http.MethodPut)
	router.HandleFunc("/transactions/{id}", middlewares.LoginRequired(TransactionService.DeleteTransaction)).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", middlewares.LoginRequired(TransactionService.GetAllTransactions)).Methods(http.MethodGet)
}
