package models

import (
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("lmizaniayajdk")

type Claims struct {
	ID         string
	FirstName  string
	FamilyName string
	Email      string
	jwt.StandardClaims
}
