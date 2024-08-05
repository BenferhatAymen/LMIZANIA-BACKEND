package repository

import (
	"context"
	"errors"
	"lmizania/helpers"
	"lmizania/models"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UserRepo) RegisterUser(user *models.User) (interface{}, string, error) {
	passwordHelper := helpers.PasswordHelper{}

	// Check if the user already exists
	var existingUser models.User
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return nil, "", errors.New("user already exists")
	}

	// Hash the user's password
	hash, err := passwordHelper.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
		return nil, "", err
	}
	user.Password = hash
	user.ID = uuid.New().String()

	// Insert the new user into the database
	result, err := r.MongoCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, "", err
	}

	// Create JWT token
	claims := models.Claims{
		FirstName:      user.FirstName,
		FamilyName:     user.FamilyName,
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
	}
	jwtHelper := helpers.JWTHelper{Claims: claims}
	token, err := jwtHelper.GenerateJWT(claims)
	if err != nil {
		return nil, "", err
	}

	return result, token, nil
}
func (r *UserRepo) UserLogin(email, password string) (string, error) {
	var user models.User
	passwordHelper := helpers.PasswordHelper{}

	// Find user by email
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	// Validate password
	err = passwordHelper.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	claims := models.Claims{
		FirstName:      user.FirstName,
		FamilyName:     user.FamilyName,
		Email:          user.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
	}
	jwtHelper := helpers.JWTHelper{Claims: claims}
	token, err := jwtHelper.GenerateJWT(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

// func main() {
// 	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://baroud:baroud123@cluster0.32co7ow.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mongoTestClient.Disconnect(context.Background())

// 	userRepo := UserRepo{MongoCollection: mongoTestClient.Database("test").Collection("users")}
// 	result, err := userRepo.RegisterUser(&models.User{FirstName: "Aymen", FamilyName: "Lmizania", Email: "aymen@gmail.com", Password: "aymen2004", ID: uuid.New().String()})

// 	fmt.Println(err)
// 	fmt.Print(result)

// }
