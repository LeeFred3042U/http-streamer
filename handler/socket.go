package handler

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func SocketStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// [ADDED] manually dial TCP socket
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Fprintf(w, "data: [dial error] %s\n\n", err.Error())
		flusher.Flush()
		return
	}
	defer conn.Close()

	// [ADDED] write raw HTTP request
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: google.com\r\nConnection: close\r\n\r\n")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintf(w, "data: %s\n\n", scanner.Text())
		flusher.Flush()
	}
}