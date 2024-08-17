package controllers

import (
	"encoding/json"
	"fmt"
	"lmizania/models"
	"lmizania/pkg/types"
	"lmizania/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	MongoCollection *mongo.Collection
}

func (svc *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &types.AuthResponse{}
	defer json.NewEncoder(w).Encode(res)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("invalid body ", err)
		res.Error = err.Error()
		return
	}
	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	token, err := repo.UserLogin(user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in login ", err)
		res.StatusCode = http.StatusBadRequest
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Token = token
}
func (svc *AuthService) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &types.AuthResponse{}
	defer json.NewEncoder(w).Encode(res)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Ensure email and password are provided
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		res.Error = "Email and password are required"
		return
	}

	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	// Register the new user
	result, token, err := repo.RegisterUser(&user)
	if err != nil {
		if err.Error() == "user already exists" {
			w.WriteHeader(http.StatusConflict)
			res.StatusCode = http.StatusConflict
			res.Error = "User already exists"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			res.StatusCode = http.StatusInternalServerError
			res.Error = err.Error()

		}
		log.Println("Error registering user", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res.StatusCode = http.StatusCreated
	res.Data = result
	res.Token = token
}

func (svc *AuthService) VerifyUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &types.AuthResponse{}
	defer json.NewEncoder(w).Encode(res)

	userID := mux.Vars(r)["id"]
	userOtp := r.URL.Query().Get("otp")

	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	err := repo.VerifyUser(userID, userOtp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in verifying user ", err)
		res.StatusCode = http.StatusBadRequest
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Data = fmt.Sprintf("User with id %s has been verified", userID)

}


func (svc *AuthService) ResetPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &types.AuthResponse{}
	defer json.NewEncoder(w).Encode(res)

	var request struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Ensure email and new password are provided
	if request.Email == "" || request.NewPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.StatusCode = http.StatusBadRequest
		res.Error = "Email and new password are required"
		return
	}

	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	// Reset the user's password
	token, err := repo.ResetPassword(request.Email, request.NewPassword)
	if err != nil {
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			res.StatusCode = http.StatusNotFound
			res.Error = "User not found"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			res.StatusCode = http.StatusInternalServerError
			res.Error = err.Error()
		}
		log.Println("Error resetting password", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res.StatusCode = http.StatusOK
	res.Token = token
}
