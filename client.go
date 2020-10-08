package genius

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}
