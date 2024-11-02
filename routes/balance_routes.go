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
	userColl := database.MongoClient.Database(config.DB_NAME).Collection("users")

	userRepo := &repository.UserRepo{MongoCollection: userColl}

	balanceService := controllers.BalanceService{
		UserRepo: userRepo,
	}

	router.HandleFunc("/balance/wallet", middlewares.LoginRequired(balanceService.GetWallet)).Methods(http.MethodGet)
	router.HandleFunc("/balance/wallet", middlewares.LoginRequired(balanceService.SetWallet)).Methods(http.MethodPut)
	router.HandleFunc("/balance/target", middlewares.LoginRequired(balanceService.GetTarget)).Methods(http.MethodGet)
	router.HandleFunc("/balance/target", middlewares.LoginRequired(balanceService.SetTarget)).Methods(http.MethodPut)
	router.HandleFunc("/balance/savings", middlewares.LoginRequired(balanceService.DepositSavings)).Methods(http.MethodPost)
	router.HandleFunc("/balance/savings", middlewares.LoginRequired(balanceService.GetSavings)).Methods(http.MethodGet)
	router.HandleFunc("/balance/income", middlewares.LoginRequired(balanceService.GetIncome)).Methods(http.MethodGet)
	router.HandleFunc("/balance/expense", middlewares.LoginRequired(balanceService.GetExpense)).Methods(http.MethodGet)
}
