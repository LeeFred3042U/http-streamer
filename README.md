# http-streamer

Minimal Go server that streams terminal command output to the browser using Server-Sent Events (SSE).

## Features

- Runs a shell command (`ping google.com`) and streams stdout to clients
- Uses raw `net/http` and `text/event-stream` (no frameworks, no magic)
- HTML frontend uses native EventSource API for real-time display
- Built with `os/exec`, `bufio`, and `http.Flusher`

## Install

```bash
git clone https://github.com/LeeFred3042U/http-streamer
cd http-streamer
go run main.go
```
---
## Usage
- Visit http://localhost:3000

- You’ll see ping output streamed live as data: events
- Frontend appends each message to <p id="resp">...</p>
```go
fmt.Fprintf(w, "data: %s\n\n", line)
```
---
## TODO
- Replace ping with any arbitrary command (e.g. traceroute, curl, top)

- Replace ping with user-defined command input

- Add web form to submit commands

- Support ANSI-colored output (parse and style in browser)

- Handle disconnects / retries client-side
---