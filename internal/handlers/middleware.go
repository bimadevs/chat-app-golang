package handlers

import (
	"net/http"

	"chat-app-golang/internal/config"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Token tidak ditemukan!", http.StatusUnauthorized)
			return
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWTKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token tidak valid!", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
