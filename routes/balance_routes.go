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

func BalanceRoutes(router *mux.Router) {
	// Initialize the MongoDB collection for users
	userColl := database.MongoClient.Database(config.DB_NAME).Collection("users")

	// Initialize the UserRepo with the user collection
	userRepo := &repository.UserRepo{MongoCollection: userColl}

	// Initialize the BalanceService with UserRepo
	balanceService := controllers.BalanceService{
		UserRepo: userRepo,
	}

	// Define the balance routes with the necessary middlewares
	router.HandleFunc("/balance/wallet", middlewares.LoginRequired(balanceService.GetWallet)).Methods(http.MethodGet)
	router.HandleFunc("/balance/wallet", middlewares.LoginRequired(balanceService.SetWallet)).Methods(http.MethodPut)
	router.HandleFunc("/balance/target", middlewares.LoginRequired(balanceService.GetTarget)).Methods(http.MethodGet)
	router.HandleFunc("/balance/target", middlewares.LoginRequired(balanceService.SetTarget)).Methods(http.MethodPut)
	router.HandleFunc("/balance/savings", middlewares.LoginRequired(balanceService.DepositSavings)).Methods(http.MethodPost)
	router.HandleFunc("/balance/savings", middlewares.LoginRequired(balanceService.GetSavings)).Methods(http.MethodGet)
	router.HandleFunc("/balance/income", middlewares.LoginRequired(balanceService.GetIncome)).Methods(http.MethodGet)
	router.HandleFunc("/balance/expense", middlewares.LoginRequired(balanceService.GetExpense)).Methods(http.MethodGet)
}
