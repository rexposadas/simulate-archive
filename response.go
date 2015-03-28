package simulate

import (
	"time"

	"github.com/franela/goreq"
)

// Response contains information about the response we received
// from the API
type Response struct {
	*goreq.Response
	Duration time.Duration // how long did we wait for the response
}
