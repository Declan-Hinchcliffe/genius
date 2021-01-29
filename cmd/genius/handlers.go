package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joe-bricknell/genius/internal"
	"github.com/joe-bricknell/genius/internal/log"
	"github.com/joe-bricknell/genius/internal/models"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("home page coming soon!")
}

// GetAllSongs will get the top 20 songs by a given artist
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	var test foo

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		err := fmt.Errorf("error when reading request body: %w", err)
		log.Logger.Errorf("GetAllSongs failed: %v", err)
		return
	}

	if test.Name == "" {
		log.Logger.Infof("GetAllSongs: request body was empty: %v", test)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("GetAllSongs: successfully read request body: %v", test)

	id, err := internal.GetArtistID(test.Name)
	if err != nil {
		err := fmt.Errorf("error when retrieving artist id: %w", err)
		log.Logger.Errorf("GetAllSongs failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	songs, err := internal.SongsByArtist(*id)
	if err != nil {
		err := fmt.Errorf("error when searching songs by artist: %w", err)
		log.Logger.Errorf("GetAllSongs failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		err := fmt.Errorf("error when encoding response: %w", err)
		log.Logger.Errorf("GetAllSongs failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

// GetLyricsByArtist will get the lyrics to the top 20 songs by a particular artist
func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {
	var test foo

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		err := fmt.Errorf("error when reading request body: %w", err)
		log.Logger.Errorf("GetAllSongs failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if test.Name == "" {
		log.Logger.Infof("GetAllSongs: request body was empty: %v", test)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("GetLyricsByArtist: successfully read request body: %v", test)

	songData, err := internal.GetAllLyricsByArtist(test.Name)
	if err != nil {
		err := fmt.Errorf("error when getting all lyrics by artist: %w", err)
		log.Logger.Errorf("GetLyricsByArtist failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//var wordMap map[string]int
	//var data []models.Song
	//wordsInput := "hello and the"
	//
	//wordMap, err = internal.FindWords(songData, &wordsInput)
	//if err != nil {
	//	http.Error(w, http.StatusText(400), 400)
	//}

	if err := json.NewEncoder(w).Encode(songData); err != nil {
		err := fmt.Errorf("error when encoding response: %w", err)
		log.Logger.Errorf("GetLyricsByArtist failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

// GetLyricsBySearch will get all the lyrics for the 20 results of a given search
func GetLyricsBySearch(w http.ResponseWriter, r *http.Request) {
	var test foo

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		err := fmt.Errorf("error when reading request body: %w", err)
		log.Logger.Errorf("GetLyricsBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if test.Name == "" {
		log.Logger.Infof("GetAllSongs: request body was empty: %v", test)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("GetLyricsBySearch: successfully read request body: %v", test)

	songData, err := internal.GetLyricsBySearch(test.Name)
	if err != nil {
		err := fmt.Errorf("error when getting lyrics by search: %w", err)
		log.Logger.Errorf("GetLyricsBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(songData); err != nil {
		err := fmt.Errorf("error when encoding response: %w", err)
		log.Logger.Errorf("GetLyricsBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

// GetOneSongBySearch will return the closest match to a given song from a
// search result of 20
func GetOneSongBySearch(w http.ResponseWriter, r *http.Request) {
	var test foo

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		err := fmt.Errorf("error when decoding request body: %w", err)
		log.Logger.Errorf("GetOneSongBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if test.Name == "" {
		log.Logger.Infof("GetAllSongs: request body was empty: %v", test)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("GetOneSongBySearch: successfully read request body: %v", test)

	songs, err := internal.SearchSongs(test.Name)
	if err != nil {
		err := fmt.Errorf("error when searching songs: %w", err)
		log.Logger.Errorf("GetOneSongBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	song, err := internal.GetOneSong(*songs)
	if err != nil {
		err := fmt.Errorf("error when getting one song: %w", err)
		log.Logger.Errorf("GetOneSongBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(song); err != nil {
		err := fmt.Errorf("error when encoding response: %w", err)
		log.Logger.Errorf("GetOneSongBySearch failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

}

type foo struct {
	Name string `json:"name"`
}

// GetLyricsOneSong will retrieve the lyrics for one song by searching all songs, find
// the given song and find the lyrics for this song
func GetLyricsOneSong(w http.ResponseWriter, r *http.Request) {
	var test foo

	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		err := fmt.Errorf("error when reading request body: %w", err)
		log.Logger.Errorf("GetLyricsOneSong failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if test.Name == "" {
		log.Logger.Infof("GetAllSongs: request body was empty: %v", test)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Logger.Infof("GetLyricsOneSong: successfully read request body: %v", test)

	songs, err := internal.SearchSongs(test.Name)
	if err != nil {
		err := fmt.Errorf("error when searching songs: %w", err)
		log.Logger.Errorf("GetLyricsOneSong failed: %w", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	singleSong, err := internal.GetOneSong(*songs)
	if err != nil {
		err := fmt.Errorf("error when getting song: %w", err)
		log.Logger.Errorf("GetLyricsOneSong failed: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	songWithLyrics, err := internal.GetLyricsForSingleSong(*singleSong)
	if err != nil {
		err := fmt.Errorf("error when getting lyrics for song: %w", err)
		log.Logger.Errorf("GetLyricsOneSong failed: %w", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
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

	if err := json.NewEncoder(w).Encode(data); err != nil {
		err := fmt.Errorf("error when encoding response: %w", err)
		log.Logger.Errorf("GetLyricsOneSong failed: %v", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
