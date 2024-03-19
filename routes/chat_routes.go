package routes

import (
	"astra-ai-server/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.GetAllChats).Methods(http.MethodGet)

}
