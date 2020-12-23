package internal

import (
	"fmt"
	"net/http"
	"time"
)

const (
	genius = "https://api.genius.com"
	lyrics = "https://api.lyrics.ovh/v1"
	apiKey = "SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp"
)

var netClient = http.Client{
	Timeout: time.Second * 10,
}

func makeRequestGenius(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", genius, endpoint), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", lyrics, endpoint), nil)
	if err != nil {
		return nil, err
	}

	resp, err := netClient.Do(req)
	if err != nil {

		return nil, err
	}

	return resp, nil
}
