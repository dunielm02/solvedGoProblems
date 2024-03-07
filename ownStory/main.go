package main

import (
	"encoding/json"
	"flag"
	"os"
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

var story map[string]StoryChapter = make(map[string]StoryChapter)

func main() {
	storyFile := flag.String("f", "story.json", "Select story file")

	flag.Parse()

	data, err := os.ReadFile(*storyFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &story)
}
