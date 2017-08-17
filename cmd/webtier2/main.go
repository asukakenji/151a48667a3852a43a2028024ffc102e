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
	for {
		var q *taskqueue.Question
		err := taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
			qid, _q, _err := taskqueue.FetchQuestion(conn, timeLimit)
			if _err != nil {
				glog.Errorf("Cannot fetch question")
				return _err
			}
			glog.Infof("qid: %d", qid)
			q = _q

			aidip, _err := taskqueue.SetAnswer(conn, _q.Token, &common.DrivingRoute{
				Status: "in progress",
			})
			if _err != nil {
				glog.Errorf("Cannot set answer (in progress)")
				return _err
			}
			glog.Infof("aid (in progress): %d", aidip)

			dmr, _err := GetDistanceMatrix(apiKey, q.Locations)
			if _err != nil {
				glog.Errorf("Cannot get distance matrix")
				return _err
			}

			m := GoogleMapsMatrixToMatrix(dmr)

			var cost int
			var path []int
			size := len(q.Locations)
			if size < heldKarpThreshold {
				cost, path = bruteforce.TravellingSalesmanTour(m)
			} else {
				cost, path = heldkarp.TravellingSalesmanTour(m)
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
				TotalTime:     totalTime,
			}

			aids, _err := taskqueue.SetAnswer(conn, _q.Token, dr)
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
				aidf, _err := taskqueue.SetAnswer(conn, q.Token, &common.DrivingRoute{
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

func GetDistanceMatrix(apiKey string, locs common.Locations) (*maps.DistanceMatrixResponse, error) {
	origins := LocationsToGoogleMapsLocations(locs)
	destinations := origins

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	r := &maps.DistanceMatrixRequest{
		Origins:                  origins,
		Destinations:             destinations,
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

func main1() {
	c, err := maps.NewClient(maps.WithAPIKey(os.Args[1]))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", c)
	ctx := context.Background()
	r := &maps.DistanceMatrixRequest{
		Origins:                  []string{"22.372081,114.107877", "22.284419,114.159510", "22.326442,114.167811"},
		Destinations:             []string{"22.372081,114.107877", "22.284419,114.159510", "22.326442,114.167811"},
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
		panic(err)
	}
	fmt.Printf("Origin Addresses:\n")
	for _, addr := range resp.OriginAddresses {
		fmt.Printf("  %q\n", addr)
	}
	fmt.Printf("Destination Addresses:\n")
	for _, addr := range resp.DestinationAddresses {
		fmt.Printf("  %q\n", addr)
	}
	fmt.Printf("Distance Matrix Elements Row:\n")
	for i, row := range resp.Rows {
		fmt.Printf("  %d:\n", i)
		for j, elem := range row.Elements {
			duration := (elem.Duration.Seconds())
			durationInTraffic := (elem.DurationInTraffic.Seconds())
			fmt.Printf("    %d:\n", j)
			fmt.Printf("      Status: %q\n", elem.Status)
			fmt.Printf("      Duration: %f\n", duration)
			fmt.Printf("      DurationInTraffic: %f\n", durationInTraffic)
			fmt.Printf("      Distance: %d\n", elem.Distance.Meters)
		}
	}
	fmt.Printf("%#v\n\n", resp)
}
