package simulate

import (
	"fmt"
	"time"

	"github.com/influxdb/influxdb/client"
)

// Stats goes to influxDB
type InfluxDBStats struct {
	Database string
	Count    *Counter
	Client   *client.Client
}

func (i *InfluxDBStats) mustDB() {

	d := &client.ClientConfig{}

	c, err := client.New(d)
	if err != nil {
		panic(err)
	}

	dbs, err := c.GetDatabaseList()
	if err != nil {
		panic(err)
	}

	if len(dbs) > 0 {
		c.DeleteDatabase(i.Database)
	}

	if err = c.CreateDatabase(i.Database); err != nil {
		panic(err)
	}

	i.Client, err = client.NewClient(&client.ClientConfig{
		Database: i.Database,
	})

	if err != nil {
		panic(err)
	}
}

func (i *InfluxDBStats) Run() {
	i.Count = NewCounter()
	i.mustDB()
	go i.Send()
}

// TrackResponse sends response times to influxDB
func (i *InfluxDBStats) TrackResponse(url string, duration time.Duration) {
	u := &client.Series{
		Name:    "api",
		Columns: []string{"url", "response_time"},
		Points: [][]interface{}{
			{url, duration.Seconds() * 1000},
		},
	}

	if err := i.Client.WriteSeries([]*client.Series{u}); err != nil {
		panic(err)
	}
}

// sendCounts sends count stats to influxDB
func sendCounts(c *client.Client, series string, count *Counter) {
	count.Lock()
	defer count.Unlock()

	for k, v := range count.M {
		u := &client.Series{
			Name:    series,
			Columns: []string{"key", "count"},
			Points: [][]interface{}{
				{k, v},
			},
		}

		fmt.Printf("adding tick %+v\n", u)
		if err := c.WriteSeries([]*client.Series{u}); err != nil {
			panic(err)
		}
	}
}

// Send periodically send stats to influxDB
func (i *InfluxDBStats) Send() {

	t := time.NewTicker(time.Second * 5)
	for {
		<-t.C
		go sendCounts(i.Client, "counts", i.Count)
	}
}

func (i *InfluxDBStats) Tally(t string, c int) {
	i.Count.Add(t, c)
}

// Add adds a one (+1) to the Map for a given key
func (i *InfluxDBStats) Add(t string) {
	go i.Tally(t, 1)
}

// Add substracts one (-1) to the Map for a given key
func (i *InfluxDBStats) Sub(t string) {
	go i.Tally(t, -1)
}
