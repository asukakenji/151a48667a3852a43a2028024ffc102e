package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type Config struct {
	ListenAddress             string        `json:"listen_address"`                // Default: ":8080"
	TaskQueueAddress          string        `json:"task_queue_address"`            // Default: "127.0.0.1:11300"
	TimeLimitForFindingAnswer time.Duration `json:"time_limit_for_finding_answer"` // Default: 5 * time.Second
}

func ReadConfig() *Config {
	file, err := os.Open("frontier.json")
	if err != nil {
		glog.Fatalf(`Cannot open "frontier.json": %s`, err.Error())
		return nil
	}
	defer file.Close()

	var config *Config
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		glog.Fatalf(`Failed to read "frontier.json: %s"`, err.Error())
		return nil
	}
	return config
}

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

	// Read the config file
	config := ReadConfig()

	// Setup and start an HTTP server.
	router := mux.NewRouter()
	router.HandleFunc("/route", func(w http.ResponseWriter, req *http.Request) {
		SubmitStartPointAndDropOffLocations(config, w, req)
	}).Methods("POST")
	router.HandleFunc("/route/{token}", func(w http.ResponseWriter, req *http.Request) {
		GetShortestDrivingRoute(config, w, req)
	}).Methods("GET")
	glog.Fatal(http.ListenAndServe(config.ListenAddress, router))
}
