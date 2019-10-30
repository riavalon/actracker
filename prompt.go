package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Saver interface for saving a struct to local file
type Saver interface {
	Save(io.Writer, ACTrackerData) error
}

// GetString prompt user wiht `promptMessage` and get input
func GetString(promptMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(aurora.Bold(aurora.Cyan(promptMessage)))
	response, err := reader.ReadString('\n')
	if err != nil {
		GeneralErr(err, "Failed to parse input (ln 25 prompt.go)")
	}
	return strings.TrimRight(response, "\r\n")
}

// GetFileHandler creates or opens (append mode) a file in the current working directory
// in order to write to it. Returns an io.Writer
func GetFileHandler(filename string, fileMode int) *os.File {
	filePath := fmt.Sprintf("%s/%s", WorkingDir, filename)
	if _, err := os.Stat(filePath); err == nil {
		// file exists, use that file. Perm: Owner -> r, w, e; Other -> r, w
		f, err := os.OpenFile(filePath, fileMode, 0644)
		GeneralErr(err, fmt.Sprintf("Failed to read file: %s\nError: %v", filePath, err))
		return f
	}
	f, err := os.Create(filePath)
	GeneralErr(err, fmt.Sprintf("Failed to create file: %s\nError: %v", filePath, err))
	return f
}

func getLocalData(data *[]byte) (ACTrackerData, error) {
	var acTrackerData ACTrackerData
	if len(*data) == 0 {
		return acTrackerData, nil
	}
	err := json.Unmarshal(*data, &acTrackerData)
	return acTrackerData, err
}

// WriteObjectToMemory takes a struct and saves it to the a file in the local directory
func WriteObjectToMemory(obj Saver) error {
	allData, readErr := ReadObjectsFromMemory()
	GeneralErr(readErr, "Failed to read from local file")
	localFile := GetFileHandler(LocalFile, os.O_WRONLY)
	defer localFile.Close()

	parsedData, err := getLocalData(&allData)
	GeneralErr(err, "Failed to parse data stored in local file")

	return obj.Save(localFile, parsedData)
}

// ReadObjectsFromMemory loads local file and unmarshals data into object
func ReadObjectsFromMemory() ([]byte, error) {
	localFile := GetFileHandler(LocalFile, os.O_RDONLY)
	defer localFile.Close()
	data, err := ioutil.ReadAll(localFile)
	if err != nil {
		panic(err)
	}
	return data, localFile.Sync()
}
