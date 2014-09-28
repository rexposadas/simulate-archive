package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StartAPI(port int) {

	r := mux.NewRouter()
	r.HandleFunc("/jobs", AddJobs).Methods("Post")

	http.Handle("/", r)

	fmt.Printf("Serving API in PORT %d", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}

}

// AddJobs adds a job in the job chanel.
// This is part of the on-the-fly simulation effort
func AddJobs(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	u := r.FormValue("url")

	Add(u)
}
