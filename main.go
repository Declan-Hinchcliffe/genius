package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// want to call to the genius api and get information about artists
// make it so you can search for an artist or song
// can get the lyrics for a song
// can get all the songs by artist

type Song struct {
	Title  string
	Artist string
}

// need to define our song struct
type data struct {
	Response struct {
		Hits []Hit
	} `json:"response"`
}

type Hit struct {
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
			Concurrents           int  `json:"concurrents"`
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
}

func main() {
	var svar string
	flag.StringVar(&svar, "search", "", "specify your search term")
	flag.Parse()

	encodedSearch := url.QueryEscape(svar)

	// build request to genius api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%v", encodedSearch), strings.NewReader(""))
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
	var songs data
	if err := json.Unmarshal(body, &songs); err != nil {
		panic(err)
	}

	// range over the results and insert into new slice
	var song []string

	for _, songs := range songs.Response.Hits {
		song = append(song, songs.Result.FullTitle)
	}
	fmt.Println(song)

	//
	//splitString := strings.Split(songList[6], "by")
	//splitSong := strings.TrimSpace(splitString[0])
	//splitArtist := strings.TrimSpace(splitString[1])
	//
	////	get the lyrics for this particular song
	//req, err = http.NewRequest("GET", fmt.Sprintf("https://api.lyrics.ovh/v1/%v/%v", splitArtist, splitSong), strings.NewReader(""))
	//if err != nil {
	//	panic(err)
	//}
	//
	//resp, err = http.DefaultClient.Do(req)
	//if err != nil {
	//	panic(err)
	//}

	//lyrics, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	//var lyric []string

	// can get all the songs by artist

}
