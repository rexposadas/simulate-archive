package simulate

import (
	"fmt"
	"time"

	"github.com/influxdb/influxdb/client"
)

type StatsObj struct {
	Count  map[string]int
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

	fmt.Printf("config being used is %+v", c)

	s.Client = c
}

func (s *StatsObj) Run() {
	s.Count = make(map[string]int)
	s.MChan = make(chan CountData, 1000)
	s.mustDB()

	for {
		m := <-s.MChan

		if v, ok := s.Count[m.name]; ok {
			s.Count[m.name] += v
		} else {
			s.Count[m.name] = 1
		}
	}
}

func (s *StatsObj) TrackResponse(url string, duration time.Duration) {
	u := &client.Series{
		Name:    "apireponse",
		Columns: []string{"url", "response_time"},
		Points: [][]interface{}{
			{url, duration},
		},
	}

	if err := s.Client.WriteSeries([]*client.Series{u}); err != nil {
		panic(err)
	}
}

func (s *StatsObj) Tick(t string, c int) {

	d := CountData{
		name:  t,
		value: c,
	}
	s.MChan <- d
}

// Add adds a one (+1) to the Map for a given key
func (s *StatsObj) Add(t string) {
	s.Tick(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (s *StatsObj) Sub(t string) {
	s.Tick(t, -1)
}
