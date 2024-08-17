package repository

import (
	"context"
	"errors"
	"lmizania/models"
	"lmizania/pkg/jwthelper"
	"lmizania/pkg/mail"

	"lmizania/pkg/passwordhelper"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	MongoCollection *mongo.Collection
}

func (r *AuthRepo) RegisterUser(user *models.User) (interface{}, string, error) {
	PasswordHelper := passwordhelper.PasswordHelper{}

	// Check if the user already exists
	var existingUser models.User
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return nil, "", errors.New("user already exists")
	}

	// Hash the user's password
	hash, err := PasswordHelper.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
		return nil, "", err
	}
	user.Password = hash
	user.ID = uuid.New().String()
	user.IsVerified = false

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
		ID:             user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
	}
	verifier := mail.NewVerifier()
	err = verifier.SendVerfication(user.ID, []string{user.Email})
	if err != nil {
		return nil, "", err
	}
	jwtHelper := jwthelper.JWTHelper{Claims: claims}

	token, err := jwtHelper.GenerateJWT(claims)
	if err != nil {
		return nil, "", err
	}

	return result, token, nil
}
func (r *AuthRepo) UserLogin(email, password string) (string, error) {
	var user models.User
	PasswordHelper := passwordhelper.PasswordHelper{}

	// Find user by email
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	// Validate password
	err = PasswordHelper.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	claims := models.Claims{
		FirstName:      user.FirstName,
		FamilyName:     user.FamilyName,
		Email:          user.Email,
		ID:             user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
	}
	jwtHelper := jwthelper.JWTHelper{Claims: claims}
	token, err := jwtHelper.GenerateJWT(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *AuthRepo) VerifyUser(userID, otp string) error {
	var user models.User
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return err
	}
	verifier := mail.NewVerifier()
	err = verifier.Verify(userID, otp)
	if err != nil {
		return err
	}
	_, err = r.MongoCollection.UpdateOne(context.Background(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"is_verified": true}})
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) ResetPassword(email, newPassword string) ( string, error) {
	PasswordHelper := passwordhelper.PasswordHelper{}

	var user models.User
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	hash, err := PasswordHelper.HashPassword(newPassword)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	user.Password = hash

	update := bson.M{
		"$set": bson.M{"password": user.Password},
	}
	_, err = r.MongoCollection.UpdateOne(context.Background(), bson.M{"email": email}, update)
	if err != nil {
		return  "", err
	}

	// Create JWT token
	claims := models.Claims{
		FirstName:      user.FirstName,
		FamilyName:     user.FamilyName,
		Email:          user.Email,
		ID:             user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
	}

	jwtHelper := jwthelper.JWTHelper{Claims: claims}
	token, err := jwtHelper.GenerateJWT(claims)
	if err != nil {
		return  "", err
	}

	return token, nil
}
