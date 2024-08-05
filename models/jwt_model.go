package models

import "github.com/dgrijalva/jwt-go"

var JwtKey []byte

type Claims struct {
	FirstName  string
	FamilyName string
	Email      string
	jwt.StandardClaims
}
