package simulate

import (
	"fmt"
	"net/http"
	"time"
)

type PrintStats struct {
	Count *Counter // non error counters
}

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

func (p *PrintStats) Error(err error, msg string) {
	if err != nil {
		fmt.Printf("%+v - %s", err, msg)
	}
}

// Print prints counter data to stdout
// Satisfies the Stats interface
func (p *PrintStats) Send() {
	p.Count.Lock()
	defer p.Count.Unlock()

	for k, v := range p.Count.M {
		fmt.Printf("\nkey: %s response: %v", k, v)
	}
}

// TrackResponse tracks api response times.
func (p *PrintStats) TrackResponse(req *http.Request, duration time.Duration) {
	fmt.Printf("\n %s url: %s took: %v", req.Method, req.URL.String(), duration)
}

func (p *PrintStats) Tally(key string, count int) {
	go p.Count.Add(key, count)
}

// Add adds a one (+1) to the Map for a given key
func (p *PrintStats) Add(t string) {
	go p.Count.Add(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (p *PrintStats) Sub(t string) {
	go p.Count.Add(t, -1)
}
