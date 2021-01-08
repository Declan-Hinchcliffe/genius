package internal

//
//import (
//	"errors"
//	"fmt"
//	"log"
//	"testing"
//
//	"github.com/joho/godotenv"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetAllLyricsByArtist(t *testing.T) {
//	validArtist := "kanye west"
//	invalidArtist := "hewbcjhwbcjhwe"
//	testCases := []struct {
//		desc        string
//		flag        *string
//		expectedErr error
//	}{
//		{
//			desc:        "1. can successfully retrieve all lyrics by given artist",
//			flag:        &validArtist,
//			expectedErr: nil,
//		},
//		{
//			desc:        "2. returns error as can't find artist id for given artist",
//			flag:        &invalidArtist,
//			expectedErr: errors.New("couldn't find id for given artist"),
//		},
//	}
//
//	for _, tc := range testCases {
//		loadEnv()
//		lyrics, err := GetAllLyricsByArtist(tc.flag)
//		if tc.expectedErr != nil {
//			assert.Equal(t, tc.expectedErr, err)
//		}
//		if len(lyrics) > 0 {
//			assert.NotEmpty(t, lyrics[0].Lyrics)
//			fmt.Printf("successfully found lyrics: %v\n", lyrics[0].Lyrics)
//		}
//	}
//}
//
//func TestGetArtistID(t *testing.T) {
//	id := 33385
//	testCases := []struct {
//		desc        string
//		flag        string
//		expectedID  *int
//		expectedErr error
//	}{
//		{
//			desc:        "1. can successfully retrieve the ID for an artist",
//			flag:        "abba",
//			expectedID:  &id,
//			expectedErr: nil,
//		},
//		{
//			desc:        "2. response returns error as artist ID can't be retrieved",
//			flag:        "ekbwhbwejhcbwj",
//			expectedID:  nil,
//			expectedErr: errors.New("couldn't find id for given artist"),
//		},
//	}
//
//	for _, tc := range testCases {
//		c := CustomClient{
//			httpClient: client,
//			url:        "https://api.genius.com",
//		}
//
//		id, err := GetArtistID(tc.flag, c)
//		if tc.expectedErr != nil {
//			assert.Equal(t, tc.expectedErr, err)
//		}
//		assert.Equal(t, tc.expectedID, id)
//
//	}
//}
//
//func BenchmarkGetAllLyricsByArtist(b *testing.B) {
//	flag := "fetty wap"
//	for i := 0; i < b.N; i++ {
//		actual, _ := GetAllLyricsByArtist(&flag)
//		_ = actual
//	}
//
//	// original before concurrency
//	// BenchmarkGetAllLyricsByArtist-8   	       1	5960556438 ns/op
//	// BenchmarkGetAllLyricsByArtist-8   	       1	4302533566 ns/op
//}
//
//func BenchmarkGetArtistID(b *testing.B) {
//	flag := "chris brown"
//	c := CustomClient{
//		url:        "https://api.genius.com",
//		httpClient: client,
//	}
//
//	for i := 0; i < b.N; i++ {
//		actual, _ := GetArtistID(flag, c)
//		_ = actual
//	}
//
//	// original before concurrency
//	// BenchmarkGetArtistID-8   	       1	1237129692 ns/op
//	// BenchmarkGetArtistID-8   	       1	1481188656 ns/op
//
//}
//
//func BenchmarkSongsByArtist(b *testing.B) {
//	flag := 820
//	c := CustomClient{
//		url:        "https://api.genius.com",
//		httpClient: client,
//	}
//	actual, _ := SongsByArtist(flag, c)
//	_ = actual
//
//	// original before concurrency
//	// BenchmarkSongsByArtist-8   	       1	1704035128 ns/op
//	// BenchmarkSongsByArtist-8   	       1	1826186030 ns/op
//
//}
//
//func loadEnv() {
//	var err error
//	err = godotenv.Load(".env")
//	if err != nil {
//		log.Fatalln("error when loading .env")
//	}
//}
