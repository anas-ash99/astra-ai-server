package main

import (
	"astra-ai-server/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Print("Serer running on port: " + port)
	}
}
