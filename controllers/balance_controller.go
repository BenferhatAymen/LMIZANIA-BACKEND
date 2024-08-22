package controllers

import (
	"encoding/json"
	"lmizania/pkg/types"
	"lmizania/repository"
	"log"
	"net/http"
)

type BalanceService struct {
	UserRepo *repository.UserRepo // UserRepo to interact with user data
}

var data struct {
	Amount float64 `json:"amount"`
}

// GetWallet - Retrieves the user's wallet balance
func (svc *BalanceService) GetWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	wallet, err := svc.UserRepo.GetWallet(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving wallet", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = wallet
}

// SetWallet - Updates the user's wallet balance
func (svc *BalanceService) SetWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	var data struct {
		Wallet float64 `json:"wallet"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	err = svc.UserRepo.SetWallet(userID, data.Wallet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error setting wallet", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Wallet updated successfully"
}

// GetTarget - Retrieves the user's financial target
func (svc *BalanceService) GetTarget(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	target, err := svc.UserRepo.GetTarget(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving target", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = target
}

// SetTarget - Updates the user's financial target
func (svc *BalanceService) SetTarget(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	err = svc.UserRepo.SetTarget(userID, data.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error setting target", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Target updated successfully"
}

// DepositSavings - Adds an amount to the user's savings
func (svc *BalanceService) DepositSavings(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	err = svc.UserRepo.DepositSavings(userID, data.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error depositing savings", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = "Savings deposited successfully"
}

// GetSavings - Retrieves the user's total savings
func (svc *BalanceService) GetSavings(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	savings, err := svc.UserRepo.GetSavings(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving savings", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = savings
}

// GetIncome - Retrieves the user's total income
func (svc *BalanceService) GetIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	income, err := svc.UserRepo.GetIncome(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving income", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = income
}

// GetExpense - Retrieves the user's total expenses
func (svc *BalanceService) GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &types.APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := r.Context().Value("userID").(string)

	expenses, err := svc.UserRepo.GetExpense(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()
		log.Println("Error retrieving expenses", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = expenses
}
