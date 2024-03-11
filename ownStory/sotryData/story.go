package storyData

import "encoding/json"

type Story map[string]StoryChapter

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryChapter struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

func NewStory(jsonData []byte) (Story, error) {
	var story Story
	err := json.Unmarshal(jsonData, &story)

	if err != nil {
		return nil, err
	}

	return story, nil
}
