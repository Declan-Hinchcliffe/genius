package main

import (
	"encoding/json"
	"net/http"

	"github.com/joe-bricknell/genius/internal/models"

	"github.com/gorilla/mux"
	"github.com/joe-bricknell/genius/internal"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("home page coming soon!")
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
		http.Error(w, http.StatusText(400), 400)
	}
}

// GetLyricsByArtist will get the lyrics to the top 20 songs by a particular artist
func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wordsInput := "hello and the"

	songData, err := internal.GetAllLyricsByArtist(vars["artist"])
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	var wordMap map[string]int
	var data models.Response

	if songData != nil {
		wordMap, err = internal.FindWords(songData.Lyrics, &wordsInput)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
		}

		data = models.Response{
			Status:  200,
			Songs:   songData.Songs,
			Lyrics:  songData.Lyrics,
			WordMap: wordMap,
		}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
}

// GetLyricsBySearch will get all the lyrics for the 20 results of a given search
func GetLyricsBySearch(w http.ResponseWriter, r *http.Request) {
	//search := flag.String("search", "", "specify your search term")
	//flag.Parse()

	vars := mux.Vars(r)

	songData, err := internal.GetLyricsBySearch(vars["search"])
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(songData); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
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

	if err := json.NewEncoder(w).Encode(song); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

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

	data := models.Response{
		Songs:   []models.Song{*song},
		Lyrics:  []models.Lyric{*lyrics},
		WordMap: nil,
	}

	json.NewEncoder(w).Encode(data)
}
