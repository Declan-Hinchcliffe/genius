package main

import (
	"fmt"
	"time"

	"github.com/joe-bricknell/genius"
)

// This project will allow you to search for a particular artist or term
// and will return you the number of times a given word is used within
// the lyrics search result
func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("exectution time - %v\n", time.Since(start))
	}()

	genius.Genius()
}

// todo tasks
// need to add flag for search words
// add different searches, for a specific song or something

// Completed tasks
// 1. add unit tests for as much as possible
// 2. add benchmarks for all the tests
// 3. need to add go routines for concurrent requests
