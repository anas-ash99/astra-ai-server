package main

import (
	"astra-ai-server/controllers"
	"fmt"

	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/get-user", controllers.GetAllChats)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print("Serer running on port: " + port)
	}
}
