package internal

import (
	"sort"

	"github.com/joe-bricknell/genius/internal/models"
)

// sortSongsAndLyrics sorts the slice of lyrics and then creates a new slice with each song
// and the correct lyric
func sortSongsAndLyrics(songs []models.Song, lyrics []models.Lyrics) []models.Song {
	sort.Slice(lyrics, func(i, j int) bool { return lyrics[i].ID < lyrics[j].ID })
	var songsWithLyrics []models.Song

	for i, song := range songs {
		for i2, lyric := range lyrics {
			if i == i2 {
				song = models.Song{
					ID:     song.ID,
					Title:  song.Title,
					Artist: song.Artist,
					Lyrics: models.Lyrics{
						Lyrics: lyric.Lyrics,
					},
				}
			}
		}
		songsWithLyrics = append(songsWithLyrics, song)
	}

	return songsWithLyrics
}
