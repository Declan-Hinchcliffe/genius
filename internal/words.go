package internal

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/joe-bricknell/genius/internal/log"

	"github.com/joe-bricknell/genius/internal/models"
)

var stripRegexWords = regexp.MustCompile(`[^a-zA-Z0-9']+`)
var stripRegexInput = regexp.MustCompile(`[^a-zA-Z0-9', ]+`)

// scanWords will search through the lyrics and count the number of matches
// for particular words
func ScanWords(songs []models.Song, flag *string) (map[string]int, error) {
	wordsList := strings.Fields(stripRegexInput.ReplaceAllString(*flag, ""))
	wordCount := make([]int, len(wordsList))
	wordMap := make(map[string]int)

	numOfWords := 0
	var words []string

	log.Logger.Infof("ScanWords: searching for the following words: %v", wordsList)

	for _, song := range songs {
		lyricsLowercase := strings.ToLower(stripRegexWords.ReplaceAllString(song.Lyrics.Lyrics, " "))
		formatString := strings.ReplaceAll(lyricsLowercase, "'", "")

		scan := bufio.NewScanner(strings.NewReader(formatString))
		scan.Split(bufio.ScanWords)

		for scan.Scan() {
			numOfWords++

			words = append(words, scan.Text())

			for i, lookUpWord := range wordsList {
				if lookUpWord == scan.Text() {
					wordCount[i]++
				}
			}
		}
		if err := scan.Err(); err != nil {
			return nil, err
		}
	}

	wordMap["numOfWords"] = numOfWords

	for i, v := range wordsList {
		wordMap[v] = wordCount[i]
	}

	return wordMap, nil
}
