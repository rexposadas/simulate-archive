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

	_return := fmt.Errorf("Request reresponse time %f seconds. \n", r.Duration.Seconds())
	return nil, _return
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
	_return := fmt.Errorf(url + " - response time %f seconds. \n", r.Duration.Seconds())
	return nil, _return
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

	_return := fmt.Errorf(url + " - response time %f seconds. \n", r.Duration.Seconds())
	return nil, _return
}
