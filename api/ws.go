package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edgarSucre/bochinche/domain"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	session, _ := s.sessionStore.Get(r, "chat-session")
	var user string
	if username, ok := session.Values["username"]; ok {
		user = username.(string)
	}

	var (
		chatter  domain.Chatter
		roomName string
	)

	if user == "" {
		w.WriteHeader(http.StatusForbidden)
		setErrorResponse(w, "Invalid chat session")
		return
	}

	chatter, err := s.repo.GetChatter(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		setErrorResponse(w, "Invalid chat session")
		return
	}

	vars := mux.Vars(r)
	if room, ok := vars["roomID"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Invalid room")
		return
	} else {
		roomName = room
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Could not stablish chat connection")
		return
	}
	client := &Client{
		room:    roomName,
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte, 256),
		chatter: chatter,
	}

	//send chats when join a room
	chats, err := s.repo.ListChats(r.Context(), roomName)
	if err != nil {
		log.Printf("couldn't retrieve chats for room %s\n", roomName)
	}

	for _, v := range chats {
		client.send <- []byte(fmt.Sprintf("%s: %s", v.Author, v.Message))
	}

	chatRecord := chatterRecord{roomName, client}
	client.hub.register <- chatRecord

	go client.readMessages()
	go client.writeMessages()
}
