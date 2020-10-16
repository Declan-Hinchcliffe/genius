package genius

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/home", homeHandler)

	// looking for songs/lyrics by artist
	router.HandleFunc("/songs/{artist}", GetAllSongs).Methods("GET")
	router.HandleFunc("/songs/lyrics/{artist}", GetLyricsByArtist).Methods("GET")

	// looking for one song or one song lyrics
	router.HandleFunc("/lyrics/{song}", GetLyricsOneSong).Methods("GET")
	router.HandleFunc("/lyrics/{search}", GetLyricsBySearch).Methods("GET")

	// looking via search
	router.HandleFunc("/search/{song}", GetOneSongBySearch).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home page coming soon...")
}

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := getArtistID(vars["artist"])
	if err != nil {
		panic(err)
	}

	songs, err := songsByArtist(*id)
	if err != nil {
		panic(err)
	}

	for i, song := range songs {
		fmt.Fprintf(w, "%v. %v - %v\n", i+1, song.Artist, song.Title)
	}
}

func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	words := "hello and the"

	var lyrics []Lyrics
	var err error
	lyrics, err = getAllLyricsByArtist(vars["artist"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "\n%v\n", lyrics)

	findWords(w, lyrics, &words)
}

func GetLyricsBySearch(w http.ResponseWriter, r *http.Request) {
	//search := flag.String("search", "", "specify your search term")
	//flag.Parse()

	vars := mux.Vars(r)

	var lyrics []Lyrics
	var err error
	lyrics, err = getLyricsBySearch(vars["search"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "\n%v\n", lyrics)
}

func GetOneSongBySearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songs, err := searchSongs(vars["song"])
	if err != nil {
		panic(err)
	}

	song, err := getOneSong(*songs)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%v - %v", song.Artist, song.Title)

}

func GetLyricsOneSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songs, err := searchSongs(vars["song"])
	if err != nil {
		panic(err)
	}

	song, err := getOneSong(*songs)
	if err != nil {
		panic(err)
	}

	lyrics, err := getLyricsOneSong(*song)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%v", lyrics)
}
