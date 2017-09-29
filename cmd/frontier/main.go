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

var (
	config *Config
)

func init() {
	config = &Config{
		ListenAddress:             ":8080",
		TaskQueueAddress:          "127.0.0.1:11300",
		TimeLimitForFindingAnswer: 5 * time.Second,
	}
	flag.StringVar(&config.ListenAddress, "listen", config.ListenAddress, "The address to which the server listens")
	flag.StringVar(&config.TaskQueueAddress, "queue", config.TaskQueueAddress, "The address of the task queue")
	flag.Int64Var((*int64)(&config.TimeLimitForFindingAnswer), "timeLimitA", int64(config.TimeLimitForFindingAnswer), "The time limit for finding the answer in nanoseconds")
}

func ReadConfig(config *Config) {
	file, err := os.Open("frontier.json")
	if err != nil {
		glog.Infof(`Cannot open "frontier.json": %s`, err.Error())
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		glog.Errorf(`Failed to read "frontier.json: %s"`, err.Error())
	}
}

func main() {
	// Read the config file
	ReadConfig(config)

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
	router.HandleFunc("/route", func(w http.ResponseWriter, req *http.Request) {
		SubmitStartPointAndDropOffLocations(config, w, req)
	}).Methods("POST")
	router.HandleFunc("/route/{token}", func(w http.ResponseWriter, req *http.Request) {
		GetShortestDrivingRoute(config, w, req)
	}).Methods("GET")
	glog.Fatal(http.ListenAndServe(config.ListenAddress, router))
}
