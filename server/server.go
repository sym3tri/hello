package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	mux.Handle("/version", s.versionHandler())
	mux.Handle("/prestop", s.prestopHandler())
	mux.Handle("/poststart", s.poststartHandler())
	mux.Handle("/mount/", s.mountHandler())
	mux.Handle("/environment", s.environmentHandler())
	return http.Handler(mux)
}

func (s *Server) Version() string {
	return version
}

func (s *Server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, message: %s\n", s.Config.Message)
	}
}

func (s *Server) versionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: %s\n", version)
	}
}

func (s *Server) poststartHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("post-start")
		fmt.Fprint(w, "post-start")
	}
}

func (s *Server) environmentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("environment")
		fmt.Fprint(w, strings.Join(os.Environ(), "\n"))
	}
}

func (s *Server) prestopHandler() http.HandlerFunc {
	waitSecs := 5
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "initiated pre-stop\n")
		log.Println("initiated pre-stop")
		for i := waitSecs; i > 0; i-- {
			log.Printf("shutting down in: %ds", i)
			fmt.Fprintf(w, "shutting down in: %ds\n", i)
			time.Sleep(time.Second * 1)
		}
		log.Println("pre-stop complete")
		fmt.Fprint(w, "pre-stop complete")
	}
}

// mountHandler Prints the contents of the file specified in the 'file' query param.
func (s *Server) mountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		filename := q.Get("file")
		if filename == "" {
			fmt.Fprintf(w, "No file param specified")
			return
		}

		b, err := ioutil.ReadFile(filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "file not mounted: %s\n", filename)
			return
		}
		fmt.Fprintf(w, "file contents for: %s\n", filename)
		w.Write(b)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "not found")
}
