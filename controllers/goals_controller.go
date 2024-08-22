package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"lmizania/models"
	"lmizania/pkg/types"
	"lmizania/repository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoalService struct {
	MongoCollection *mongo.Collection
	UserRepo        *repository.UserRepo // Added UserRepo to GoalService
}

// AddGoal handles the request to add a new goal
func (svc *GoalService) AddGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	var goal models.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Assign the userID from the context
	goal.UserID = r.Context().Value("userID").(string)

	// Initialize the GoalRepo with UserRepo
	repo := repository.GoalRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	result, err := repo.AddGoal(&goal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error adding goal", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res.StatusCode = http.StatusCreated
	res.Data = result
}

// UpdateGoal handles the request to update an existing goal
func (svc *GoalService) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	goalID := mux.Vars(r)["id"]
	var goal models.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Initialize the GoalRepo with UserRepo
	repo := repository.GoalRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}
	goal.UserID = r.Context().Value("userID").(string)

	result, err := repo.UpdateGoal(goalID, &goal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error updating goal", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = result
}

// DeleteGoal handles the request to delete a goal
func (svc *GoalService) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	goalID := mux.Vars(r)["id"]

	// Initialize the GoalRepo with UserRepo
	repo := repository.GoalRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	err := repo.DeleteGoal(goalID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error deleting goal", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Goal deleted successfully"
}

// GetAllGoals handles the request to retrieve all goals for a user
func (svc *GoalService) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	// Initialize the GoalRepo with UserRepo
	repo := repository.GoalRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	goals, err := repo.GetAllGoals(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving goals", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = goals
}

func (svc *GoalService) DepositGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	goalID := mux.Vars(r)["id"]

	var depositRequest struct {
		Amount float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&depositRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid request payload", err)
		res.Error = "Invalid request payload"
		return
	}

	// Initialize the GoalRepo
	repo := repository.GoalRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	err = repo.DepositGoal(goalID, depositRequest.Amount)
	if err != nil {
		if err.Error() == "deposit amount cannot be negative" || err.Error() == "deposit exceeds the total goal amount" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error depositing goal", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Deposit successful"
}
