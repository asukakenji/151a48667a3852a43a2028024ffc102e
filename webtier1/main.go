package main

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/route", SubmitStartPointAndDropOffLocations).Methods("POST")
	router.HandleFunc("/route/{token}", GetShortestDrivingRoute).Methods("GET")
	glog.Fatal(http.ListenAndServe(":8080", router))
}
