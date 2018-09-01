package router

import (
	"net/http"
	. "Go-Web/go-exercise/storyJsonParser/parser"
	. "Go-Web/go-exercise/storyJsonParser/config"
			)

var story map[string]Chapter
func init() {
	story = ParseJsonFile("story.json")
}

func Intro (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("intro", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}

func NewYork (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("new-york", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}

func Debate (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("debate", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}

func SeanKelly (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("sean-kelly", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}

func MarkBates (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("mark-bates", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}

func Denver (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("denver", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}


func Home (w http.ResponseWriter, r *http.Request) {
	introChapter := findChapter("home", story)
	Tpl.ExecuteTemplate(w, "tpl.gohtml", introChapter)
}



func findChapter(chapter string, book map[string]Chapter) Chapter {
	if val, ok := book[chapter]; ok {
		return val
	}
	return Chapter{}
}

