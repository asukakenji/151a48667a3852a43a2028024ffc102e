package main

import (
	"flag"
	"os"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/webtier2/lib"
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
				glog.Errorf("Cannot fetch question")
				return _err
			}
			glog.Infof("qid: %d", qid)

			gid, _err := taskqueue.RegisterGarbage(conn, qid, q.Token)
			if _err != nil {
				glog.Errorf("Cannot register garbage")
				return _err
			}
			glog.Info("gid: %d", gid)

			aidp, a, _err := taskqueue.GetAnswer2(conn, q.Token)
			if _err != nil {
				glog.Errorf("Cannot get answer")
				return _err
			}
			glog.Infof("aidp: %d", aidp)

			if a != nil {
				if a.QuestionID != qid {
					// TODO: Clear answer now
				}

				tc = a.TrialCount + 1
				if tc >= maxTrialCount {
					// TODO: return error
				}
			}

			aidip, _err := taskqueue.SetAnswer(conn, qid, tc, q.Token, &common.DrivingRoute{
				Status: "in progress",
			})
			if _err != nil {
				glog.Errorf("Cannot set answer (in progress)")
				return _err
			}
			glog.Infof("aid (in progress): %d", aidip)

			glocs := lib.LocationsToGoogleMapsLocations(q.Locations)

			dmr, _err := lib.GetDistanceMatrix(apiKey, glocs)
			if _err != nil {
				glog.Errorf("Cannot get distance matrix")
				return _err
			}

			m, _err := lib.GoogleMapsMatrixToMatrix(dmr)
			if _err != nil {
				glog.Errorf("Cannot convert Google Maps Matrix to Matrix")
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

			locationPath := make([][]string, size)
			for i, index := range path {
				locationPath[i] = q.Locations[index]
			}

			totalTime := lib.CalculateTotalTime(dmr, path)

			dr := &common.DrivingRoute{
				Status:        "success",
				Path:          common.Locations(locationPath),
				TotalDistance: cost,
				TotalTime:     int(totalTime.Seconds()),
			}

			aids, _err := taskqueue.SetAnswer(conn, qid, tc, q.Token, dr)
			if _err != nil {
				glog.Errorf("Cannot set answer (success)")
				return _err
			}
			glog.Infof("aid (success): %d", aids)

			return nil
		})
		if err != nil {
			glog.Errorf("main: %#v", err)
			err2 := taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
				aidf, _err := taskqueue.SetAnswer(conn, qid, tc, q.Token, &common.DrivingRoute{
					Status: "failure",
					Error:  err.Error(),
				})
				if _err != nil {
					glog.Errorf("Cannot set answer (failure)")
					return _err
				}
				glog.Infof("aid (failure): %d", aidf)

				return nil
			})
			if err2 != nil {
				break
			}
		}
	}
}
