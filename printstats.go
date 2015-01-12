package simulate

import (
	"fmt"
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

func (p *PrintStats) Error(err error, msg ...string) {
	if err != nil {
		fmt.Printf("%+v", err)
	}
	for m := range msg {
		fmt.Printf("\n%s", m)
	}
}

// Print prints counter data to stdout
// Satisfies the Stats interface
func (p *PrintStats) Send() {
	p.Count.Lock()
	defer p.Count.Unlock()

	for k, v := range p.Count.M {
		fmt.Printf("key: `%s` response: `%v`", k, v)
	}
}

// TrackResponse tracks api response times.
func (p *PrintStats) TrackResponse(url string, duration time.Duration) {
	fmt.Printf("\nurl: %s duration: %v", url, duration)
}

// Add adds a one (+1) to the Map for a given key
func (p *PrintStats) Add(t string) {
	go p.Count.Add(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (p *PrintStats) Sub(t string) {
	go p.Count.Add(t, -1)
}
