package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	// Check whether "-logtostderr=true" or "-logtostderr=false" is provided in
	// command line. If yes, obey the command line option. Otherwise, use the
	// default, "true".
	isLogToStderrProvided := false
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "logtostderr" {
			isLogToStderrProvided = true
		}
	})
	if !isLogToStderrProvided {
		flag.Set("logtostderr", "true")
	}

	// Setup and start an HTTP server.
	router := mux.NewRouter()
	router.HandleFunc("/route", SubmitStartPointAndDropOffLocations).Methods("POST")
	router.HandleFunc("/route/{token}", GetShortestDrivingRoute).Methods("GET")
	glog.Fatal(http.ListenAndServe(":8080", router))
}
