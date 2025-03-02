package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"chat-app-golang/internal/auth"
	"chat-app-golang/internal/config"
	"chat-app-golang/internal/database"
	"chat-app-golang/internal/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username dan password tidak boleh kosong", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal memproses password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			http.Error(w, "Username sudah digunakan!", http.StatusBadRequest)
		} else {
			http.Error(w, "Gagal mendaftarkan user: "+result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User berhasil didaftarkan!"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if input.Username == "" || input.Password == "" {
		http.Error(w, "Username dan password tidak boleh kosong", http.StatusBadRequest)
		return
	}

	var user models.User
	result := database.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "Username atau password salah!", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token":    tokenString,
		"username": user.Username,
	})
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token tidak ditemukan!", http.StatusUnauthorized)
		return
	}

	username, err := auth.GetUsernameFromToken(tokenString)
	if err != nil {
		http.Error(w, "Token tidak valid: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"username": username})
}
