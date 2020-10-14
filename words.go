package genius

import (
	"fmt"
	"strings"
)

// findWords will search through the lyrics and count the number of matches
// for particular words
func findWords(allLyrics []Lyrics, flag *string) {
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

	fmt.Printf("total words counted: %v\n", numOfWords)

	wordMap := make(map[string]int)

	for i, v := range wordsFlag {
		wordMap[v] = wordCounter[i]
	}

	for i, v := range wordsFlag {
		fmt.Printf("'%v' total: %v\n", v, wordCounter[i])
	}
}
