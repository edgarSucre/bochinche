package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/edgarSucre/bochinche/domain"
)

func (s *Server) RegisterChatterHandler(w http.ResponseWriter, r *http.Request) {
	var request domain.ChatterParams

	//TODO: verify content type, verify the body object fir unknow fields,
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Couldn't read the chatter parameters")
		return
	}

	err = s.repo.RegisterChatter(r.Context(), request)
	if err != nil {
		if errors.Is(err, domain.ErrChatterConflict) {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		setErrorResponse(w, fmt.Sprintf("Couldn't save the chatter, %s", err.Error()))
		return
	}

	setResponse(w, nil)

}

func (s *Server) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	var params domain.VerifyChatterParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setErrorResponse(w, "Couldn't read the login parameters")
		return
	}

	err = s.repo.AreCredentialsValid(r.Context(), params)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		setErrorResponse(w, err.Error())
		return
	}

	session, _ := s.sessionStore.New(r, "chat-session")
	session.Values["username"] = params.UserName
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		setErrorResponse(w, "Could not authenticate the chatter")
	}
}

func (s *Server) DestroySessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := s.sessionStore.Get(r, "chat-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	setResponse(w, nil)
}
