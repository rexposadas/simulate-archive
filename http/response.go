package http

import (
	"net/http"
	"time"
)

// Response contains informatino about the response we received
// from the API
type SimResponse struct {
	*http.Response
	Duration time.Duration // how long did we wait for the response
}
