package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", testRouter)

	handler := cors.Default().Handler(r)

	if err := http.ListenAndServe("localhost:4000", handler); err != nil {
		err := fmt.Errorf("error when starting server: %w", err)
		log.Fatalf("listen and serve error: %v", err)
	}
}

type foo struct {
	Name string `json:"name"`
}

func testRouter(w http.ResponseWriter, r *http.Request) {
	var foo foo

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := fmt.Errorf("error when reading request body: %w", err)
		log.Printf("testRouter failed: %v", err)
		return
	}

	log.Printf("request recieved: %v", string(body))

	if err := json.Unmarshal(body, &foo); err != nil {
		err := fmt.Errorf("error when decoding request body: %w", err)
		log.Printf("testRouter failed: %v", err)
		return
	}

	if foo.Name == "" {
		log.Print("request body was empty, returning...")
		return
	}

	log.Printf("request body read: %v, sending response...", foo)

	data := "this is the response"

	if err := json.NewEncoder(w).Encode(data); err != nil {
		err := fmt.Errorf("error when unmarshalling request body: %w", err)
		log.Printf("testRouter failed: %v", err)
		return
	}

	log.Printf("response successfully sent: %v", data)
}
