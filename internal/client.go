package internal

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

type CustomClient struct {
	httpClient *http.Client
	url        string
}

// New creates a new custom client
func New(url string) (CustomClient, error) {
	return CustomClient{
		httpClient: client,
		url:        url,
	}, nil
}

func makeRequestGenius(endpoint string) (*http.Response, error) {
	c, err := New(os.Getenv("GENIUS"))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", c.url, endpoint), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(req.URL)

	req.Header.Set("Authorization", "Bearer SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func makeRequestLyrics(endpoint string) (*http.Response, error) {
	c, err := New(os.Getenv("LYRICS"))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", c.url, endpoint), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {

		return nil, err
	}

	return resp, nil
}
