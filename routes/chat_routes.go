package routes

import (
	"astra-ai-server/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(router *mux.Router) {

	router.HandleFunc("/", controllers.WelcomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/chats", controllers.GetAllChats).Methods(http.MethodGet)
	router.HandleFunc("/chat", controllers.CreateNewChat).Methods(http.MethodPost)

}
