package simulate

import (
	"fmt"
	"time"

	"github.com/influxdb/influxdb/client"
)

type Stats interface {
	TrackResponse(url string, duration time.Duration)
	Tally(t string, c int) // counter for non-error strings
	Run()
}
