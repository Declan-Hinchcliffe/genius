package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"

	"github.com/joe-bricknell/genius/internal/models"
)

// allSongs represents the data structure of the response from the genius api
type geniusAllSongsResponse struct {
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

// GetLyricsOneSong will retrieve the lyrics for a given song
func GetLyricsForSingleSong(song models.Song) (models.Lyrics, error) {
	errCh := make(chan error)
	resultCh := make(chan models.Lyrics)

	go doRequests(resultCh, errCh, nil, song)

	select {
	case err := <-errCh:
		return models.Lyrics{}, err
	case lyrics := <-resultCh:
		//data := models.Response{
		//	Songs: []models.Song{
		//		{
		//			ID:     song.ID,
		//			Title:  song.Title,
		//			Artist: song.Artist,
		//			Lyrics: models.Lyrics{
		//				ID:     lyrics.ID,
		//				Lyrics: lyrics.Lyrics,
		//			},
		//		},
		//	},
		//}
		return lyrics, nil
	}
}

// getAllLyricsByArtist will return the lyrics to the first 20 songs by a given artist
func GetAllLyricsByArtist(artist string) (models.Response, error) {
	id, err := GetArtistID(artist)
	if err != nil {
		return models.Response{}, err
	}

	songs, err := SongsByArtist(*id)
	if err != nil {
		return models.Response{}, err
	}

	lyrics, err := getLyrics(songs)
	if err != nil {
		return models.Response{}, err
	}

	sort.Slice(lyrics, func(i, j int) bool { return lyrics[i].ID < lyrics[j].ID })

	var songsWithLyrics []models.Song

	for i, song := range songs {

		for i2, lyric := range lyrics {

			if i == i2 {
				song = models.Song{
					ID:     song.ID,
					Title:  song.Title,
					Artist: song.Artist,
					Lyrics: models.Lyrics{
						Lyrics: lyric.Lyrics,
					},
				}
			}

		}

		songsWithLyrics = append(songsWithLyrics, song)

	}

	data := models.Response{
		Songs: songsWithLyrics,
	}

	return data, nil
}

// getArtistID will call to the genius api search and pull out the artist id from the first search result
func GetArtistID(artist string) (*int, error) {
	endpoint := fmt.Sprintf("search?q=%v", url.QueryEscape(artist))

	resp, err := makeRequestGenius(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var songResponse geniusApiResponse
	if err := json.Unmarshal(body, &songResponse); err != nil {
		return nil, err
	}

	if len(songResponse.Response.Hits) == 0 {
		return nil, errors.New("couldn't find id for given artist")
	}

	id := songResponse.Response.Hits[0].Result.PrimaryArtist.ID

	fmt.Printf("successfully retrieved id for %v - %v\n", artist, id)

	return &id, nil
}

// songsByArtist will retrieve all the songs by an artist using the artist id
func SongsByArtist(id int) ([]models.Song, error) {
	endpoint := fmt.Sprintf("artists/%v/songs?sort=popularity", id)

	resp, err := makeRequestGenius(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var allSongs geniusAllSongsResponse
	if err := json.Unmarshal(body, &allSongs); err != nil {
		return nil, err
	}

	songList, err := getSongsForArtist(allSongs)
	if err != nil {
		return nil, err
	}

	return songList, nil
}

// getSongs will loop over a slice of Song data and retrieve the artist and title for each Song
func getSongsForArtist(allSongs geniusAllSongsResponse) ([]models.Song, error) {
	var songList []models.Song
	for i, songs := range allSongs.Response.Songs {
		song := models.Song{
			ID:     i,
			Title:  strings.TrimSpace(songs.Title),
			Artist: strings.TrimSpace(songs.PrimaryArtist.Name),
		}

		songList = append(songList, song)
	}

	return songList, nil
}
