package handlers

import (
	"net/http"

	"chat-app-golang/internal/database"
	"chat-app-golang/internal/models"
	"chat-app-golang/internal/websocket"
)

func ClearChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result := database.DB.Exec("DELETE FROM messages")
	if result.Error != nil {
		http.Error(w, "Gagal membersihkan chat", http.StatusInternalServerError)
		return
	}

	clearMessage := models.Message{
		Sender:  "System",
		Message: "Chat history has been cleared!",
	}
	websocket.Broadcast <- clearMessage

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chat cleared"))
}
