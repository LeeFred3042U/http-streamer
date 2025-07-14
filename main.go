package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"bufio"
	"io"
	"log"
)


var addr *string = flag.String("addr", ":3000", "address")

func main(){
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.HandleFunc("/fetch", fetch)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	fmt.Println("Listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}


func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, ": connected to fetch stream\n\n")
	flusher.Flush()

	resp, err := http.Get("https://httpbin.org/html")
	if err != nil {
		fmt.Fprintf(w, "data: [error] %s\n\n", err.Error())
		flusher.Flush()
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(w, "data: %s\n\n", line)
		flusher.Flush()
		time.Sleep(100 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(w, "data: [error] %s\n\n", err.Error())
		flusher.Flush()
	}
}



func fetch(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, resp.Body) // raw stream, my eyes
}


func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "typer.html")
}