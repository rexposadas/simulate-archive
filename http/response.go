package http

import (
	"time"
)

// Response contains informatino about the response we received
// from the API
type Response struct {
	Duration time.Duration // how long did we wait for the response
}
