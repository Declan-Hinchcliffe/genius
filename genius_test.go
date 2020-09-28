package genius

import (
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
			expectedResponse: LyricsResp,
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
			t.Fatalf("error when calling getLyrics %v", err)
		}
		assert.Equal(t, stripNewlineChar(tc.expectedResponse[0].Lyrics), stripNewlineChar(lyrics[0].Lyrics))
	}
}

func stripNewlineChar(lyrics string) string {
	return strings.Replace(lyrics, "\r\n", " ", -1)

}
