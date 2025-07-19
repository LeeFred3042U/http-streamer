package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"http-streamer/handler" // [ADDED]
)

var addr = flag.String("addr", ":3000", "address")

func main() {
	flag.Parse()

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/events", handler.Events)
	http.HandleFunc("/fetch", handler.Fetch)
	http.HandleFunc("/socket", handler.SocketStream) // [ADDED]

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
