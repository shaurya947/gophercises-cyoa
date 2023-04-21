package main

import (
	"fmt"
	"gophercises-cyoa/story"
	"os"
)

func main() {
	jsonBlob, err := os.ReadFile("gopher.json")
	if err != nil {
		exit(fmt.Sprintln("Error reading file:", err))
	}

	rootStoryMap, err := story.ParseJsonBlob(jsonBlob)
	if err != nil {
		exit(fmt.Sprintln("Error parsing json:", err))
	}

	fmt.Printf("%+v", rootStoryMap)
}

func exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
