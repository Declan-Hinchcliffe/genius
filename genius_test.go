package genius

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLyrics(t *testing.T) {
	testCases := []struct {
		desc           string
		songs          []Song
		expectedLyrics []Lyrics
	}{
		{
			desc: "1. response returns lyrics",
			songs: []Song{
				{
					Title:  "God's Plan",
					Artist: "Drake",
				},
			},
			expectedLyrics: testLyrics,
		},
		{
			desc: "2. response returns empty lyrics as api can't find song",
			songs: []Song{
				{
					Title:  "m.A.A.d city",
					Artist: "Kendrick Lamar",
				},
			},
			expectedLyrics: []Lyrics{{Lyrics: ""}},
		},
	}

	for _, tc := range testCases {
		lyrics, err := getLyrics(tc.songs)
		if err != nil {
			t.Fatalf("error when calling getLyrics. err: %v", err)
		}
		assert.Equal(t, stripNewlineChar(tc.expectedLyrics[0].Lyrics), stripNewlineChar(lyrics[0].Lyrics))
	}
}

func stripNewlineChar(lyrics string) string {
	return strings.Replace(lyrics, "\r\n", " ", -1)
}

func TestFindWords(t *testing.T) {
	testCases := []struct {
		desc          string
		lyrics        []Lyrics
		expectedWords map[string]int
	}{
		{
			desc:   "1. returns word map with no error",
			lyrics: testLyrics,
			expectedWords: map[string]int{
				"fuckCount":  12,
				"shitCount":  6,
				"bitchCount": 13,
				"pussyCount": 21,
			},
		},
	}

	for _, tc := range testCases {
		wordMap := findWords(tc.lyrics, nil)
		assert.Equal(t, wordMap, tc.expectedWords)

		fmt.Println(wordMap)
	}
}

func BenchmarkGetLyrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		actual, _ := getLyrics(testSongs)
		_ = actual
	}

	// benchmark before concurrency
	// BenchmarkGetLyrics-8   	       1	2061720173 ns/op (2.06s)
}

func BenchmarkGenius(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Genius()
	}

	// search original benchmark
	// BenchmarkGenius-8   	       1	3145686350 ns/op (3.14s)

	// artist original benchmark
	// BenchmarkGenius-8   	       1	7300000000 ns/op (7.3s ish)
	// BenchmarkGenius-8   	       1	5330076596 ns/op (5.3s)
}
