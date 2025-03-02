package main

import (
	"fmt"
	"log"
	"net/http"

	"chat-app-golang/internal/database"
	"chat-app-golang/internal/handlers"
	"chat-app-golang/internal/websocket"
)

func main() {
	// Inisialisasi database
	database.InitDB()

	// Jalankan handler untuk pesan WebSocket
	go websocket.HandleMessages()

	// Routing
	http.HandleFunc("/ws", websocket.HandleConnections)
	http.HandleFunc("/clear-chat", handlers.AuthMiddleware(handlers.ClearChatHandler))
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", handlers.AuthMiddleware(handlers.ProfileHandler))

	// Serve file statis dari folder views
	http.Handle("/", http.FileServer(http.Dir("./views")))

	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
