package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	message string
	port    int
)

func init() {
	flag.StringVar(&message, "message", "there", "a message to print")
	flag.IntVar(&port, "port", 8080, "port to listen on")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", message)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
