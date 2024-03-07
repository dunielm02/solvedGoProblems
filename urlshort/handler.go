package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if value, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, value, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)

		defer fmt.Println(path)
	}
}

func YAMLHandler(fileData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []Route
	yaml.Unmarshal(fileData, &routes)

	return MapHandler(pathsToUrls(routes), fallback), nil
}

func JsonHandler(fileData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var routes []Route
	json.Unmarshal(fileData, &routes)

	return MapHandler(pathsToUrls(routes), fallback), nil
}

func pathsToUrls(routes []Route) map[string]string {
	pathsToUrls := make(map[string]string)

	for _, route := range routes {
		pathsToUrls[route.Path] = route.Url
	}

	return pathsToUrls
}
