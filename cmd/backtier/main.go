package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/backtier/lib"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/bruteforce"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/heldkarp"
	"github.com/golang/glog"
)

type Config struct {
	MapProviderAPIKey               string        `json:"map_provider_api_key"`
	TaskQueueAddress                string        `json:"task_queue_address"`                   // Default: "127.0.0.1:11300"
	RetryCountLimitForFindingAnswer int           `json:"retry_count_limit_for_finding_answer"` // Default: 3
	TimeWaitedBeforeBecomingGarbage time.Duration `json:"time_waited_before_becoming_garbage"`  // Default: 10 * time.Minute
	TimeLimitForRemovingGarbage     time.Duration `json:"time_limit_for_removing_garbage"`      // Default: 1 * time.Minute
}

func ReadConfig() *Config {
	file, err := os.Open("backtier.json")
	if err != nil {
		glog.Fatalf(`Cannot open "backtier.json": %s`, err.Error())
		return nil
	}
	defer file.Close()

	var config *Config
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		glog.Fatalf(`Failed to read "backtier.json: %s"`, err.Error())
		return nil
	}
	return config
}

const (
	// heldKarpThreshold is determined by benchmarks.
	// When the problem size is less than heldKarpThreshold,
	// bruteforce performs better than heldkarp;
	// When the problem size is greater than or equal to heldKarpThreshold,
	// heldharp performs better.
	heldKarpThreshold = 8 // TODO: Customize: heldKarpThreshold
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

	// Read the config file
	config := ReadConfig()

	for {
		var qid uint64
		var q *taskqueue.Question
		var rc int
		err := taskqueue.WithConnection(config.TaskQueueAddress, func(conn *taskqueue.Connection) common.Error {
			var err2 common.Error
			for {
				qid, q, err2 = taskqueue.FetchQuestion(conn)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot fetch question", err2.Hash())
					return err2
				}
				glog.Infof("main: qid: %d", qid)

				gid, err2 := taskqueue.RegisterGarbage(
					conn,
					q.Token,
					qid,
					config.TimeWaitedBeforeBecomingGarbage,
					config.TimeLimitForRemovingGarbage,
				)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot register garbage", err2.Hash())
					return err2
				}
				glog.Info("main: gid: %d", gid)

				aidp, a, err2 := taskqueue.GetAnswer2(conn, q.Token)
				if err2 != nil {
					if _, ok := err2.(taskqueue.NotFoundError); !ok {
						glog.Errorf("[%s] main: error occurred while getting answer", err2.Hash())
						return err2
					}
				}
				glog.Infof("main: aidp: %d", aidp)

				if a != nil {
					if a.QuestionID != qid {
						hash := common.NewToken()
						glog.Errorf("[%s] main: token collision", hash)
						return lib.NewTokenCollisionError(q.Token, qid, a.QuestionID, hash)
					}

					rc = a.RetryCount + 1
					if rc == config.RetryCountLimitForFindingAnswer {
						hash := common.NewToken()
						glog.Errorf("[%s] main: retry count limit exceeded", hash)
						return lib.NewRetryCountLimitExceededError(q.Token, qid, rc, hash)
					}
				}

				aidip, err2 := taskqueue.SetAnswer(conn, q.Token, qid, rc, &common.DrivingRoute{
					Status: "in progress",
				})
				if err2 != nil {
					glog.Errorf("[%s] main: cannot set answer (in progress)", err2.Hash())
					return err2
				}
				glog.Infof("main: aid (in progress): %d", aidip)

				glocs := lib.LocationsToGoogleMapsLocations(q.Locations)

				resp, err2 := lib.GetDistanceMatrix(config.MapProviderAPIKey, glocs)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot get distance matrix", err2.Hash())
					return err2
				}

				m, err2 := lib.GoogleMapsMatrixToMatrix(resp)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot convert Google Maps Matrix to Matrix", err2.Hash())
					return err2
				}

				var cost int
				var path []int
				size := len(q.Locations)
				if size < heldKarpThreshold {
					cost, path = bruteforce.TravellingSalesmanPath(m)
				} else {
					cost, path = heldkarp.TravellingSalesmanPath(m)
				}
				glog.Infof("main: cost: %d, path: %v", cost, path)

				locationPath := lib.PathToLocationPath(q.Locations, path)
				totalTime := int(lib.CalculateTotalTime(resp, path).Seconds())

				aids, err2 := taskqueue.SetAnswer(conn, q.Token, qid, rc, &common.DrivingRoute{
					Status:        "success",
					Path:          locationPath,
					TotalDistance: cost,
					TotalTime:     totalTime,
				})
				if err2 != nil {
					glog.Errorf("[%s] main: cannot set answer (success)", err2.Hash())
					return err2
				}
				glog.Infof("main: aid (success): %d", aids)
			}
		})
		if err != nil {
			glog.Errorf("[%s] main: %#v", err.Hash(), err)
			taskqueue.WithConnection(config.TaskQueueAddress, func(conn *taskqueue.Connection) common.Error {
				aidf, err2 := taskqueue.SetAnswer(conn, q.Token, qid, rc, &common.DrivingRoute{
					Status: "failure",
					Error:  err.Error(),
				})
				if err2 != nil {
					glog.Errorf("[%s] main: cannot set answer (failure)", err2.Hash())
					return err2
				}
				glog.Infof("main: aid (failure): %d", aidf)

				return nil
			})
		}
	}
}
