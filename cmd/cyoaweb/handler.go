package main

import (
	"gophercises-cyoa/story"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type StoryWebHandler struct {
	story story.Story
	tmpl  *template.Template
}

func (handler *StoryWebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "intro"
	}

	arc, found := handler.story[path]
	if !found {
		http.Error(w, "Chapter not found", http.StatusNotFound)
		return
	}

	if err := handler.tmpl.Execute(w, arc); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}
