package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sym3tri/hello/server"
)

func main() {
	fs := flag.NewFlagSet("hello", flag.ExitOnError)
	listen := fs.String("listen", "0.0.0.0:8080", "address/port to listen on")
	message := fs.String("message", "", "a message to print")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

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
