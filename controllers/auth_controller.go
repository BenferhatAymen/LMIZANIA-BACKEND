package controllers

import (
	"encoding/json"
	"lmizania/models"
	"lmizania/repository"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	MongoCollection *mongo.Collection
}
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Token string      `json:"token,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body ", err)
		res.Error = err.Error()
		return
	}
	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	token, err := repo.UserLogin(user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in login ", err)
		res.Error = err.Error()
		return
	}
	w.WriteHeader(http.StatusOK)
	res.Token = token
}
func (svc *AuthService) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = "Invalid request payload"
		return
	}

	// Ensure email and password are provided
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Email and password are required"
		return
	}

	repo := repository.AuthRepo{MongoCollection: svc.MongoCollection}

	// Register the new user
	result, token, err := repo.RegisterUser(&user)
	if err != nil {
		if err.Error() == "user already exists" {
			w.WriteHeader(http.StatusConflict)
			res.Error = "User already exists"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			res.Error = err.Error()
		}
		log.Println("Error registering user", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res.Data = result
	res.Token = token
}
