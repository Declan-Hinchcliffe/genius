package internal

import (
	"strings"

	"github.com/joe-bricknell/genius/internal/models"
)

// findWords will search through the lyrics and count the number of matches
// for particular words
func FindWords(songData *models.Response, flag *string) (map[string]int, error) {
	wordsFlag := strings.Fields(*flag)
	wordCounter := make([]int, len(wordsFlag))
	var numOfWords int

	for _, lyrics := range songData.Lyrics {
		for _, songWord := range strings.Fields(lyrics.Lyric) {
			numOfWords++
			for i, lookupWord := range wordsFlag {
				if lookupWord == songWord {
					wordCounter[i]++
				}
			}
		}
	}

	wordMap := make(map[string]int)

	wordMap["numOfWords"] = numOfWords

	for i, v := range wordsFlag {
		wordMap[v] = wordCounter[i]
	}

	return wordMap, nil
}
