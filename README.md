# GoProxy

> A terminal-first networking toolkit written in Go.

GoProxy combines an HTTP client and an HTTP forward proxy into a single CLI application. The project focuses on understanding HTTP from first principles by implementing request construction, forwarding, proxying, and response handling directly using Go's standard library.

---

## Features

### HTTP Client

- Custom HTTP methods (GET, POST, PUT, DELETE, ...)
- Query parameter support
- Custom headers
- Request body support
- Configurable timeout
- Pretty JSON formatting
- Save responses to disk
- Request summaries

### HTTP Forward Proxy

- Dynamic request forwarding
- Request/Response forwarding
- Header propagation
- Request body forwarding
- Response body forwarding
- Status code forwarding
- Hop-by-hop header filtering
- Built using Go's `net/http`

---

## Architecture

<p align="center">
  <img src="https://github.com/user-attachments/assets/9e41cbf4-0168-4cb9-97bb-8333ed8159c0" width="900">
</p>

### Project Structure

```text
goProxy/
│
├── cmd/                     # Cobra commands
│   ├── inspect.go
│   ├── proxy.go
│   ├── root.go
│   └── version.go
│
├── internal/
│   ├── httpclient/          # HTTP client implementation
│   └── tui/                 # Bubble Tea interface
│
├── main.go
└── README.md
```

---

## Tech Stack

- Go
- Cobra
- Bubble Tea
- Lip Gloss
- net/http

---

## Running

Clone the repository

```bash
git clone https://github.com/Suryansh-singh-137/goProxy.git
cd goProxy
```

Install dependencies

```bash
go mod tidy
```

Run

```bash
go run .
```

Start the proxy

```bash
go run . proxy start
```

---

## Current Capabilities

- HTTP Client
- HTTP Forward Proxy
- Cobra-based CLI
- Modular HTTP package
- Bubble Tea powered interface
- Hop-by-hop header filtering

---

## Roadmap

- HTTPS CONNECT tunneling
- Interactive request builder
- Proxy dashboard
- Request history
- Configuration profiles

---

## Why this project?

Most HTTP tools abstract networking behind a polished interface.

GoProxy takes the opposite approach.

The goal is to understand how HTTP works by implementing the networking stack yourself—from building requests and forwarding traffic to writing a forward proxy and eventually wrapping everything inside a terminal interface.
