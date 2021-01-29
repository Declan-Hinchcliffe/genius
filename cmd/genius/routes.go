package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home", homeHandler)

	// looking for songs/lyrics by artist
	r.HandleFunc("/songs/search", GetAllSongs)
	r.HandleFunc("/songs/lyrics/artist", GetLyricsByArtist)

	// looking for one song or one song lyrics
	r.HandleFunc("/song/lyrics", GetLyricsOneSong)
	r.HandleFunc("/search/lyrics/", GetLyricsBySearch)

	// looking via search
	r.HandleFunc("/search/song", GetOneSongBySearch)

	return r
}
