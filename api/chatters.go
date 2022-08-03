package api

import (
	"encoding/json"
	"net/http"
)

//TODO: add validation
type RegisterChatterRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *Server) RegisterChatterHandler(w http.ResponseWriter, r *http.Request) {
	var request RegisterChatterRequest

	//TODO: verify content type, verify the body object fir unknow fields,
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Couldn't read the room name")
	}

}
