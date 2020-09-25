package genius

import "net/url"

// GetLyricsBySearch will call to the genius api to get the songs and then call
// to the lyrics api to get the lyrics
func GetLyricsBySearch(svar string) ([]Lyrics, error) {
	encodedSearch := url.QueryEscape(svar)

	songList, err := searchSongs(encodedSearch)
	if err != nil {
		return nil, err
	}

	if songList == nil {
		return nil, err
	}

	allLyrics, err := getLyrics(songList)
	if err != nil {
		return nil, err
	}

	return allLyrics, nil

}
