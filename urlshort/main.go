package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Route struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func main() {
	yamlFileName := flag.String("yaml", "data.yaml", "Select YAML file")
	useJson := flag.Bool("j", false, "Use Json")
	jsonFileName := flag.String("json", "data.json", "Select Json File")
	flag.Parse()

	mux := defaultMux()

	fileToRead := yamlFileName
	if *useJson {
		fileToRead = jsonFileName
	}

	fileData, err := os.ReadFile(*fileToRead)
	if err != nil {
		log.Fatal(err)
	}

	var handler http.HandlerFunc

	if *useJson {
		handler, err = JsonHandler(fileData, mux)
	} else {
		handler, err = YAMLHandler(fileData, mux)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
