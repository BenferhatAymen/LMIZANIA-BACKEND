package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"context"
	"lmizania/models"

	"github.com/dgrijalva/jwt-go"
)

func LoginRequired(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"message":"missing token"}`, http.StatusUnauthorized)
			return
		}

		args := strings.Split(authHeader, " ")
		if len(args) < 2 {

			http.Error(w, `{"message":"invalid token format"}`, http.StatusUnauthorized)
			return
		}

		tokenStr := args[1]
		fmt.Println(tokenStr)
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})
		if err != nil {
			fmt.Println(err)
			http.Error(w, `{"message":"invalid token"}`, http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			fmt.Println(err)

			http.Error(w, `{"message":"token is not valid"}`, http.StatusUnauthorized)
			return
		}

		// Extract user ID from custom claims
		id := claims.ID
	
		fmt.Println(id)

		// Add userID to the request context
		ctx := context.WithValue(r.Context(), "userID", id)
		r = r.WithContext(ctx)

		// Call the next handler
		f(w, r)
	}
}
