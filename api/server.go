package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	repo "github.com/edgarSucre/bochinche/repository"
)

type Server struct {
	Router mux.Router
	repo   repo.ChatRepository
}

func New(repo repo.ChatRepository) *Server {
	return &Server{
		Router: *mux.NewRouter(),
		repo:   repo,
	}
}

func (s *Server) Start() error {
	s.Router.HandleFunc("/rooms", s.CreateRoomHandler).Methods("POST")
	s.Router.HandleFunc("/rooms", s.ListRoomsHandler).Methods("GET")

	log.Println("listing on port 8080")

	return http.ListenAndServe(":8080", &s.Router)
}
