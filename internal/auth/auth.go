package auth

import (
	"fmt"

	"chat-app-golang/internal/config"

	"github.com/golang-jwt/jwt"
)

func GetUsernameFromToken(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWTKey, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token: %v", err)
	}
	username, ok := (*claims)["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in token")
	}
	return username, nil
}
