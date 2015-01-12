package simulate

import (
	"time"
)

type Stats interface {
	TrackResponse(url string, duration time.Duration)
	Error(err error, msg ...string)
	Send()
	Run()
}
