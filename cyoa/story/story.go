package story

import (
	"encoding/json"
	"fmt"
	"os"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Arc

func ParseFile(path string) (*Story, error) {
	story := make(Story)
	file, readErr := os.ReadFile(path)

	if readErr != nil {
		return nil, fmt.Errorf("unable read file on path: %v. %v", path, readErr)
	}

	parseErr := json.Unmarshal(file, &story)

	if parseErr != nil {
		return nil, fmt.Errorf("unable to parse JSON: %v", parseErr)
	}

	return &story, nil
}

func (s *Story) Arcs() []Arc {
	arcs := []Arc{}
	i := 0
	for _, v := range *s {
		arcs[i] = v
		i++
	}

	return arcs
}
