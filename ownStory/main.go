package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	storyData "ownStory/sotryData"
	"strings"
)

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryChapter struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

func main() {
	storyFile := flag.String("f", "story.json", "Select story file")

	flag.Parse()

	data, err := os.ReadFile(*storyFile)
	if err != nil {
		panic(err)
	}

	story, err := storyData.NewStory(data)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc := strings.ReplaceAll(r.URL.Path, "/", "")
		if arc == "" {
			arc = "intro"
		}
		templ, err := template.New("template.html").ParseFiles("template.html")
		if err != nil {
			panic(err)
		}

		templ.Execute(w, story[arc])
	})

	log.Println("Listening...")

	http.ListenAndServe("localhost:8000", nil)
}
