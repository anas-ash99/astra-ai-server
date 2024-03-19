package main

import (
	"astra-ai-server/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)
	err := http.ListenAndServe("0.0.0.0:"+port, router)
	if err != nil {
		fmt.Print("Serer running on port: " + port)
	}
}
