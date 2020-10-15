package genius

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/", getAllLyricsHandler).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func getAllLyricsHandler(w http.ResponseWriter, r *http.Request) {

}
