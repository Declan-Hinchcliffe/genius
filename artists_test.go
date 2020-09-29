package genius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllLyricsByArtist(t *testing.T) {
	lyrics, err := getAllLyricsByArtist("kanye west")
	if err != nil {
		t.Fatalf("error when calling getAllLyricsByArtist. err: %v", err)
	}

	assert.Equal(t, testLyrics[3], lyrics[0])
}
