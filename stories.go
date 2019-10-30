package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Story object representing one ticket from a sprint
type Story struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (s Story) String() string {
	id := aurora.Bold(aurora.Yellow(fmt.Sprintf("[%v]", s.ID)))
	name := aurora.Bold(fmt.Sprintf("%v", strings.Title(s.Name)))
	return fmt.Sprintf("%v - %v", id, name)
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
	Clear[runtime.GOOS]()
	data, err := GetStoredACTrackingData()
	GeneralErr(err, "ListStories Error: Failed to read objects from memory.")
	fmt.Println(aurora.Bold(aurora.Cyan("Showing all saved stories\n")))
	// fmt.Println(aurora.Underline(aurora.Green("Sprint 13")))
	for _, story := range data.Stories {
		fmt.Println(story)
	}
}
