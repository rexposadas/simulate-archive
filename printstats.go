package simulate

import (
	"fmt"
	"time"

	"github.com/franela/goreq"
)

// PrintStats holds a counter
type PrintStats struct {
	Count *Counter // non error counters
}

// NewPrintStats returns an intialized counter
func NewPrintStats() *PrintStats {
	return &PrintStats{
		Count: NewCounter(),
	}
}

// Run is a worker which sends stats periodically
// Satisfies the Stats interface
func (p *PrintStats) Run() {
	t := time.NewTicker(time.Second * 5)
	for {
		<-t.C
		go p.Send()
	}
}

// Error print errors to stdout
// Satisfies the Stats interface
func (p *PrintStats) Error(err error, msg string) {
	if err != nil {
		fmt.Printf("%+v - %s", err, msg)
	}
}

// Send sends counts to std. satisfies the Stats interface
func (p *PrintStats) Send() {
	p.Count.Lock()
	defer p.Count.Unlock()

	for k, v := range p.Count.M {
		fmt.Printf("\nkey: %s response: %v", k, v)
	}
}

// TrackResponse tracks api response times.
func (p *PrintStats) TrackResponse(req *goreq.Request, duration time.Duration) {
	fmt.Printf("\n %s url: %s took: %v", req.Method, req.Uri, duration)
}

// Tally adds to the key
// Satisfies the Stats interface
func (p *PrintStats) Tally(key string, count int) {
	go p.Count.Add(key, count)
}

// Add adds a one (+1) to the Map for a given key
func (p *PrintStats) Add(t string) {
	go p.Count.Add(t, 1)
}

// Sub substracts one (-1) to the Map for a given key
func (p *PrintStats) Sub(t string) {
	go p.Count.Add(t, -1)
}
