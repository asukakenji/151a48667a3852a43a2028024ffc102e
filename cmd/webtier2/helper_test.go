package main

import (
	"reflect"
	"testing"
	"time"

	"googlemaps.github.io/maps"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
)

var (
	dmr1 = &maps.DistanceMatrixResponse{
		OriginAddresses: []string{
			"11 Hoi Shing Rd, Tsuen Wan, Hong Kong",
			"Laguna City, Central, Hong Kong",
			"789 Nathan Rd, Mong Kok, Hong Kong",
		},
		DestinationAddresses: []string{
			"11 Hoi Shing Rd, Tsuen Wan, Hong Kong",
			"Laguna City, Central, Hong Kong",
			"789 Nathan Rd, Mong Kok, Hong Kong",
		},
		Rows: []maps.DistanceMatrixElementsRow{
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 38000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          995000000000,
						DurationInTraffic: 1040000000000,
						Distance: maps.Distance{
							HumanReadable: "15.5 km",
							Meters:        15518,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          845000000000,
						DurationInTraffic: 925000000000,
						Distance: maps.Distance{
							HumanReadable: "9.7 km",
							Meters:        9667,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          878000000000,
						DurationInTraffic: 914000000000,
						Distance: maps.Distance{
							HumanReadable: "15.2 km",
							Meters:        15223,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 3000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          883000000000,
						DurationInTraffic: 932000000000,
						Distance: maps.Distance{
							HumanReadable: "8.3 km",
							Meters:        8333,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          816000000000,
						DurationInTraffic: 788000000000,
						Distance: maps.Distance{
							HumanReadable: "10.3 km",
							Meters:        10329,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          908000000000,
						DurationInTraffic: 938000000000,
						Distance: maps.Distance{
							HumanReadable: "8.5 km",
							Meters:        8464,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 1000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
				},
			},
		},
	}
)

func TestLocationsToGoogleMapsLocations(t *testing.T) {
	cases := []struct {
		locs     common.Locations
		expected []string
	}{
		{
			nil,
			[]string{},
		},
		{
			common.Locations{},
			[]string{},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
			},
			[]string{
				"22.372081,114.107877",
			},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
			},
			[]string{
				"22.372081,114.107877",
				"22.284419,114.159510",
			},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
				{"22.326442", "114.167811"},
			},
			[]string{
				"22.372081,114.107877",
				"22.284419,114.159510",
				"22.326442,114.167811",
			},
		},
	}
	for _, c := range cases {
		got := LocationsToGoogleMapsLocations(c.locs)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf(
				"LocationsToGoogleMapsLocations(%v) = %v, expected %v",
				c.locs, got, c.expected,
			)
		}
	}
}

func TestGoogleMapsMatrixToMatrix(t *testing.T) {
	cases := []struct {
		dmr      *maps.DistanceMatrixResponse
		expected matrix.Matrix
	}{
		{
			nil,
			matrix.NewSquareMatrix([][]int{}),
		},
		{
			dmr1,
			matrix.NewSquareMatrix([][]int{
				{constant.Infinity, 15518, 9667},
				{15223, constant.Infinity, 8333},
				{10329, 8464, constant.Infinity},
			}),
		},
	}
	for _, c := range cases {
		got := GoogleMapsMatrixToMatrix(c.dmr)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf(
				"GoogleMapsMatrixToMatrix(%v) = %v, expected %v",
				c.dmr, got, c.expected,
			)
		}
	}
}

func TestCalculateTotalTime(t *testing.T) {
	cases := []struct {
		dmr      *maps.DistanceMatrixResponse
		path     []int
		expected time.Duration
	}{
		{nil, nil, 0},
		{dmr1, []int{0, 2, 1}, 1753000000000},
	}
	for _, c := range cases {
		got := CalculateTotalTime(c.dmr, c.path)
		if got != c.expected {
			t.Errorf(
				"CalculateTotalTime(%v, %v) = %d, expected %d",
				c.dmr, c.path, got, c.expected,
			)
		}
	}
}
