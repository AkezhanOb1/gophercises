package parser

import (
	"os"
	"log"
	"encoding/json"
)

// Types for successfully parsing data
type Chapter struct{
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`

}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

// Parses given string into Chapter struct and returns map of this chapters
func ParseJsonFile(str string) map[string]Chapter {

	story := make(map[string]Chapter)
	file, err := os.Open(str)
	if err != nil {
		log.Fatalln(err, "Error white opening the project")
	}
	defer file.Close()

	reader := json.NewDecoder(file)
	reader.Decode(&story)

	return story
}


