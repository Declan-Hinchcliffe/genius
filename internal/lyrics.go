package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/joe-bricknell/genius/internal/models"
)

// getLyrics will call to the lyrics api and return the lyrics for a particular Song
func getLyrics(songList []models.Song) ([]models.Response, error) {
	// create error channel to receive errors from go routines
	errCh := make(chan error)
	resultCh := make(chan models.Response)

	allLyrics := make([]models.Response, 0, 20)

	var wg sync.WaitGroup

	wg.Add(len(songList))

	for _, song := range songList {
		fmt.Printf("%v - %v\n", song.Artist, song.Title)
		go doRequests(resultCh, errCh, &wg, song)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(resultCh)
	}()

	go func() {
		for lyrics := range resultCh {
			allLyrics = append(allLyrics, lyrics)
		}
	}()

	for err := range errCh {
		return nil, err
	}

	return allLyrics, nil
}

// GetLyricsOneSong will retrieve the lyrics for a given song
func GetLyricsForSingleSong(song models.Song) (*models.Response, error) {
	errCh := make(chan error)
	resultCh := make(chan models.Response)

	go doRequests(resultCh, errCh, nil, song)

	select {
	case err := <-errCh:
		return nil, err
	case lyrics := <-resultCh:
		return &lyrics, nil
	}
}

func doRequests(resultCh chan<- models.Response, errCh chan<- error, wg *sync.WaitGroup, song models.Song) {
	if wg != nil {
		defer wg.Done()
	}
	var lyrics models.Response
	endpoint := fmt.Sprintf("%v/%v", song.Artist, song.Title)

	resp, err := makeRequestLyrics(endpoint)
	if err != nil {
		errCh <- err
		return
	}

	defer resp.Body.Close()

	// read body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- err
		return
	}

	// unmarshal json into lyrics struct
	if err := json.Unmarshal(body, &lyrics); err != nil {
		errCh <- err
		return
	}

	if lyrics.Lyrics == "" {
		fmt.Printf("failed to find lyrics for: %v - %v\n", song.Artist, song.Title)
	}

	resultCh <- lyrics
}
