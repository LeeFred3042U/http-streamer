package handler

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

// Events is a handler that fetches an external URL and 
// streams its content line-by-line to the client
func Events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	lines := make(chan string)

	go func() {
		defer close(lines)

		resp, err := http.Get("https://httpbin.org/html")
		if err != nil {
			lines <- fmt.Sprintf("[error] %s", err.Error())
			close(lines)
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			// send each scaned line to channel
			lines <- scanner.Text()
			time.Sleep(100 * time.Millisecond)
		}
		if err := scanner.Err(); err != nil {
			lines <- fmt.Sprintf("[error] %s", err.Error())
		}
		close(lines)
	}()

	// loop and read from channel
	for line := range lines {
		fmt.Fprintf(w, "data: %s\n\n", line)
		flusher.Flush()
	}
}