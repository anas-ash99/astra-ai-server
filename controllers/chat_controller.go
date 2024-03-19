package controllers

import (
	"astra-ai-server/models"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func GetAllChats(w http.ResponseWriter, r *http.Request) {
	var chats []models.Chat
	chats = append(chats, models.Chat{ID: "1", Title: "First chat", LastMsgTimestamp: "02103210"})
	chats = append(chats, models.Chat{ID: "2", Title: "Second chat", LastMsgTimestamp: "02132130"})

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(chats)
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
	}

}

func CreateNewChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var newChat models.Chat
	err := json.NewDecoder(r.Body).Decode(&newChat)
	if err != nil {
		http.Error(w, "Error decoding user data", http.StatusBadRequest)
		return
	}
	newChat.ID = uuid.New().String()
	// Respond with success (usually a 201 Created status)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newChat)
}
