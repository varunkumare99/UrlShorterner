package main

import (
	"fmt"
	"net/http"

	"github.com/varunkumare99/UrlShorterner/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/yt":             "https://youtube.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
- path: /gg
  url: https://google.com
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonData := `[
  {"path": "/docs", "url": "https://go.dev/doc/"},
  {"path": "/blog", "url": "https://go.dev/blog/"},
  {"path": "/tour", "url": "https://go.dev/tour/"}
]`
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonData), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
