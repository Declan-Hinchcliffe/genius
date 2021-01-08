package models

// Response is the top level struct that will be returned to the user
// here we can return the list of songs or
// a lyrics we've searched for, the status code
// of the response and the word map that tells us the word count
type Response struct {
	Songs   []Song         `json:"song"`
	Lyrics  []Lyrics       `json:"lyrics"`
	WordMap map[string]int `json:"word_map"`
}

// Song represents a Song returned from the API
type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

type Lyrics struct {
	ID     int    `json:"id"`
	Lyrics string `json:"lyrics"`
}
