package simulate

import (
	"fmt"
	"time"

	"github.com/influxdb/influxdb/client"
)

type StatsObj struct {
	Count  *Counter
	MChan  chan CountData
	Client *client.Client
}

type CountData struct {
	name  string
	value int
}

func (s *StatsObj) mustDB() {

	d := &client.ClientConfig{
		Database: "leaf",
	}

	c, err := client.New(d)
	if err != nil {
		panic(err)
	}

	s.Client = c
}

func (s *StatsObj) Run() {
	s.MChan = make(chan CountData, 1000)
	s.Count = &Counter{}
	s.Count.M = make(map[string]int)
	s.mustDB()
	go s.SendStats()

	for {
		m := <-s.MChan
		s.Count.Add(m.name, m.value)
	}
}

// TrackResponse sends response times to influxDB
func (s *StatsObj) TrackResponse(url string, duration time.Duration) {
	u := &client.Series{
		Name:    "api",
		Columns: []string{"url", "response_time"},
		Points: [][]interface{}{
			{url, duration.Seconds() * 1000},
		},
	}

	if err := s.Client.WriteSeries([]*client.Series{u}); err != nil {
		panic(err)
	}
}

// SendStats periodically send stats to influxDB
func (s *StatsObj) SendStats() {

	t := time.NewTicker(time.Second * 5)
	for {
		<-t.C
		s.Count.lock.Lock()

		for k, v := range s.Count.M {
			u := &client.Series{
				Name:    "counts",
				Columns: []string{"key", "count"},
				Points: [][]interface{}{
					{k, v},
				},
			}

			fmt.Printf("adding tick %+v\n", u)
			if err := s.Client.WriteSeries([]*client.Series{u}); err != nil {
				panic(err)
			}
		}
		s.Count.lock.Unlock()
	}
}

func (s *StatsObj) Tick(t string, c int) {

	d := CountData{
		name:  t,
		value: c,
	}
	fmt.Printf("adding count %+v\n\n", d)
	s.MChan <- d

}

// Add adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {
	go s.Tick(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (s *StatsObj) Sub(t string) {
	go s.Tick(t, -1)
}
