package genius

import (
	"fmt"
	"sort"
	"strings"
)

// findWords will search through the lyrics and count the number of matches
// for particular words
func findWords(allLyrics []Lyrics, flag *string) map[string]int {
	//wordFlags := strings.Fields(*flag)
	//fmt.Println("wordflags:", wordFlags)

	var lyricCount int
	var fuckCount int
	var shitCount int
	var bitchCount int
	var pussyCount int

	for _, lyrics := range allLyrics {
		for _, word := range strings.Fields(lyrics.Lyrics) {
			lyricCount++
			switch {
			case
				strings.Contains(strings.ToLower(word), "fuck"),
				strings.Contains(strings.ToLower(word), "f-ck"),
				strings.Contains(strings.ToLower(word), "f*ck"):
				fuckCount++
			case strings.Contains(strings.ToLower(word), "shit"):
				shitCount++
			case
				strings.Contains(strings.ToLower(word), "bitch"),
				strings.Contains(strings.ToLower(word), "b*tch"),
				strings.Contains(strings.ToLower(word), "b-tch"):
				bitchCount++
			case
				strings.Contains(strings.ToLower(word), "pussy"),
				strings.Contains(strings.ToLower(word), "p*ssy"),
				strings.Contains(strings.ToLower(word), "p-ssy"):
				pussyCount++
			}

		}
	}

	fmt.Printf("total words counted: %v\n", lyricCount)

	wordMap := map[string]int{
		"fuckCount":  fuckCount,
		"shitCount":  shitCount,
		"bitchCount": bitchCount,
		"pussyCount": pussyCount,
	}

	return wordMap
}

func displayWordCount(wordMap map[string]int) {
	// we range over the map to get the keys and store them in a slice
	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Printf("%v:%v,\n%v:%v,\n%v:%v,\n%v:%v\n",
		keys[0], wordMap[keys[0]],
		keys[1], wordMap[keys[1]],
		keys[2], wordMap[keys[2]],
		keys[3], wordMap[keys[3]])

}
