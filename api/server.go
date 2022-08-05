package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	repo "github.com/edgarSucre/bochinche/repository"
)

type Server struct {
	Router       mux.Router
	repo         repo.ChatRepository
	sessionStore sessions.Store
}

func New(repo repo.ChatRepository) *Server {
	return &Server{
		Router:       *mux.NewRouter(),
		repo:         repo,
		sessionStore: sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))),
	}
}

func (s *Server) Start() error {
	s.Router.HandleFunc("/rooms", s.CreateRoomHandler).Methods("POST")
	s.Router.HandleFunc("/rooms", s.ListRoomsHandler).Methods("GET")
	s.Router.HandleFunc("/chatter", s.RegisterChatterHandler).Methods("POST")
	s.Router.HandleFunc("/login", s.CreateSessionHandler).Methods("POST")
	s.Router.HandleFunc("/logout", s.DestroySessionHandler).Methods("DELETE")

	log.Println("listing on port 8080")

	return http.ListenAndServe(":8080", handlers.CORS()(&s.Router))
}
