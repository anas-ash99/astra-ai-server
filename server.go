package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}
	return server
}
