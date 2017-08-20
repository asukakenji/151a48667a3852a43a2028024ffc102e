package main

import (
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

	addr := "127.0.0.1:11300" // TODO: Customize: addr
	apiKey := os.Args[1]      // TODO: Customize: apiKey
	maxRetryCount := 3        // TODO: Customize: maxRetryCount
	delay := 10 * time.Minute // TODO: Customize: delay
	ttr := 1 * time.Minute    // TODO: Customize: delay
	for {
		var qid uint64
		var q *taskqueue.Question
		var rc int
		err := taskqueue.WithConnection(addr, func(conn *taskqueue.Connection) common.Error {
			var err2 common.Error
			for {
				qid, q, err2 = taskqueue.FetchQuestion(conn)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot fetch question", err2.Hash())
					return err2
				}
				glog.Infof("main: qid: %d", qid)

				gid, err2 := taskqueue.RegisterGarbage(conn, q.Token, qid, delay, ttr)
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
					if rc == maxRetryCount {
						hash := common.NewToken()
						glog.Errorf("[%s] main: retry count limit exceeded", hash)
						return lib.NewRetryCountLimitExceededError(q.Token, qid, maxRetryCount, hash)
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

				resp, err2 := lib.GetDistanceMatrix(apiKey, glocs)
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
			taskqueue.WithConnection(addr, func(conn *taskqueue.Connection) common.Error {
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
