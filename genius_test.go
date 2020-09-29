package genius

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLyrics(t *testing.T) {
	testCases := []struct {
		desc             string
		songs            []Song
		expectedResponse []Lyrics
	}{
		{
			desc: "1. response returns lyrics",
			songs: []Song{
				{
					Title:  "God's Plan",
					Artist: "Drake",
				},
			},
			expectedResponse: testLyrics,
		},
		{
			desc: "2. response returns empty lyrics as api can't find song",
			songs: []Song{
				{
					Title:  "m.A.A.d city",
					Artist: "Kendrick Lamar",
				},
			},
			expectedResponse: []Lyrics{{Lyrics: ""}},
		},
	}

	for _, tc := range testCases {
		lyrics, err := getLyrics(tc.songs)
		if err != nil {
			t.Fatalf("error when calling getLyrics. err: %v", err)
		}
		assert.Equal(t, stripNewlineChar(tc.expectedResponse[0].Lyrics), stripNewlineChar(lyrics[0].Lyrics))
	}
}

func stripNewlineChar(lyrics string) string {
	return strings.Replace(lyrics, "\r\n", " ", -1)
}

func TestFindWords(t *testing.T) {
	testCases := []struct {
		desc             string
		lyrics           []Lyrics
		expectedResponse map[string]int
	}{
		{
			desc:   "1. returns word map with no error",
			lyrics: testLyrics,
			expectedResponse: map[string]int{
				"fuckCount":  4,
				"shitCount":  3,
				"bitchCount": 1,
				"pussyCount": 18,
			},
		},
	}

	for _, tc := range testCases {
		wordMap := findWords(tc.lyrics, "")
		assert.Equal(t, wordMap, tc.expectedResponse)

		fmt.Println(wordMap)
	}
}
