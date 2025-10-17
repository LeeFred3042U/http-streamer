package handler

import "net/http"

// NewRouter creates and configures a new server router
func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/", Home)
    mux.HandleFunc("/events", Events)
    mux.HandleFunc("/fetch", Fetch)
    mux.HandleFunc("/socket", SocketStream)
    mux.HandleFunc("/ws", WebSocketHandler)

    // Serve static files
    staticHandler := http.FileServer(http.Dir("static"))
    mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

    return mux
}