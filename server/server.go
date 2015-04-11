package server

import (
	"fmt"
	"net/http"
)

const version = "v2.0.0"

type Config struct {
	Message string
}

type Server struct {
	Config Config
}

func (s *Server) HTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", s.rootHandler())
	return http.Handler(mux)
}

func (s *Server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s.\n", s.Config.Message)
		fmt.Fprintf(w, "Version: %s\n", version)
	}
}

func (s *Server) Version() string {
	return version
}
