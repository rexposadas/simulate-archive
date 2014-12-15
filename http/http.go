package http

import (
	"github.com/rexposadas/simulate"
	"net/http"
	"net/url"
	"time"
)

// MakeRequest makes an HTTP request.
// The caller is in charge of closing the response body. ( #todo: is this proper? )
func MakeRequest(req *http.Request) (*SimResponse, error) {

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		// todo: add time to series
		simulate.Stats.Error(err.Error())
		return nil, err
	}
	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	simulate.Stats.TrackResponse(req.URL.String(), since)

	if resp.StatusCode != 200 {
		simulate.Stats.Error(resp.Status)
	}

	return r, nil
}

// Get, runs a simple GET request on the specified URL.
// The caller is in charge of closing the response body. ( #todo: is this proper? )

func Get(url string) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		simulate.Stats.Error(err.Error())
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	if resp.StatusCode != 200 {
		simulate.Stats.Error(string(resp.StatusCode))
	}

	simulate.Stats.TrackResponse(url, since)
	return r, nil
}

// The caller is in charge of closing the response body. ( #todo: is this proper? )
func Post(url string, payload url.Values) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
	if err != nil {
		simulate.Stats.Error(err.Error())

		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	simulate.Stats.TrackResponse(url, since)
	if resp.StatusCode != 200 {
		simulate.Stats.Error(string(resp.StatusCode))
	}

	return r, nil
}
