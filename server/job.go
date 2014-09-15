package server

import (
	"time"
)

const (
	METHOD_GET = "GET"
)

type Job struct {
	URL      string
	Method   string
	Interval time.Duration
}

func NewJob() *Job {
	return &Job{
		Interval: time.Second * 2,
		Method:   METHOD_GET,
	}
}
