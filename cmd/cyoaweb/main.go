package main

import (
	"flag"
	"fmt"
	"gophercises-cyoa/story"
	"html/template"
	"net/http"
	"os"
)

func main() {
	filenamePtr := flag.String("file", "gopher.json", "JSON file containing full story")
	flag.Parse()

	fmt.Println("Using the story in", *filenamePtr)
	jsonBlob, err := os.ReadFile(*filenamePtr)
	if err != nil {
		exit(fmt.Sprintln("Error reading file", *filenamePtr, err))
	}

	story, err := story.ParseJsonBlob(jsonBlob)
	if err != nil {
		exit(fmt.Sprintln("Error parsing json:", err))
	}

	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	handler := StoryWebHandler{story: story, tmpl: tmpl}
	http.ListenAndServe(":8080", &handler)
}

func exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
