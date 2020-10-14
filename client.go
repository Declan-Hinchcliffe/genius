package genius

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

func makeRequest(c CustomClient, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", c.url, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if c.url == "https://api.genius.com" {
		req.Header.Set("Authorization", "Bearer SWIZahaJ5gY3S8ZOAwLbTlpREdKOXMakvPPM_0vD5q1AXId4J4fGTDJ-VO-h0Ojp")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {

		return nil, err
	}

	return resp, nil
}
