package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edgarSucre/bochinche/domain"
)

func (s *Server) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request domain.RoomParams

	//TODO: verify content type, verify the body object fir unknow fields,
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Couldn't read the room name")
		return
	}

	room, err := s.repo.CreateRoom(ctx, request.Name)
	//TODO have the errors reside in the domain
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Couldn't create the room")
		return
	}

	data, err := json.Marshal(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, err.Error())
		return
	}

	setResponse(w, data)
}

func (s *Server) ListRoomsHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := s.sessionStore.Get(r, "chat-session")
	var user string
	if username, ok := session.Values["username"]; ok {
		user = username.(string)
	}

	if user == "" {
		w.WriteHeader(http.StatusForbidden)
		setErrorResponse(w, "Invalid chat session")
		return
	}

	ctx := r.Context()

	rooms, err := s.repo.ListRooms(ctx)

	//TODO have the errors reside in the domain
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Couldn't retrieve the room list")
		return
	}

	data, err := json.Marshal(rooms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, err.Error())
		return
	}

	setResponse(w, data)
}

//TODO: find a better place for utility functions
func setErrorResponse(w http.ResponseWriter, err string) {
	fmt.Fprintf(w, `{err: %s}`, err)
}

func setResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
