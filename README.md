# http-streamer

Minimal Go server that streams terminal or HTTP output to the browser using Server-Sent Events (SSE).

## Features

- Streams HTTP response from `https://httpbin.org/html` line by line
- Uses raw `net/http`, `http.Get`, `bufio.Scanner`, and `http.Flusher`
- HTML frontend uses native EventSource API for real-time display
- Static styles served from `/static/styles.css`
- No frameworks, no magic — just standard lib

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

## TODO

- Replace hardcoded URL with user-defined query param (/events?target=...)

- Add web form to submit URL or command from frontend

- Support ANSI-colored output (parse and style in browser)

- Add client disconnect/retry logic
---

## License

MIT

---