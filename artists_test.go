package genius

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllLyricsByArtist(t *testing.T) {
	testCases := []struct {
		desc           string
		flag           string
		expectedLyrics []Lyrics
		expectedErr    error
	}{
		{
			desc:           "1. can successfully retrieve all lyrics by given artist",
			flag:           "Kanye West",
			expectedLyrics: []Lyrics{testLyrics[3]},
			expectedErr:    nil,
		},
		{
			desc:           "2. returns error as can't find artist id for given artist",
			flag:           "hewbcjhwbcjhwe",
			expectedLyrics: nil,
			expectedErr:    errors.New("couldn't find id for given artist"),
		},
	}

	for _, tc := range testCases {
		lyrics, err := getAllLyricsByArtist(tc.flag)
		if tc.expectedErr != nil {
			assert.Equal(t, tc.expectedErr, err)
		}
		if lyrics != nil {
			assert.Equal(t, tc.expectedLyrics, []Lyrics{lyrics[0]})
		}
	}
}

func TestGetArtistID(t *testing.T) {
	id := 33385
	testCases := []struct {
		desc        string
		flag        string
		expectedID  *int
		expectedErr error
	}{
		{
			desc:        "1. can successfully retrieve the ID for an artist",
			flag:        "abba",
			expectedID:  &id,
			expectedErr: nil,
		},
		{
			desc:        "2. response returns error as artist ID can't be retrieved",
			flag:        "ekbwhbwejhcbwj",
			expectedID:  nil,
			expectedErr: errors.New("couldn't find id for given artist"),
		},
	}

	for _, tc := range testCases {
		id, err := getArtistID(tc.flag)
		if tc.expectedErr != nil {
			assert.Equal(t, tc.expectedErr, err)
		}
		assert.Equal(t, tc.expectedID, id)

	}
}
