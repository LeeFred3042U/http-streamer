# http-streamer

Toy Go web server that streams live response data from Google.com using raw `net/http` and raw sockets -_-

## Features

- Streams HTML responses line-by-line via Server-Sent Events (SSE)
- Uses `http.Get` or raw TCP sockets (`net.Dial`) to fetch content
- Practiced concurrency with goroutines and channels
- Buffered streaming using `bufio.Scanner`
- Built entirely with raw `net/http` and zero frameworks
- Used Insomnia to test endpoints and monitor real-time responses

## Routes

| Route      | Description                              |
|------------|------------------------------------------|
| `/`        | Home page (`typer.html`)                 |
| `/events`  | Streams from `httpbin.org/html` (SSE)    |
| `/fetch`   | Raw HTTP stream from Google.com          |
| `/socket`  | Google stream via raw TCP socket         |

## Install

```bash
git clone https://github.com/LeeFred3042U/http-streamer
cd http-streamer
go run main.go
```

---

## Usage

- Visit http://localhost:3000

- You’ll see a clean stream of an HTTP response like but messier:

```html
<title>Herman Melville - Moby-Dick</title>
<p>Call me Ishmael...</p>
```

- Output scrolls in real time as each line is received from the response body

---
## FLOW: Request to Response
```rust
[Browser] ---- GET /events --> [Go Server]
                                |
                                +--> GET https://httpbin.org/html
                                        |
                                        +--> Read line-by-line (bufio.Scanner)
                                                |
                                                +--> format as SSE "data: line\n\n"
                                                         |
                                                         +--> Flush to client //Ya its readable kindoff <_>
```
---

## TODO

- Replace hardcoded URL with user-defined query param (`/events?target=...`)
- Add frontend form to submit URL or command
- Parse and display ANSI-colored terminal output
- Add SSE reconnect logic + client disconnect awareness
- Add `/events?cmd=...` support to stream shell command output

---

## LICENSE

MIT

---