package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strings"
)

// AC model for describing acceptance criteria
type AC struct {
	Task string `json:"task"`
}

// ACList struct keeping track of the items and the story
type ACList struct {
	Items   []AC   `json:"items"`
	StoryID string `json:"storyID"`
}

// Save serializes the AC for a given story into the local file
func (a *ACList) Save(w io.Writer, data ACTrackerData) error {
	var newStory Story
	for idx, story := range data.Stories {
		if story.ID == a.StoryID {
			newStory = story
			newStory.AC = *a
			if len(data.Stories) == 1 {
				data.Stories = []Story{newStory}
				break
			}
			data.Stories = append(data.Stories[:idx], data.Stories[idx+1:]...)
			data.Stories = append(data.Stories, newStory)
			break
		}
	}
	return json.NewEncoder(w).Encode(data)
}

// AddACToStory setups a loop for creating AC items for a given story with arg `storyID`
func AddACToStory(storyID string) {
	data, err := GetStoredACTrackingData()
	GeneralErr(err, fmt.Sprintf("Failed to read existing ACTrackerData\n\n%v", err))
	var allAC ACList
	for _, story := range data.Stories {
		fmt.Println("Comparing", story.ID, "with cli arg", storyID)
		if strings.ToLower(story.ID) == strings.ToLower(storyID) {
			allAC = getACForStoryInLoop(&story)
			Clear[runtime.GOOS]()
			fmt.Printf("Added %v criteria to story [%v]\n", len(allAC.Items), allAC.StoryID)
			break
		}
	}
	if allAC.StoryID == "" {
		NotFoundErr(true, fmt.Sprintf("Failed to find story with ID: [%v]", storyID))
	}
	WriteObjectToMemory(&allAC)
}

func getACForStoryInLoop(story *Story) ACList {
	acList := ACList{
		StoryID: story.ID,
		Items:   []AC{},
	}
	for {
		Clear[runtime.GOOS]()
		fmt.Print("Add details for the AC (type \\quit to stop adding AC):\n\n")
		task := GetString("Task: ")
		if strippedTask := stripSpaceChars(task); strippedTask == "\\quit" {
			break
		}
		acList.Items = append(acList.Items, AC{Task: task})
		fmt.Println("AC_LIST", acList)
	}
	return acList
}

func stripSpaceChars(task string) string {
	stripped := strings.Trim(task, "\n")
	stripped = strings.ReplaceAll(stripped, " ", "")
	stripped = strings.ToLower(stripped)
	return stripped
}
