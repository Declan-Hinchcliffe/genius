package genius

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//need to define our song struct
type apiSearchResponse struct {
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

// GetLyricsBySearch will call to the genius api to get the songs and then call
// to the lyrics api to get the lyrics
func getLyricsBySearch(flag *string) ([]lyrics, error) {
	encodedSearch := url.QueryEscape(*flag)

	songList, err := searchSongs(encodedSearch)
	if err != nil {
		return nil, err
	}

	if songList == nil {
		return nil, err
	}

	allLyrics, err := getLyrics(songList)
	if err != nil {
		return nil, err
	}

	return allLyrics, nil

}

// searchSongs will call to the genius api and return a list of songs matching
// a particular search
func searchSongs(search string) ([]song, error) {
	// build request to genius api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%v", search), strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	token := "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "Application/json")

	// make request to genius api
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read the body of the request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal json into song response struct
	var apiSearchRepsonse apiSearchResponse
	if err := json.Unmarshal(body, &apiSearchRepsonse); err != nil {
		return nil, err
	}

	// define our song list variable and range over the songs and add the
	// song name and artist to the song struct
	var songList []song
	for _, songs := range apiSearchRepsonse.Response.Hits {
		song := song{
			Title:  strings.TrimSpace(songs.Result.Title),
			Artist: strings.TrimSpace(songs.Result.PrimaryArtist.Name),
		}

		songList = append(songList, song)
	}

	return songList, nil
}
