package internal

import (
	"bufio"
	"strings"

	"github.com/joe-bricknell/genius/internal/log"

	"github.com/joe-bricknell/genius/internal/models"
)

// scanWords will search through the lyrics and count the number of matches
// for particular words
func ScanWords(songs []models.Song, flag *string) (map[string]int, error) {
	wordsList := strings.Fields(*flag)
	wordCount := make([]int, len(wordsList))
	wordMap := make(map[string]int)

	numOfWords := 0
	var words []string

	log.Logger.Infof("ScanWords: searching for the following words: %v", wordsList)

	for _, song := range songs {
		lyricsLowercase := strings.ToLower(song.Lyrics.Lyrics)

		scan := bufio.NewScanner(strings.NewReader(lyricsLowercase))
		scan.Split(bufio.ScanWords)

		for scan.Scan() {
			numOfWords++
			words = append(words, scan.Text())
		}
		if err := scan.Err(); err != nil {
			return nil, err
		}
	}

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
