package handler

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func Events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// [ADDED] channel for line streaming
	lines := make(chan string)

	// [ADDED] Goroutine to fetch + scan lines concurrently
	go func() {
		resp, err := http.Get("https://httpbin.org/html")
		if err != nil {
			lines <- fmt.Sprintf("[error] %s", err.Error())
			close(lines)
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			lines <- scanner.Text()
			time.Sleep(100 * time.Millisecond)
		}
		if err := scanner.Err(); err != nil {
			lines <- fmt.Sprintf("[error] %s", err.Error())
		}
		close(lines)
	}()

	// [MODIFIED] send from channel to client
	for line := range lines {
		fmt.Fprintf(w, "data: %s\n\n", line)
		flusher.Flush()
	}
}