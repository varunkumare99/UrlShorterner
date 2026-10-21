# URL Shortener

A simple URL redirection service built with Go's standard `net/http` package. It supports URL mappings defined in Go maps, YAML, and JSON, with fallback handlers for unmatched paths.

## Features

- **Map Handler:** Redirects URLs using a Go map.
- **YAML Handler:** Loads URL mappings from YAML data.
- **JSON Handler:** Loads URL mappings from JSON data.
- **Fallback Handler:** Returns a default response when no matching path is found.
- **HTTP Redirects:** Uses temporary redirects (`307`).

## Getting Started

### Prerequisites

- Go 1.23 or later

### Run the Server

```bash
git clone https://github.com/varunkumare99/UrlShorterner.git
cd UrlShorterner
go mod tidy
go run .
```

The server starts on `http://localhost:8080`.

## Example Routes

| Path | Redirects To |
|---|---|
| `/yt` | https://youtube.com |
| `/gg` | https://google.com |
| `/docs` | https://go.dev/doc/ |
| `/blog` | https://go.dev/blog/ |
| `/tour` | https://go.dev/tour/ |

Visit any of these paths in your browser to test the redirects.

## Technologies

- Go
- `net/http`
- YAML (`go.yaml.in/yaml/v3`)
- JSON (`encoding/json`)

## Learning Objectives

This project explores HTTP handlers, handler composition, URL redirection, JSON/YAML parsing, and fallback handling in Go.
