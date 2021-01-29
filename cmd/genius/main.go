package main

import (
	"fmt"
	goLog "log"
	"net/http"

	"github.com/joe-bricknell/genius/internal/log"
	"github.com/rs/cors"
)

func main() {
	if err := log.InitLogger(); err != nil {
		goLog.Fatalf("failed to initalise logger: %w", err)
	}

	r := NewRouter()

	handler := cors.Default().Handler(r)

	fmt.Println("starting server on port 9000...")

	if err := http.ListenAndServe("localhost:9000", handler); err != nil {
		log.Logger.Fatalf("failed to start server: %w", err.Error())
	}
}
