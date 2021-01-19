package internal

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/joe-bricknell/genius/internal/models"
)

// findWords will search through the lyrics and count the number of matches
// for particular words
func FindWords(songData models.Response, flag *string) (map[string]int, error) {
	wordsFlag := strings.Fields(*flag)
	wordCounter := make([]int, len(wordsFlag))
	var numOfWords int

	for _, songs := range songData.Songs {
		for _, songWord := range strings.Fields(songs.Lyrics.Lyrics) {
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

func scanWords(lyrics []models.Lyrics, flag *string) (map[string]int, error) {
	wordsList := strings.Fields(*flag)
	wordCount := make([]int, len(wordsList))
	wordMap := make(map[string]int)

	numOfWords := 0
	var words []string

	for _, v := range lyrics {
		scan := bufio.NewScanner(strings.NewReader(v.Lyrics))
		scan.Split(bufio.ScanWords)

		for scan.Scan() {
			numOfWords++
			words = append(words, scan.Text())
		}
		if err := scan.Err(); err != nil {
			return nil, err
		}
	}

	fmt.Println(words)

	for _, songWord := range words {
		for i, lookUpWord := range wordsList {
			if lookUpWord == songWord {
				wordCount[i]++
			}
		}
	}

	wordMap["numOfWords"] = numOfWords

	for i, v := range wordsList {
		wordMap[v] = wordCount[i]
	}

	return wordMap, nil
}
