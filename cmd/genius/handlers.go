package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joe-bricknell/genius/internal"
	"github.com/joe-bricknell/genius/internal/models"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode("home page coming soon!")
}

type Artist struct {
	Name string `json:"name"`
}

// GetAllSongs will get the top 20 songs by a given artist
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	var test Artist

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Printf("failed to decode response body: %v", err)
	}

	id, err := internal.GetArtistID(test.Name)
	if err != nil {
		panic(err)
	}

	songs, err := internal.SongsByArtist(*id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
}

// GetLyricsByArtist will get the lyrics to the top 20 songs by a particular artist
func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {
	var test Artist

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Printf("failed to decode response body: %v", err)
	}

	songData, err := internal.GetAllLyricsByArtist(test.Name)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	//var wordMap map[string]int
	//var data []models.Song
	//wordsInput := "hello and the"
	//
	//wordMap, err = internal.FindWords(songData, &wordsInput)
	//if err != nil {
	//	http.Error(w, http.StatusText(400), 400)
	//}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(songData); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
}

// GetLyricsBySearch will get all the lyrics for the 20 results of a given search
func GetLyricsBySearch(w http.ResponseWriter, r *http.Request) {
	var test Artist

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Printf("failed to decode response body: %v", err)
	}

	songData, err := internal.GetLyricsBySearch(test.Name)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(songData); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
}

// GetOneSongBySearch will return the closest match to a given song from a
// search result of 20
func GetOneSongBySearch(w http.ResponseWriter, r *http.Request) {
	var test Artist

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Printf("failed to decode response body: %v", err)
	}

	songs, err := internal.SearchSongs(test.Name)
	if err != nil {
		panic(err)
	}

	song, err := internal.GetOneSong(*songs)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(song); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

}

// GetLyricsOneSong will retrieve the lyrics for one song by searching all songs, find
// the given song and find the lyrics for this song
func GetLyricsOneSong(w http.ResponseWriter, r *http.Request) {
	var test Artist

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		log.Printf("failed to decode response body: %v", err)
	}

	songs, err := internal.SearchSongs(test.Name)
	if err != nil {
		panic(err)
	}

	singleSong, err := internal.GetOneSong(*songs)
	if err != nil {
		panic(err)
	}

	songWithLyrics, err := internal.GetLyricsForSingleSong(*singleSong)
	if err != nil {
		panic(err)
	}

	//wordsInput := "hello and the"
	//
	//wordMap, err := internal.FindWords(data, &wordsInput)
	//if err != nil {
	//	http.Error(w, http.StatusText(400), 400)
	//}

	//_ = wordMap

	data := models.Song{
		ID:     singleSong.ID,
		Title:  singleSong.Title,
		Artist: singleSong.Artist,
		Lyrics: models.Lyrics{
			ID:     songWithLyrics.ID,
			Lyrics: songWithLyrics.Lyrics,
		},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
}
