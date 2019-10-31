package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Clear function for clearing the terminal screen
var Clear map[string]func()

// WorkingDir current working directory that util is called in
var WorkingDir string

// LocalFile reference to the local file name
var LocalFile string = ".ac_tracking.json"

// ACTrackerData main object that holds all saved data
type ACTrackerData struct {
	Stories []Story `json:"stories"`
}

func init() {
	workingDir, err := os.Getwd()
	WorkingDir = workingDir
	GeneralErr(err, "Failed to get working directiony...")

	Clear = make(map[string]func())
	Clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Clear["Windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	Clear["darwin"] = Clear["linux"]
}

func add(arg string) {
	switch strings.ToLower(arg) {
	case "story":
		AddStory()
	case "ac":
		ArgumentErr(len(os.Args) < 4, "Must specify StoryID to add AC for.\nUsage: `actracker new AC RE-360`")
		AddACToStory(os.Args[3])
	default:
		ArgumentErr(true, fmt.Sprintf("Invalid option: %s", arg))
	}
}

func show() {
	ListStories()
}

func main() {
	ArgumentErr(len(os.Args) == 1, "Need at least one command arg.")

	switch strings.ToLower(os.Args[1]) {
	case "add":
		ArgumentErr(len(os.Args) < 3, "Must specify what `new` is creating. (Hint: story, AC...)")
		add(os.Args[2])
	case "show":
		// for now we just show all items regardless of what "sprint" they are apart of
		show()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
