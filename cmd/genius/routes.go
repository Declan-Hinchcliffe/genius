package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/home", homeHandler)

	// looking for songs/lyrics by artist
	router.HandleFunc("/songs/search", GetAllSongs).Methods("GET", "POST")
	router.HandleFunc("/songs/lyrics/artist", GetLyricsByArtist).Methods("GET", "POST")

	// looking for one song or one song lyrics
	router.HandleFunc("/song/lyrics", GetLyricsOneSong).Methods("GET", "POST")
	router.HandleFunc("/search/lyrics/", GetLyricsBySearch).Methods("GET", "POST")

	// looking via search
	router.HandleFunc("/search/song", GetOneSongBySearch).Methods("GET", "POST")

	return router
}
