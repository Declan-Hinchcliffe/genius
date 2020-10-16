package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/home", homeHandler)

	// looking for songs/lyrics by artist
	router.HandleFunc("/songs/{search}", GetAllSongs).Methods("GET")
	router.HandleFunc("/songs/lyrics/{artist}", GetLyricsByArtist).Methods("GET")

	// looking for one song or one song lyrics
	router.HandleFunc("/lyrics/{song}", GetLyricsOneSong).Methods("GET")
	router.HandleFunc("/lyrics/{search}", GetLyricsBySearch).Methods("GET")

	// looking via search
	router.HandleFunc("/search/{song}", GetOneSongBySearch).Methods("GET")

	return router
}
