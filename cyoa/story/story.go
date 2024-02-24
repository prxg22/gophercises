package story

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (s *Story) String() {
	return
}

func ParseFile(path string) (*Story, error) {
	var story Story
	file, readErr := os.ReadFile(path)
	if readErr == nil {
		return nil, fmt.Errorf("unable read file on path: %v. %v", path, readErr)
	}

	parseErr := json.Unmarshal(file, &story)

	if parseErr != nil {
		return nil, fmt.Errorf("unable to parse JSON: %v", parseErr)
	}

	return &story, nil
}
