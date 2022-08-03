package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//TODO: add validation

type CreateRoomRequest struct {
	Name string `json:"name"`
}

//TODO: use this
type RoomResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Server) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request CreateRoomRequest

	//TODO: verify content type, verify the body object fir unknow fields,
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Couldn't read the room name")
	}

	room, err := s.repo.CreateRoom(ctx, request.Name)
	//TODO have the errors reside in the domain
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Couldn't create the room")
	}

	data, err := json.Marshal(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, err.Error())
	}

	setResponse(w, data)
}

func (s *Server) ListRoomsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rooms, err := s.repo.ListRooms(ctx)

	//TODO have the errors reside in the domain
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Couldn't retrieve the room list")
	}

	data, err := json.Marshal(rooms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, err.Error())
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
