package handler

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

// SocketStream manually makes an HTTP request 
// and streams the entire raw response
func SocketStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Flusher to send data incrementally
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Manually open a TCP connection to the server on port 80 (HTTP)
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Fprintf(w, "data: [dial error] %s\n\n", err.Error())
		flusher.Flush()
		return
	}
	defer conn.Close()

	// Manually write a raw HTTP GET request to the socket
	// This shows what an HTTP request looks like at the protocol level
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: google.com\r\nConnection: close\r\n\r\n")

	// Use a scanner to read the response from the server line by line
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Format and send each line of the raw response to the client
		fmt.Fprintf(w, "data: %s\n\n", scanner.Text())
		flusher.Flush()
	}
}