package genius

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/artist/songs/lyrics", GetLyricsByArtist).Methods("GET")

	// possible other routes?
	router.HandleFunc("/artist/songs", nil).Methods("GET")
	router.HandleFunc("/search/songs", nil).Methods("GET")
	router.HandleFunc("/search/songs/lyrics", nil).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func GetLyricsByArtist(w http.ResponseWriter, r *http.Request) {

}
