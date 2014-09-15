package http

import (
	"net/http"
	"time"
)

// Get, runs a simple GET request on the specified URL.
func Get(url string) (*Response, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	since := time.Since(start)
	r := &Response{
		Duration: since,
	}

	return r, nil
}
