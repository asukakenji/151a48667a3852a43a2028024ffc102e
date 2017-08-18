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
	"github.com/kr/beanstalk"
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

	addr := "127.0.0.1:11300"    // TODO: Customize: addr
	timeLimit := 5 * time.Second // TODO: Customize: timeLimit
	apiKey := os.Args[1]         // TODO: Customize: apiKey
	maxTrialCount := 3           // TODO: Customize: maxTrialCount
	for {
		var qid uint64
		var q *taskqueue.Question
		var tc int
		err := taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
			var _err error
			qid, q, _err = taskqueue.FetchQuestion(conn, timeLimit)
			if _err != nil {
				glog.Errorf("main: cannot fetch question")
				return _err
			}
			glog.Infof("main: qid: %d", qid)

			gid, _err := taskqueue.RegisterGarbage(conn, q.Token, qid)
			if _err != nil {
				glog.Errorf("main: cannot register garbage")
				return _err
			}
			glog.Info("main: gid: %d", gid)

			aidp, a, _err := taskqueue.GetAnswer2(conn, q.Token)
			if _err != nil {
				if cerr, ok := _err.(beanstalk.ConnError); !ok {
					glog.Errorf("main: error occurred while getting answer")
					return _err
				} else if cerr.Err != beanstalk.ErrNotFound {
					glog.Errorf("main: error occurred while getting answer")
					return _err
				}
			}
			glog.Infof("main: aidp: %d", aidp)

			if a != nil {
				if a.QuestionID != qid {
					glog.Errorf("main: token collision")
					return lib.NewTokenCollisionError(q.Token, qid, a.QuestionID)
				}

				tc = a.TrialCount + 1
				if tc == maxTrialCount {
					glog.Errorf("main: trial count limit exceeded")
					return lib.NewTrialCountLimitExceededError(q.Token, qid, maxTrialCount)
				}
			}

			aidip, _err := taskqueue.SetAnswer(conn, q.Token, qid, tc, &common.DrivingRoute{
				Status: "in progress",
			})
			if _err != nil {
				glog.Errorf("main: cannot set answer (in progress)")
				return _err
			}
			glog.Infof("main: aid (in progress): %d", aidip)

			glocs := lib.LocationsToGoogleMapsLocations(q.Locations)

			dmr, _err := lib.GetDistanceMatrix(apiKey, glocs)
			if _err != nil {
				glog.Errorf("main: cannot get distance matrix")
				return _err
			}

			m, _err := lib.GoogleMapsMatrixToMatrix(dmr)
			if _err != nil {
				glog.Errorf("main: cannot convert Google Maps Matrix to Matrix")
				return _err
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
			totalTime := int(lib.CalculateTotalTime(dmr, path).Seconds())

			aids, _err := taskqueue.SetAnswer(conn, q.Token, qid, tc, &common.DrivingRoute{
				Status:        "success",
				Path:          locationPath,
				TotalDistance: cost,
				TotalTime:     totalTime,
			})
			if _err != nil {
				glog.Errorf("main: cannot set answer (success)")
				return _err
			}
			glog.Infof("main: aid (success): %d", aids)

			return nil
		})
		if err != nil {
			glog.Errorf("main: %#v", err)
			err2 := taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
				aidf, _err := taskqueue.SetAnswer(conn, q.Token, qid, tc, &common.DrivingRoute{
					Status: "failure",
					Error:  err.Error(),
				})
				if _err != nil {
					glog.Errorf("main: cannot set answer (failure)")
					return _err
				}
				glog.Infof("main: aid (failure): %d", aidf)

				return nil
			})
			if err2 != nil {
				break
			}
		}
	}
}
