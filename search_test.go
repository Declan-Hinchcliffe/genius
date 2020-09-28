package genius

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLyricsBySearch(t *testing.T) {
	flag := "cardi b"
	lyrics, err := getLyricsBySearch(flag)
	if err != nil {
		t.Fatalf("error when calling getLyricsBySearch. err: %v", err)
	}

	fmt.Println(lyrics)
}

func TestSearchSongs(t *testing.T) {
	testCases := []struct {
		desc             string
		search           string
		expectedResponse []Song
	}{
		{
			desc:   "1. successfully returns songs using search term",
			search: "drake",
			expectedResponse: []Song{
				{
					Title:  "God’s Plan",
					Artist: "Drake",
				},
				{
					Title:  "In My Feelings",
					Artist: "Drake",
				},
				{
					Title:  "Hotline Bling",
					Artist: "Drake",
				},
			},
		},
		{
			desc:             "2. response returns empty songs as api can't find song",
			search:           "krvbhrbvjhrbvhjrbv",
			expectedResponse: nil,
		},
	}

	for _, tc := range testCases {
		songs, err := searchSongs(tc.search)
		if err != nil {
			t.Fatalf("error when calling searchSongs. err: %v", err)
		}

		if songs != nil {
			assert.Equal(t, tc.expectedResponse[0:2], songs[0:2])
		} else {
			assert.Equal(t, tc.expectedResponse, songs)
		}

	}
}
