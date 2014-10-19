package simulate

import (
	"time"
)

type call func() error

// Job is a job which simulate runs.  Jobs are sent to the Jobs channel
// for processing.
type Job struct {
	Interval time.Duration
	Run      call
}

// NewJob returns a job with a default values.
func NewJob() *Job {
	return &Job{
		Interval: time.Second * 2, // arbitrary interval to do the job.
	}
}
