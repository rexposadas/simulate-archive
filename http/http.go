package http

import (
	"net/http"
	"net/url"
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

func Post(url string, payload url.Values) (*Response, error) {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
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
