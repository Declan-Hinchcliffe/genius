package genius

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type allSongs struct {
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

// allSongsByArtist will return all the songs by a given artist
func getAllLyricsByArtist(flag *string) ([]lyrics, error) {
	id, err := getArtistID(*flag)
	if err != nil {
		return nil, err
	}

	songs, err := songsByArtist(*id)

	lyrics, err := getLyrics(songs)
	if err != nil {
		return nil, err
	}

	return lyrics, nil
}

// getArtistID will call to the genius api search and pull out the artist id from the first search result
func getArtistID(artistFlag string) (*int, error) {
	encodedSearch := url.QueryEscape(artistFlag)
	// build request to genius api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/search?q=%v", encodedSearch), strings.NewReader(""))
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
	var apiSongResponse data
	if err := json.Unmarshal(body, &apiSongResponse); err != nil {
		return nil, err
	}

	id := apiSongResponse.Response.Hits[0].Result.PrimaryArtist.ID

	return &id, nil
}

func songsByArtist(id int) ([]song, error) {
	// build request using the id obtained earlier
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.genius.com/artists/%v/songs?sort=popularity", id), strings.NewReader(""))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiSongResponse allSongs
	if err := json.Unmarshal(body, &apiSongResponse); err != nil {
		return nil, err
	}

	songList, err := getSongs(apiSongResponse)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

func getSongs(apiResponse allSongs) ([]song, error) {
	// define our song list variable and range over the songs and add the
	// song name and artist to the song struct
	var songList []song
	for _, songs := range apiResponse.Response.Songs {
		song := song{
			Title:  strings.TrimSpace(songs.Title),
			Artist: strings.TrimSpace(songs.PrimaryArtist.Name),
		}

		songList = append(songList, song)
	}

	return songList, nil
}
