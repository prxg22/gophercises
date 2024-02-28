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
	file, readErr := os.ReadFile(path)

	if readErr != nil {
		return nil, fmt.Errorf("unable read file on path: %v. %v", path, readErr)
	}

	j, parseErr := ParseJSON(file)

	if parseErr != nil {
		return nil, parseErr
	}

	return j, nil
}

func ParseJSON(j []byte) (*Story, error) {
	story := &Story{}

	parseErr := json.Unmarshal(j, story)

	if parseErr != nil {
		return nil, fmt.Errorf("unable to parse JSON: %v", parseErr)
	}

	return story, nil
}

func (s *Story) GetArc(key string) *Arc {
	arc := (*s)[key]

	if arc.Title == "" {
		return nil
	}

	return &arc
}

func (s *Story) Intro() *Arc {
	return s.GetArc("intro")
}

func (s *Story) Marshal() ([]byte, error) {
	return json.Marshal(*s)
}

func (a *Arc) Marshal() ([]byte, error) {
	return json.Marshal(*a)
}
