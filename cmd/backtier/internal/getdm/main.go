package main

import (
	"fmt"
	"os"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/backtier/lib"

	"googlemaps.github.io/maps"
)

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

	dmr, err := lib.GetDistanceMatrix(apiKey, glocs)
	if err != nil {
		panic(err)
	}

	DumpDMR(dmr)
}
