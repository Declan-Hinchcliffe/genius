package genius

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type ClientGenius struct {
	client      *http.Client
	baseURL     string
	secretToken string
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
	fmt.Println(lyrics)

	wordMap := findWords(lyrics, word)
	displayWordCount(wordMap)
}

// getLyrics will call to the lyrics api and return the lyrics for a particular Song
func getLyrics(songList []Song) ([]Lyrics, error) {
	var allLyrics []Lyrics
	var lyrics Lyrics

	for _, song := range songList {
		fmt.Printf("%v - %v\n", song.Artist, song.Title)
		//	build request to lyrics api
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.lyrics.ovh/v1/%v/%v", song.Artist, song.Title), strings.NewReader(""))
		if err != nil {
			return nil, err
		}

		// make request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		// read body of the response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// unmarshal json into lyrics struct
		if err := json.Unmarshal(body, &lyrics); err != nil {
			return nil, err
		}

		allLyrics = append(allLyrics, lyrics)
	}

	return allLyrics, nil
}

// findWords will search through the lyrics and count the number of matches
// for particular words
func findWords(allLyrics []Lyrics, flag *string) map[string]int {
	//wordFlags := strings.Fields(*flag)
	//fmt.Println("wordflags:", wordFlags)

	var lyricCount int
	var fuckCount int
	var shitCount int
	var bitchCount int
	var pussyCount int

	for _, lyrics := range allLyrics {
		for _, word := range strings.Fields(lyrics.Lyrics) {
			lyricCount++
			switch {
			case
				strings.Contains(strings.ToLower(word), "fuck"),
				strings.Contains(strings.ToLower(word), "f-ck"),
				strings.Contains(strings.ToLower(word), "f*ck"):
				fuckCount++
			case strings.Contains(strings.ToLower(word), "shit"):
				shitCount++
			case
				strings.Contains(strings.ToLower(word), "bitch"),
				strings.Contains(strings.ToLower(word), "b*tch"),
				strings.Contains(strings.ToLower(word), "b-tch"):
				bitchCount++
			case
				strings.Contains(strings.ToLower(word), "pussy"),
				strings.Contains(strings.ToLower(word), "p*ssy"),
				strings.Contains(strings.ToLower(word), "p-ssy"):
				pussyCount++
			}

		}
	}

	fmt.Printf("total words counted: %v\n", lyricCount)

	wordMap := map[string]int{
		"fuckCount":  fuckCount,
		"shitCount":  shitCount,
		"bitchCount": bitchCount,
		"pussyCount": pussyCount,
	}

	return wordMap
}

func displayWordCount(wordMap map[string]int) {
	// we range over the map to get the keys and store them in a slice
	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Printf("%v:%v,\n%v:%v,\n%v:%v,\n%v:%v\n",
		keys[0], wordMap[keys[0]],
		keys[1], wordMap[keys[1]],
		keys[2], wordMap[keys[2]],
		keys[3], wordMap[keys[3]])

}
