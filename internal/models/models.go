package models

// Response is the top level struct that will be returned to the user
// here we can return the list of songs or
// a lyrics we've searched for, the status code
// of the response and the word map that tells us the word count
type Response struct {
	Songs []Song
	//WordMap map[string]int
}

type SongsList struct {
	Songs []Song `json:"song"`
}

// Song represents a Song returned from the API
type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Lyrics Lyrics `json:"lyrics"`
}

type Lyrics struct {
	ID     int    `json:"-"`
	Lyrics string `json:"lyrics"`
}
