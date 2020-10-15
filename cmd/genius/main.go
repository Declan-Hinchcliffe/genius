package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/joe-bricknell/genius"
)

// init loads env vars in .env
func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error when loading .env")
	}
}

// This project will allow you to search for a particular artist or term
// and will return you the number of times a given word is used within
// the lyrics search result
func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("exectution time - %v\n", time.Since(start))
	}()

	r := genius.Router()

	fmt.Println("starting server on port 9000...")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatalf("failed to start server (err: %v)", err)
	}

	genius.Genius()
}

// todo tasks
// add different searches, for a specific song or something
// make server? different routes
// add database? could cache song lyrics and return them quickly?

// Completed tasks
// 1. add unit tests for as much as possible
// 2. add benchmarks for all the tests
// 3. need to add go routines for concurrent
// need to add flag for search words
