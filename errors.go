package main

import (
	"fmt"
	"os"
)

// ArgumentErr if predicate is true, print given error and exit program with
// Non-zero status code
func ArgumentErr(predicate bool, err string) {
	if predicate {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NotFoundErr if predicate is true, exit program with non-zero status code
func NotFoundErr(predicate bool, errMessage string) {
	if predicate {
		fmt.Printf("NotFoundErr: %v\n", errMessage)
		os.Exit(1)
	}
}

// GeneralErr exits with non-zero status code if non-nil err passed in
func GeneralErr(err error, errMessage string) {
	if err != nil {
		fmt.Printf("Error: %s\n\n%v\n", errMessage, err)
		os.Exit(1)
	}
}
