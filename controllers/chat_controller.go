package controllers

import (
	"astra-ai-server/models"
	"encoding/json"
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
