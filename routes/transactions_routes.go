package routes

import (
	"lmizania/config"
	"lmizania/controllers"
	"lmizania/database"
	"lmizania/middlewares"
	"lmizania/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func TransactionRoutes(router *mux.Router) {
	
	transactionColl := database.MongoClient.Database(config.DB_NAME).Collection("transactions")
	userColl := database.MongoClient.Database(config.DB_NAME).Collection("users")

	// Initialize the UserRepo with the user collection
	userRepo := &repository.UserRepo{MongoCollection: userColl}

	// Initialize the TransactionService with the transaction collection and UserRepo
	transactionService := controllers.TransactionService{
		MongoCollection: transactionColl,
		UserRepo:        userRepo,
	}

	// Define the transaction routes with the necessary middlewares
	router.HandleFunc("/transactions", middlewares.LoginRequired(transactionService.AddTransaction)).Methods(http.MethodPost)
	router.HandleFunc("/transactions/{id}", middlewares.LoginRequired(transactionService.UpdateTransaction)).Methods(http.MethodPut)
	router.HandleFunc("/transactions/{id}", middlewares.LoginRequired(transactionService.DeleteTransaction)).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", middlewares.LoginRequired(transactionService.GetAllTransactions)).Methods(http.MethodGet)
}
