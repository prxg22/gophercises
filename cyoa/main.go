package main

import (
	"log"
	"net/http"

	"github.com/prxg22/gophercises/cyoa/handler"
	"github.com/prxg22/gophercises/cyoa/story"
)

func main() {
	st, parseErr := story.ParseFile("goopher.json")

	if parseErr != nil {
		log.Fatal(parseErr)
	}

	handler := &handler.StoryHandler{Story: st}

	log.Println("Starting server on :8080...")

	listeningErr := http.ListenAndServe(":8080", handler)

	if listeningErr != nil {
		log.Fatal(listeningErr)
	}

}
