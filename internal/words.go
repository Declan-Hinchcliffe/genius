package internal

import (
	"net/http"
	"strings"
)

// findWords will search through the lyrics and count the number of matches
// for particular words
func FindWords(w http.ResponseWriter, allLyrics []Lyrics, flag *string) (map[string]int, error) {
	wordsFlag := strings.Fields(*flag)
	wordCounter := make([]int, len(wordsFlag))
	var numOfWords int

	for _, lyrics := range allLyrics {
		for _, songWord := range strings.Fields(lyrics.Lyrics) {
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
