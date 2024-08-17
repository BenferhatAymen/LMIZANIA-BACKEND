package jwthelper

import (
	"lmizania/models"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type JWTHelper struct {
	Claims models.Claims
}

var SECRET_KEY = []byte("lmizaniayajdk")

func (j *JWTHelper) GenerateJWT(claims models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func (j *JWTHelper) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
