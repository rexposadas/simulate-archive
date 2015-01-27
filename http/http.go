package http

import (
	"bytes"
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

	defer resp.Body.Close()

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	simulate.Metrics.TrackResponse(req.URL.String(), since)

	if resp.StatusCode != 200 {
		simulate.Metrics.Error(nil, resp.Status)
	}

	return r, nil
}

func Get(url string) (*SimResponse, error) {

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Post(url string, payload url.Values) (*SimResponse, error) {

	req, _ := http.NewRequest("PostForm", url, bytes.NewBufferString(payload.Encode()))

	resp, err := MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
