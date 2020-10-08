package genius

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

// Song represents a Song returned from the API
type Song struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Lyrics represents the lyrics returned from the lyric api
type Lyrics struct {
	Lyrics string `json:"lyrics"`
}

func Genius() {
	search := flag.String("search", "", "specify your search term")
	artist := flag.String("artist", "", "specify your artist")
	word := flag.String("word", "", "specify the words you want to look for")
	flag.Parse()

	var lyrics []Lyrics
	var err error
	if *search != "" {
		lyrics, err = getLyricsBySearch(search)
		if err != nil {
			panic(err)
		}
	}

	if *artist != "" {
		lyrics, err = getAllLyricsByArtist(artist)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%v\n", lyrics)

	wordMap := findWords(lyrics, word)
	displayWordCount(wordMap)
}

// getLyrics will call to the lyrics api and return the lyrics for a particular Song
func getLyrics(songList []Song) ([]Lyrics, error) {
	// create error channel to receive errors from go routines
	errCh := make(chan error)

	allLyrics := make([]Lyrics, 0, 20)
	var lyrics Lyrics

	// wait group waits for goroutines to finish
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(songList))

	for _, song := range songList {
		fmt.Printf("%v - %v\n", song.Artist, song.Title)
		go func(song Song, errCh chan<- error, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			req, err := http.NewRequest("GET", fmt.Sprintf("https://api.lyrics.ovh/v1/%v/%v", song.Artist, song.Title), strings.NewReader(""))
			if err != nil {
				errCh <- err
				return
			}

			// make request
			resp, err := client.Do(req)
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

			mu.Lock()
			allLyrics = append(allLyrics, lyrics)
			mu.Unlock()

		}(song, errCh, &wg, &mu)
	}

	// need to place this into a go routine otherwise blocks here before values are pulled off
	// without this we would hit a deadlock
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// we range the errCh to see if there are multiple errors
	// this will return the first error
	for err := range errCh {
		return nil, err
	}

	return allLyrics, nil
}
