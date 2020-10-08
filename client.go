package genius

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

func makeRequest(song Song, c CustomClient, endpoint string) (*http.Response, error) {
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
