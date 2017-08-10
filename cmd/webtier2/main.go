package main

import (
	"context"
	"fmt"
	"os"

	"googlemaps.github.io/maps"
)

func main() {
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
