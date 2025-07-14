package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"os/exec"
	"bufio"
)


var addr *string = flag.String("addr", ":3000", "address")

func main(){
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.ListenAndServe(*addr, nil)
}


func events(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}	

	cmd := exec.Command("ping", "google.com")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		http.Error(w, "Could not get stdout", http.StatusInternalServerError)
		return
	}
	scanner := bufio.NewScanner(stdout)

	cmd.Start()
	defer cmd.Process.Kill()

	fmt.Fprintf(w, ": connected to ping stream\n\n")
	flusher.Flush()


	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(w, "data: %s\n\n", line)
		flusher.Flush()
		time.Sleep(time.Millisecond * 100) 
	}

	if err := scanner.Err(); err != nil {
    	fmt.Fprintf(w, "data: [error] %s\n\n", err.Error())
	}
	

}


func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "typer.html")
}