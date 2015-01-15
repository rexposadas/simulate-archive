package http

import (
	"net/http"
	"net/url"
	"time"

	"github.com/rexposadas/simulate"
)

// MakeRequest makes an HTTP request.
// The caller is in charge of closing the response body. ( #todo: is this proper? )
func MakeRequest(req *http.Request) (*SimResponse, error) {

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		// todo: add time to series
		simulate.Metrics.Error(err, "failed to make a request")
		return nil, err
	}
	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	simulate.Metrics.TrackResponse(req, since)

	if resp.StatusCode != 200 {
		simulate.Metrics.Error(nil, resp.Status)
	}

	return r, nil
}

// Get, runs a simple GET request on the specified URL.
// The caller is in charge of closing the response body. ( #todo: is this proper? )

func Get(url string) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		simulate.Metrics.Error(err)
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	if resp.StatusCode != 200 {
		simulate.Metrics.Error(nil, string(resp.StatusCode))
	}

	return r, nil
}

// The caller is in charge of closing the response body. ( #todo: is this proper? )
func Post(url string, payload url.Values) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
	if err != nil {
		simulate.Metrics.Error(err, "Failed to post form")

		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	if resp.StatusCode != 200 {
		simulate.Metrics.Error(nil, string(resp.StatusCode))
	}
	return r, nil
}
