package handler

import "net/http"

// Home is the handler for the root path ("/")
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/typer.html")
}