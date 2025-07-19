package handler

import (
	"io"
	"net/http"
)

func Fetch(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, resp.Body)
}