package controllers

import (
	"encoding/json"
	"lmizania/models"
	"lmizania/pkg/types"
	"lmizania/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	MongoCollection *mongo.Collection
	UserRepo        *repository.UserRepo // Added UserRepo to TransactionService
}

// Add Transaction
func (svc *TransactionService) AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Assign the userID from the context
	transaction.UserID = r.Context().Value("userID").(string)

	// Initialize the TransactionRepo with UserRepo
	repo := repository.TransactionRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	result, err := repo.AddTransaction(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error adding transaction", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res.StatusCode = http.StatusCreated
	res.Data = result
}

// Update Transaction
func (svc *TransactionService) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	transactionID := mux.Vars(r)["id"]
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Initialize the TransactionRepo with UserRepo
	repo := repository.TransactionRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}
	transaction.UserID = r.Context().Value("userID").(string)

	result, err := repo.UpdateTransaction(transactionID, &transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error updating transaction", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = result
}

// Delete Transaction
func (svc *TransactionService) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	transactionID := mux.Vars(r)["id"]

	// Initialize the TransactionRepo with UserRepo
	repo := repository.TransactionRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	err := repo.DeleteTransaction(transactionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error deleting transaction", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Transaction deleted successfully"
}

// Get All Transactions
func (svc *TransactionService) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	// Initialize the TransactionRepo with UserRepo
	repo := repository.TransactionRepo{
		MongoCollection: svc.MongoCollection,
		UserRepo:        svc.UserRepo,
	}

	transactions, err := repo.GetAllTransactions(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving transactions", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = transactions
}
