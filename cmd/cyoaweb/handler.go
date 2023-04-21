package main

import (
	"fmt"
	"gophercises-cyoa/story"
	"html/template"
	"net/http"
	"strings"
)

type StoryWebHandler struct {
	story story.Story
	tmpl  *template.Template
}

func (handler *StoryWebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	arc, found := handler.story[path]
	if !found {
		arc = handler.story["intro"]
	}

	if err := handler.tmpl.Execute(w, arc); err != nil {
		fmt.Println(err)
	}
}
