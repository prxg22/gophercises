package main

import (
	"github/prxg22/gophercises/cyoa/story"
	"log"
)

func main() {
	st, err := story.ParseFile("./goopher.json")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(st)
}
