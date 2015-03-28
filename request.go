package simulate

import (
	"time"

	"github.com/franela/goreq"
)

// MakeRequest makes an HTTP request.
// The caller is in charge of closing the response body. ( #todo: is this proper? )
func MakeRequest(req goreq.Request) (*Response, error) {

	start := time.Now()

	resp, err := req.Do()
	if err != nil {
		Metrics.Error(err, "failed to make a request")
		return nil, err
	}

	since := time.Since(start)
	r := &Response{
		Duration: since,
		Response: resp,
	}

	Metrics.TrackResponse(&req, since)

	if resp.StatusCode != 200 {
		Metrics.Error(nil, string(resp.StatusCode))
	}

	return r, nil
}
