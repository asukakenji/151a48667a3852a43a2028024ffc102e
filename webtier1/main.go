package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	isLogToStderrOverridden := false
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "logtostderr" {
			isLogToStderrOverridden = true
		}
	})
	if !isLogToStderrOverridden {
		flag.Set("logtostderr", "true")
	}

	router := mux.NewRouter()
	router.HandleFunc("/route", SubmitStartPointAndDropOffLocations).Methods("POST")
	router.HandleFunc("/route/{token}", GetShortestDrivingRoute).Methods("GET")
	glog.Fatal(http.ListenAndServe(":8080", router))
}
