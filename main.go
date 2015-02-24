package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "v2.0.0"

var (
	message string
	address string
)

func init() {
	flag.StringVar(&message, "message", "there", "a message to print")
	flag.StringVar(&address, "address", "0.0.0.0:8080", "address/port to listen on")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s.\n", message)
		fmt.Fprintf(w, "Version: %s\n", version)
	})
	log.Printf("listening on %s...", address)
	log.Printf("Version: %s", version)
	log.Fatal(http.ListenAndServe(address, nil))
}
