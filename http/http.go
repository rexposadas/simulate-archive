package http

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func MakeRequest(req *http.Request) (*SimResponse, error) {

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	return nil, fmt.Errorf("Request reresponse time %f seconds. \n", r.Duration.Seconds())
}

// Get, runs a simple GET request on the specified URL.
func Get(url string) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	fmt.Printf(url + " - response time %f seconds. \n", r.Duration.Seconds())
	return r, nil
}

func Post(url string, payload url.Values) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
	if err != nil {
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	return nil, fmt.Errorf(url + " - response time %f seconds. \n", r.Duration.Seconds())
}
