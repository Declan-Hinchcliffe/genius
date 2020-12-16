package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-bricknell/genius/internal"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home page coming soon...")
}

// GetAllSongs will get the top 20 songs by a given artist
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := internal.GetArtistID(vars["search"])
	if err != nil {
		panic(err)
	}

	songs, err := internal.SongsByArtist(*id)
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		panic("failed to encode response")
	}
}

// GetLyricsByArtist will get the lyrics to the top 20 songs by a particular artist
func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	words := "hello and the"

	var lyrics []internal.Lyrics
	var err error
	lyrics, err = internal.GetAllLyricsByArtist(vars["artist"])
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(lyrics); err != nil {
		panic(err)
	}

	internal.FindWords(w, lyrics, &words)
}

// GetLyricsBySearch will get all the lyrics for the 20 results of a given search
func GetLyricsBySearch(w http.ResponseWriter, r *http.Request) {
	//search := flag.String("search", "", "specify your search term")
	//flag.Parse()

	vars := mux.Vars(r)

	var lyrics []internal.Lyrics
	var err error
	lyrics, err = internal.GetLyricsBySearch(vars["search"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "\n%v\n", lyrics)
}

// GetOneSongBySearch will return the closest match to a given song from a
// search result of 20
func GetOneSongBySearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songs, err := internal.SearchSongs(vars["song"])
	if err != nil {
		panic(err)
	}

	song, err := internal.GetOneSong(*songs)
	if err != nil {
		panic(err)
	}

	fmt.Println(song)

	fmt.Fprintf(w, "%v", song)

}

// GetLyricsOneSong will retrieve the lyrics for one song by searching all songs, find
// the given song and find the lyrics for this song
func GetLyricsOneSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songs, err := internal.SearchSongs(vars["song"])
	if err != nil {
		panic(err)
	}

	song, err := internal.GetOneSong(*songs)
	if err != nil {
		panic(err)
	}

	lyrics, err := internal.GetLyricsForSingleSong(*song)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(lyrics)
}
