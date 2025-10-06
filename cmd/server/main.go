package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"http-streamer/handler"
)

// Declare a command-line flag for the server address with a default value
var addr = flag.String("addr", ":3000", "http server address")

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Create a new router from the handler package
	// centralizing all route definitions
	router := handler.NewRouter()

	fmt.Println("Server listening on", *addr)
	fmt.Println("Open http://localhost:3000 in your browser.")

	log.Fatal(http.ListenAndServe(*addr, router))
}