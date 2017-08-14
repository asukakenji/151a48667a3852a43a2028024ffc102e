package main

import (
	"context"
	"fmt"
	"os"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"

	"googlemaps.github.io/maps"
)

func main1() {
	c, err := maps.NewClient(maps.WithAPIKey(os.Args[1]))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", c)
	ctx := context.Background()
	r := &maps.DistanceMatrixRequest{
		// Origins:                  []string{"22.372081,114.107877"},
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

func main() {
	// https://youtu.be/vNqE_LDTsa0
	m := matrix.NewSquareMatrix([][]int{
		{-1, 7, 6, 8, 4},
		{7, -1, 8, 5, 6},
		{6, 8, -1, 9, 7},
		{8, 5, 9, -1, 8},
		{4, 6, 7, 8, -1},
	})
	// https://youtu.be/FJkT_dRjX94
	// https://youtu.be/KzWC-t1y8Ac
	m = matrix.NewSquareMatrix([][]int{
		{-1, 20, 30, 10, 11},
		{15, -1, 16, 4, 2},
		{3, 5, -1, 2, 4},
		{19, 6, 18, -1, 3},
		{16, 4, 7, 16, -1},
	})
	cost, path := TravellingSalesmanNaive(m)
	fmt.Printf("Cost: %d\n", cost)
	fmt.Printf("Path: %v\n", path)
}
