package internal

import (
	"testing"

	"github.com/joe-bricknell/genius/internal/models"

	"github.com/stretchr/testify/assert"
)

// can't fix this test says no lyrics are returned
func TestGetLyrics(t *testing.T) {
	testCases := []struct {
		desc           string
		songs          []models.Song
		expectedLyrics []models.Lyric
	}{
		{
			desc: "1. response returns lyrics",
			songs: []models.Song{
				{
					Title:  "God's Plan",
					Artist: "Drake",
				},
			},
		},
		{
			desc: "2. response returns empty lyrics as api can't find song",
			songs: []models.Song{
				{
					Title:  "m.A.A.d city",
					Artist: "Kendrick Lamar",
				},
			},
			expectedLyrics: []models.Lyric{{Lyric: ""}},
		},
	}

	for _, tc := range testCases {
		lyrics, err := getLyrics(tc.songs)
		if err != nil {
			t.Fatalf("error when calling getLyrics. err: %v", err)
		}

		if tc.expectedLyrics != nil {
			assert.Equal(t, tc.expectedLyrics, lyrics)
		} else {
			assert.NotEmpty(t, lyrics[0].Lyric)
		}

	}
}

//func BenchmarkGetLyrics(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		got, _ := getLyrics(testSongs)
//		_ = got
//	}
//
//	// benchmark before concurrency
//	// BenchmarkGetLyrics-8   	       1	2061720173 ns/op (2.06s)
//}
//
//func BenchmarkGenius(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Genius()
//	}
//
//	// search original benchmark
//	// BenchmarkGenius-8   	       1	3145686350 ns/op (3.14s)
//
//	// artist original benchmark
//	// BenchmarkGenius-8   	       1	7300000000 ns/op (7.3s ish)
//	// BenchmarkGenius-8   	       1	5330076596 ns/op (5.3s)
//}
//
//func BenchmarkFindWords(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		got := findWords(testLyrics, nil)
//		_ = got
//	}
//
//	// BenchmarkFindWords-8   	    1358	    881080 ns/op
//}
