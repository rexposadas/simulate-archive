package http

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"github.com/rexposadas/simulate"
)

func MakeRequest(req *http.Request) (*SimResponse, error) {

	client := &http.Client{}
	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}
	return r, nil
}

// Get, runs a simple GET request on the specified URL.
func Get(url string) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	// stats check
	s := simulate.StatsObj{make(map[string]int)} 
	s.SimpleMath(url, 1)
	
	fmt.Printf("Getting of '%s' - response time %f seconds. \n\n", url, since.Seconds())

	return r, nil
}

func Post(url string, payload url.Values) (*SimResponse, error) {

	start := time.Now()
	resp, err := http.PostForm(url, payload)
	if err != nil {
		return nil, err
	}

	since := time.Since(start)
	r := &SimResponse{
		Duration: since,
		Response: resp,
	}

	// stats check
	s := simulate.StatsObj{make(map[string]int)} 
	s.SimpleMath(url, 1)
	fmt.Printf("%+v\n", s.Count)

	fmt.Printf("Post to '%s' - response time %f seconds. \n\n", url, since.Seconds())

	return r, nil
}
