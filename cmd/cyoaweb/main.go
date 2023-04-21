package main

import (
	"flag"
	"fmt"
	"gophercises-cyoa/story"
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

	fmt.Printf("%+v", story)
}

func exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
