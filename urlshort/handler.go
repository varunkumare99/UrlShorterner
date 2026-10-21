package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.yaml.in/yaml/v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range pathsToUrls {
			fmt.Printf("url = %v, MapHandler-> k = %v, v = %v\n", r.URL.Path, k, v)
		}
		val, ok := pathsToUrls[r.URL.Path]
		if ok {
			fmt.Println("Present")
			http.Redirect(w, r, val, http.StatusTemporaryRedirect)
			return
		}
		fmt.Println("fallback")
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type path2Url struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url`
}

func parseYAML(yml []byte) ([]path2Url, error) {
	yD := []path2Url{}
	err := yaml.Unmarshal(yml, &yD)
	return yD, err
}

func parseJson(jsonByte []byte) ([]path2Url, error) {
	jD := []path2Url{}
	err := json.Unmarshal(jsonByte, &jD)
	return jD, err
}

func buildMap(yD []path2Url) map[string]string {
	path2UrlMap := make(map[string]string)
	for _, record := range yD {
		path2UrlMap[record.Path] = record.URL
	}
	return path2UrlMap
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(jsonByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJson(jsonByte)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
}
