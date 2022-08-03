package api

import "net/http"

type Server struct {
	Router http.ServeMux
}
