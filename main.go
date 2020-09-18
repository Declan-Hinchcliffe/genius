package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// want to call to the genius api and get information about artists
// make it so you can search for an artist or song
// can get the lyrics for a song
// can get all the songs by artist

// need to define our song struct
type song struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Hits []struct {
			Highlights []interface{} `json:"highlights"`
			Index      string        `json:"index"`
			Type       string        `json:"type"`
			Result     struct {
				AnnotationCount          int    `json:"annotation_count"`
				APIPath                  string `json:"api_path"`
				FullTitle                string `json:"full_title"`
				HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url"`
				HeaderImageURL           string `json:"header_image_url"`
				ID                       int    `json:"id"`
				LyricsOwnerID            int    `json:"lyrics_owner_id"`
				LyricsState              string `json:"lyrics_state"`
				Path                     string `json:"path"`
				PyongsCount              int    `json:"pyongs_count"`
				SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url"`
				SongArtImageURL          string `json:"song_art_image_url"`
				Stats                    struct {
					UnreviewedAnnotations int  `json:"unreviewed_annotations"`
					Hot                   bool `json:"hot"`
					Pageviews             int  `json:"pageviews"`
				} `json:"stats"`
				Title             string `json:"title"`
				TitleWithFeatured string `json:"title_with_featured"`
				URL               string `json:"url"`
				PrimaryArtist     struct {
					APIPath        string `json:"api_path"`
					HeaderImageURL string `json:"header_image_url"`
					ID             int    `json:"id"`
					ImageURL       string `json:"image_url"`
					IsMemeVerified bool   `json:"is_meme_verified"`
					IsVerified     bool   `json:"is_verified"`
					Name           string `json:"name"`
					URL            string `json:"url"`
					Iq             int    `json:"iq"`
				} `json:"primary_artist"`
			} `json:"result"`
		} `json:"hits"`
	} `json:"response"`
}

func main() {
	baseUrl := `https://api.genius.com/`
	reqBody := []byte("")

	// build request to genius api
	req, err := http.NewRequest("GET", baseUrl+"search?q=24kGoldn", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	token := "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "Application/json")

	// make request to genius api
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// read the body of the request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal json into struct
	var songs song
	if err := json.Unmarshal(body, &songs); err != nil {
		panic(err)
	}

	fmt.Println(songs)

	// filter down data so you can get information by artist or song

	// get the lyrics for this particular song

	// can get all the songs by artist

}
