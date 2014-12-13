package http

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func MakeRequest(req *http.Request) error {

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	fmt.Printf("Request reresponse time %f seconds. \n", r.Duration.Seconds())
	return nil
}

// Get, runs a simple GET request on the specified URL.
func Get(url string) error {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	fmt.Printf(url + " - response time %f seconds. \n", r.Duration.Seconds())
	return nil
}

func Post(url string, payload url.Values) error {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
	if err != nil {
		return err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	fmt.Printf(url + " - response time %f seconds. \n", r.Duration.Seconds())
	return nil
}
