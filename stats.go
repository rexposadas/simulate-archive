package simulate

import (
	"net/http"
	"time"
)

type Stats interface {
	// TrackResponse tracks how long a request took
	TrackResponse(req *http.Request, duration time.Duration)

	Tally(key string, count int)
	Error(err error, msg string)
	Send()
	Run()
}
