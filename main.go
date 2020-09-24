package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Song represents a song returned from the API
type Song struct {
	Title  string
	Artist string
}

// need to define our song struct
type data struct {
	Response struct {
		Hits []hit
	} `json:"response"`
}

type hit struct {
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

// Lyrics represents the lyrics returned from the lyric api
type Lyrics struct {
	Lyrics string `json:"lyrics"`
}

func main() {
	var svar string
	flag.StringVar(&svar, "search", "", "specify your search term")
	flag.Parse()

	lyrics, err := getTheLyrics(svar)
	if err != nil {
		panic(err)
	}

	fmt.Println(lyrics)

	wordMap, err := findWords(*lyrics)
	if err != nil {
		panic(err)
	}

	// we range over the map to get the keys and store them in a slice
	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Printf("%v: %v,\n%v: %v,\n %v: %v\n", keys[0], wordMap[keys[0]], keys[1], wordMap[keys[1]], keys[2], wordMap[keys[2]])
}

// getTheLyrics will call to the genius api to get the songs and then call
// to the lyrics api to get the lyrics
func getTheLyrics(svar string) (*Lyrics, error) {
	encodedSearch := url.QueryEscape(svar)

	songList, err := getSongs(encodedSearch)
	if err != nil {
		return nil, err
	}

	if songList == nil {
		return nil, err
	}

	artist := songList[0].Artist
	title := songList[0].Title

	lyrics, err := getLyrics(artist, title)
	if err != nil {
		return nil, err
	}

	return lyrics, nil

}

// getSongs will call to the genius api and return a list of songs matching
// a particular search
func getSongs(search string) ([]Song, error) {
	// build request to genius api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%v", search), strings.NewReader(""))
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

	// unmarshal json into song response struct
	var apiSongResponse data
	if err := json.Unmarshal(body, &apiSongResponse); err != nil {
		panic(err)
	}

	// define our song list variable and range over the songs and add the
	// song name and artist to the song struct
	var songList []Song
	for _, songs := range apiSongResponse.Response.Hits {
		song := Song{
			Title:  strings.TrimSpace(songs.Result.Title),
			Artist: strings.TrimSpace(songs.Result.PrimaryArtist.Name),
		}

		songList = append(songList, song)
	}

	return songList, nil
}

// getLyrics will call to the lyrics api and return the lyrics for a particular song
func getLyrics(artist, title string) (*Lyrics, error) {
	fmt.Printf("Artist: %v, Song: %v\n\n", artist, title)

	//	build request to lyrics api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.lyrics.ovh/v1/%v/%v", artist, title), strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	// make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshall json into lyrics struct
	var lyrics Lyrics
	if err := json.Unmarshal(body, &lyrics); err != nil {
		return nil, err
	}

	return &lyrics, nil
}

// findWords will search through the lyrics and count the number of matches
// for particular words
func findWords(lyrics Lyrics) (map[string]int, error) {
	var fuckCount int
	var shitCount int
	var bitchCount int

	for _, word := range strings.Fields(lyrics.Lyrics) {
		switch {
		case strings.Contains(strings.ToLower(word), "fuck"):
			fuckCount++
		case strings.Contains(strings.ToLower(word), "shit"):
			shitCount++
		case strings.Contains(strings.ToLower(word), "bitch"):
			bitchCount++
		}
	}

	wordMap := make(map[string]int)

	wordMap["fuckCount"] = fuckCount
	wordMap["shitCount"] = shitCount
	wordMap["bitch"] = bitchCount

	return wordMap, nil

}
