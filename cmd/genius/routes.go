package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home", homeHandler)

	// get top 20 songs with lyrics by search
	r.HandleFunc("/songs/lyrics/artist", GetLyricsByArtist)
	// get one song and its lyrics
	r.HandleFunc("/song/lyrics", GetLyricsOneSong)
	// get songs by searching
	r.HandleFunc("/search/lyrics/", GetLyricsBySearch)

	// i think im going to remove these
	r.HandleFunc("/songs/search", GetAllSongs)
	r.HandleFunc("/search/song", GetOneSongBySearch)

	return r
}
