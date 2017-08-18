package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/bruteforce"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/tsp/heldkarp"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"

	"googlemaps.github.io/maps"
)

const (
	// heldKarpThreshold is determined by benchmarks.
	// When the problem size is less than heldKarpThreshold,
	// bruteforce performs better than heldkarp;
	// When the problem size is greater than or equal to heldKarpThreshold,
	// heldharp performs better.
	heldKarpThreshold = 8 // TODO: Customize: heldKarpThreshold
)

func main1() {
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

			glocs := LocationsToGoogleMapsLocations(q.Locations)

			dmr, _err := GetDistanceMatrix(apiKey, glocs)
			if _err != nil {
				glog.Errorf("Cannot get distance matrix")
				return _err
			}

			m, _err := GoogleMapsMatrixToMatrix(dmr)
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

			totalTime := CalculateTotalTime(dmr, path)

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

func GetDistanceMatrix(apiKey string, glocs []string) (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	r := &maps.DistanceMatrixRequest{
		Origins:                  glocs,
		Destinations:             glocs,
		Mode:                     maps.TravelModeDriving,
		Language:                 "",
		Avoid:                    maps.Avoid(""),
		Units:                    maps.UnitsMetric,
		DepartureTime:            "now",
		ArrivalTime:              "",
		TrafficModel:             maps.TrafficModel(""),
		TransitMode:              []maps.TransitMode(nil),
		TransitRoutingPreference: maps.TransitRoutingPreference(""),
	}
	resp, err := c.DistanceMatrix(ctx, r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DumpDMR(dmr *maps.DistanceMatrixResponse) {
	fmt.Printf("Origin Addresses:\n")
	for _, addr := range dmr.OriginAddresses {
		fmt.Printf("  %q\n", addr)
	}
	fmt.Printf("Destination Addresses:\n")
	for _, addr := range dmr.DestinationAddresses {
		fmt.Printf("  %q\n", addr)
	}
	fmt.Printf("Distance Matrix Elements Row:\n")
	for i, row := range dmr.Rows {
		fmt.Printf("  %d:\n", i)
		for j, elem := range row.Elements {
			fmt.Printf("    %d:\n", j)
			fmt.Printf("      Status: %q\n", elem.Status)
			fmt.Printf("      Duration: %d\n", elem.Duration)
			fmt.Printf("      DurationInTraffic: %d\n", elem.DurationInTraffic)
			fmt.Printf("      Distance:\n")
			fmt.Printf("        Human Readable: %q\n", elem.Distance.HumanReadable)
			fmt.Printf("        Meters: %d\n", elem.Distance.Meters)
		}
	}
}

func main() {
	apiKey := os.Args[1]
	glocs := os.Args[2:]

	dmr, err := GetDistanceMatrix(apiKey, glocs)
	if err != nil {
		panic(err)
	}

	DumpDMR(dmr)
}
