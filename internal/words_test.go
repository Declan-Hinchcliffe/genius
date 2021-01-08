package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	expectedWords := map[string]int{
		"fuck":       12,
		"shit":       6,
		"bitch":      13,
		"pussy":      21,
		"numOfWords": 2747,
	}

	flag := "fuck shit bitch pussy"

	actual, err := scanWords(testLyrics, &flag)
	if err != nil {
		t.Fatalf("error when calling FindWords")
	}

	assert.Equal(t, expectedWords, actual)
	fmt.Println(actual)
}
