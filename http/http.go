package http

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexposadas/simulate"
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

	fmt.Print("tracking ", req.URL.String())

	simulate.Stats.TrackResponse(req.URL.String(), since)

	return r, nil
}

// Get, runs a simple GET request on the specified URL.
func Get(url string) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		simulate.Stats.Add(err.Error())
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	fmt.Print("stating get")
	simulate.Stats.TrackResponse(url, since)
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
	simulate.Stats.TrackResponse(url, since)

	return r, nil
}
