package genius

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type CustomClient struct {
	httpClient *http.Client
	url        string
}

// New creates a new custom client
func New(url string) (CustomClient, error) {
	return CustomClient{
		httpClient: client,
		url:        url,
	}, nil
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
	fmt.Printf("\n%v\n", lyrics)

	wordMap := findWords(lyrics, word)
	displayWordCount(wordMap)
}

// getLyrics will call to the lyrics api and return the lyrics for a particular Song
func getLyrics(songList []Song) ([]Lyrics, error) {
	// create our custom client for lyrics api
	client, err := New(os.Getenv("LYRICS"))
	if err != nil {
		return nil, err
	}

	// create error channel to receive errors from go routines
	errCh := make(chan error)
	outCh := make(chan Lyrics)

	allLyrics := make([]Lyrics, 0, 20)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(songList))

	for _, song := range songList {
		endpoint := fmt.Sprintf("%v/%v", song.Artist, song.Title)
		fmt.Printf("%v - %v\n", song.Artist, song.Title)
		go doRequests(outCh, errCh, &wg, &mu, client, endpoint)
	}

	// need to place this into a go routine otherwise blocks here before values are pulled off
	// without this we would hit a deadlock
	go func() {
		wg.Wait()
		close(errCh)
		close(outCh)
	}()

	go func() {
		for lyrics := range outCh {
			allLyrics = append(allLyrics, lyrics)
		}
	}()

	// we range the errCh to see if there are multiple errors
	// this will return the first error
	for err := range errCh {
		return nil, err
	}

	return allLyrics, nil
}

func doRequests(outCh chan<- Lyrics, errCh chan<- error, wg *sync.WaitGroup, mu *sync.Mutex, client CustomClient, endpoint string) {
	defer wg.Done()
	var lyrics Lyrics

	resp, err := makeRequest(client, endpoint)
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

	outCh <- lyrics

}
