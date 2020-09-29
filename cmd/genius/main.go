package main

import (
	"github.com/joe-bricknell/genius"
)

// This project will allow you to search for a particular artist or term
// and will return you the number of times a given word is used within
// the lyrics search result
func main() {
	genius.Genius()
}

// todo: need to add go routines for concurrent requests
// todo: possibly try and benchmark app to see before and after go routines?
// todo: need to add flag for search words
// todo: add different searches, for a specific song or something
