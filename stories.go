package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
)

// Story object representing one ticket from a sprint
type Story struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

// Save serializes story into JSON format and saves to writer
func (s *Story) Save(w io.Writer, data ACTrackerData) error {
	data.Stories = append(data.Stories, *s)
	return json.NewEncoder(w).Encode(data)
}

// AddStory command for creating a new story
func AddStory() {
	Clear[runtime.GOOS]()
	var story Story
	fmt.Print("Create a new story:\n\n")
	storyID := GetString("Story ID: ")
	storyName := GetString("Story name: ")
	storyDesc := GetString("Story description: ")
	story = Story{ID: storyID, Name: storyName, Description: storyDesc}
	WriteObjectToMemory(&story)
}

// ListStories list all stories saved in local file
func ListStories() {
	// tbi
}
