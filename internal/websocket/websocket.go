package websocket

import (
	"log"
	"net/http"
	"strings"

	"chat-app-golang/internal/auth"
	"chat-app-golang/internal/database"
	"chat-app-golang/internal/models"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	Clients   = make(map[*websocket.Conn]string)
	Broadcast = make(chan models.Message)
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		http.Error(w, "Token tidak ditemukan", http.StatusUnauthorized)
		return
	}

	username, err := auth.GetUsernameFromToken(tokenString)
	if err != nil {
		http.Error(w, "Token tidak valid: "+err.Error(), http.StatusUnauthorized)
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	Clients[conn] = username

	welcomeMsg := models.Message{
		Sender:  "System",
		Message: "Selamat datang, " + username + "!",
	}
	conn.WriteJSON(welcomeMsg)

	joinMsg := models.Message{
		Sender:  "System",
		Message: username + " telah bergabung!",
	}
	Broadcast <- joinMsg

	onlineUsers := getOnlineUsers()
	onlineMsg := models.Message{
		Sender:  "System",
		Message: "UPDATE_ONLINE_USERS:" + strings.Join(onlineUsers, ","),
	}
	log.Println("Mengirim daftar pengguna online ke", username, ":", onlineMsg.Message)
	if err := conn.WriteJSON(onlineMsg); err != nil {
		log.Println("Gagal mengirim daftar pengguna online:", err)
	}

	var messages []models.Message
	database.DB.Order("timestamp asc").Find(&messages)
	for _, msg := range messages {
		conn.WriteJSON(msg)
	}

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Client disconnected:", err)
			delete(Clients, conn)

			leaveMsg := models.Message{
				Sender:  "System",
				Message: username + " telah keluar!",
			}
			Broadcast <- leaveMsg

			onlineUsers = getOnlineUsers()
			onlineUpdateMsg := models.Message{
				Sender:  "System",
				Message: "UPDATE_ONLINE_USERS:" + strings.Join(onlineUsers, ","),
			}
			log.Println("Mengirim pembaruan daftar pengguna online:", onlineUpdateMsg.Message)
			Broadcast <- onlineUpdateMsg
			break
		}

		msg.Sender = username
		database.DB.Create(&msg)
		Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-Broadcast
		log.Println("Broadcasting pesan:", msg.Message)
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error sending message:", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}

func getOnlineUsers() []string {
	var users []string
	for _, username := range Clients {
		if !contains(users, username) {
			users = append(users, username)
		}
	}
	return users
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
