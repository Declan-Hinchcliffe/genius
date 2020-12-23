package internal

import (
	"fmt"
	"net/http"
	"time"
)

const (
	geniusURL = "https://api.genius.com"
	lyricsURL = "https://api.lyrics.ovh/v1"
	apiKey    = "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"
)

var netClient = http.Client{
	Timeout: time.Second * 10,
}

func makeRequestGenius(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", geniusURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func makeRequestLyrics(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", lyricsURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	resp, err := netClient.Do(req)
	if err != nil {

		return nil, err
	}

	return resp, nil
}
