package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	domain "github.com/edgarSucre/bochinche/domain"
)

type Server struct {
	Router       mux.Router
	repo         domain.ChatRepository
	sessionStore sessions.Store
}

func New(repo domain.ChatRepository) *Server {
	sessionKey := os.Getenv("SESSION_KEY")
	return &Server{
		Router:       *mux.NewRouter(),
		repo:         repo,
		sessionStore: sessions.NewCookieStore([]byte(sessionKey)),
	}
}

func (s *Server) Start(mqClient domain.Broker) error {
	s.Router.HandleFunc("/rooms", s.CreateRoomHandler).Methods("POST")
	s.Router.HandleFunc("/rooms", s.ListRoomsHandler).Methods("GET")
	s.Router.HandleFunc("/chatter", s.RegisterChatterHandler).Methods("POST")
	s.Router.HandleFunc("/login", s.CreateSessionHandler).Methods("POST")
	s.Router.HandleFunc("/logout", s.DestroySessionHandler).Methods("POST")

	responseConsumer, err := mqClient.GetResponseConsummer()
	if err != nil {
		return err
	}

	hub := newHub(s.repo, responseConsumer, mqClient.PublishQuoteRequest)
	//hub := newHub(s.repo, mqClient.PublishQuoteRequest)
	go hub.run()

	s.Router.HandleFunc("/ws/{roomID}", func(w http.ResponseWriter, r *http.Request) {
		s.serveWs(hub, w, r)
	})

	log.Println("listing on port 8080")

	return http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)(&s.Router))
}
