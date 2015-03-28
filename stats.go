package simulate

import (
	"time"

	"github.com/franela/goreq"
)

type Stats interface {
	// TrackResponse tracks how long a request took
	TrackResponse(req *goreq.Request, duration time.Duration)

	Tally(key string, count int)
	Error(err error, msg string)
	Send()
	Run()
}
