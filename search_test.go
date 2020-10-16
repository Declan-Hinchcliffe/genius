package genius

//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetLyricsBySearch(t *testing.T) {
//	flag := "eminem"
//	loadEnv()
//	lyrics, err := getLyricsBySearch(&flag)
//	if err != nil {
//		t.Fatalf("error calling getLyricsBySearch. err: %v", err)
//	}
//
//	if lyrics[0].Lyrics != "" {
//		fmt.Printf("successfully found lyrics: %v", lyrics[0].Lyrics)
//		assert.NotNil(t, lyrics[0].Lyrics)
//	} else {
//		t.Fatalf("\n found no lyrics, try again in a minute. lyrics: %v\n", lyrics)
//	}
//}
//
//func TestSearchSongs(t *testing.T) {
//	testCases := []struct {
//		desc   string
//		search string
//	}{
//		{
//			desc:   "1. successfully returns songs using search term",
//			search: "drake",
//		},
//		{
//			desc:   "2. response returns empty songs as api can't find song",
//			search: "krvbhrbvjhrbvhjrbv",
//		},
//	}
//
//	for _, tc := range testCases {
//		songs, err := searchSongs(tc.search)
//		if err != nil {
//			t.Fatalf("error when calling searchSongs. err: %v", err)
//		}
//
//		if len(songs) > 0 {
//			fmt.Printf("successfully found songs: %v - %v\n", songs[0].Artist, songs[0].Title)
//			assert.NotNil(t, songs[0].Artist, songs[0].Title)
//		} else {
//			assert.Equal(t, songs, []Song{})
//		}
//
//	}
//}
//
//func BenchmarkGetLyricsBySearch(b *testing.B) {
//	flag := "katy perry"
//	for i := 0; i < b.N; i++ {
//		actual, _ := getLyricsBySearch(&flag)
//		_ = actual
//	}
//
//	// original before concurrency
//	// BenchmarkGetLyricsBySearch-8   	       1	2443260642 ns/op
//	// BenchmarkGetLyricsBySearch-8   	       1	2378930002 ns/op
//
//}
//
//func BenchmarkSearchSongs(b *testing.B) {
//	search := "ariana+grande"
//	for i := 0; i < b.N; i++ {
//		actual, _ := searchSongs(search)
//		_ = actual
//	}
//
//	// original before concurrency
//	// BenchmarkSearchSongs-8   	       1	1106460523 ns/op
//	// BenchmarkSearchSongs-8   	       1	1122128581 ns/op
//
//}
