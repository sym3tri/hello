package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sym3tri/hello/server"
)

func main() {
	fs := flag.NewFlagSet("hello", flag.ExitOnError)
	listen := fs.String("listen", "0.0.0.0:8080", "address/port to listen on")
	message := fs.String("message", "", "a message to print")

	cfg := server.Config{
		Message: *message,
	}

	srv := &server.Server{
		Config: cfg,
	}

	httpsrv := &http.Server{
		Addr:    *listen,
		Handler: srv.HTTPHandler(),
	}

	log.Printf("Binding to %s...", httpsrv.Addr)
	log.Printf("Version: %s", srv.Version())
	log.Fatal(httpsrv.ListenAndServe())
}
