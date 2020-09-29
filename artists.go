package genius

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// allSongs represents the data structure of the response from the genius api
type allSongsResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Songs []struct {
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
			} `json:"stats,omitempty"`
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
			} `json:"primary_artist,omitempty"`
		} `json:"songs"`
		NextPage int `json:"next_page"`
	} `json:"response"`
}

// getAllLyricsByArtist will return the lyrics to the first 20 songs by a given artist
func getAllLyricsByArtist(flag string) ([]Lyrics, error) {
	id, err := getArtistID(flag)
	if err != nil {
		return nil, err
	}

	songs, err := songsByArtist(*id)
	if err != nil {
		return nil, err
	}

	lyrics, err := getLyrics(songs)
	if err != nil {
		return nil, err
	}

	return lyrics, nil
}

// getArtistID will call to the genius api search and pull out the artist id from the first search result
func getArtistID(artistFlag string) (*int, error) {
	// url encoded flag to use in request
	encodedSearch := url.QueryEscape(artistFlag)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%v", encodedSearch), nil)
	if err != nil {
		return nil, err
	}

	token := "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "Application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var songResponse apiSearchResponse
	if err := json.Unmarshal(body, &songResponse); err != nil {
		return nil, err
	}

	if len(songResponse.Response.Hits) == 0 {
		return nil, errors.New("couldn't find id for given artist")
	}
	id := songResponse.Response.Hits[0].Result.PrimaryArtist.ID

	return &id, nil
}

// songsByArtist will retrieve all the songs by an artist using the artist id
func songsByArtist(id int) ([]Song, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/artists/%v/songs?sort=popularity", id), nil)
	if err != nil {
		return nil, err
	}

	token := "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "Application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiSongs allSongsResponse
	if err := json.Unmarshal(body, &apiSongs); err != nil {
		return nil, err
	}

	songList, err := getSongs(apiSongs)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

// getSongs will loop over a slice of Song data and retrieve the artist and title for each Song
func getSongs(apiSongs allSongsResponse) ([]Song, error) {
	var songList []Song
	for _, songs := range apiSongs.Response.Songs {
		song := Song{
			Title:  strings.TrimSpace(songs.Title),
			Artist: strings.TrimSpace(songs.PrimaryArtist.Name),
		}

		songList = append(songList, song)
	}

	return songList, nil
}
