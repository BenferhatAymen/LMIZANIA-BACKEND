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

func GoalRoutes(router *mux.Router) {
	// Initialize the MongoDB collections for goals and users
	goalColl := database.MongoClient.Database(config.DB_NAME).Collection("goals")
	userColl := database.MongoClient.Database(config.DB_NAME).Collection("users")

	// Initialize the UserRepo with the user collection
	userRepo := &repository.UserRepo{MongoCollection: userColl}

	// Initialize the GoalService with the goal collection and UserRepo
	goalService := controllers.GoalService{
		MongoCollection: goalColl,
		UserRepo:        userRepo,
	}

	// Define the goal routes with the necessary middlewares
	router.HandleFunc("/goals", middlewares.LoginRequired(goalService.AddGoal)).Methods(http.MethodPost)
	router.HandleFunc("/goals/{id}", middlewares.LoginRequired(goalService.UpdateGoal)).Methods(http.MethodPut)
	router.HandleFunc("/goals/{id}", middlewares.LoginRequired(goalService.DeleteGoal)).Methods(http.MethodDelete)
	router.HandleFunc("/goals", middlewares.LoginRequired(goalService.GetAllGoals)).Methods(http.MethodGet)
	router.HandleFunc("/goals/{id}/deposit", middlewares.LoginRequired(goalService.DepositGoal)).Methods(http.MethodPost)
}
