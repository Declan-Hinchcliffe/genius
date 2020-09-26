package genius

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLyrics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

	req := httptest.NewRequest("GET")
}
