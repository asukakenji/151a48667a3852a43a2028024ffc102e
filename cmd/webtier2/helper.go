package main

import (
	"fmt"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"googlemaps.github.io/maps"
)

/*
LocationsToGoogleMapsLocations is a helper function which transforms the parameter
into a value that fits the Google Maps API.

It takes a common.Locations value as input,
and returns a slice of string as output.
Each element in the output slice is a latitude-longitude pair separated by a comma.
For example, an element {"22.372081", "114.107877"} in the input is transformed
into an element "22.372081,114.107877" in the output.
*/
func LocationsToGoogleMapsLocations(locs common.Locations) []string {
	glocs := make([]string, len(locs))
	for i, loc := range locs {
		glocs[i] = fmt.Sprintf("%s,%s", loc[0], loc[1])
	}
	return glocs
}

/*
GoogleMapsMatrixToMatrix is a helper function which transforms the parameter -
a value returned from the Google Maps API -
into a value that fits the internal API of this project.

It takes a *maps.DistanceMatrixRequest value as input,
and returns a matrix.Matrix as output.
The maps.DistanceMatrixRequest type is equivalent to the following:

	type DistanceMatrixResponse struct {
		OriginAddresses      []string
		DestinationAddresses []string
		Rows                 []struct {
			Elements []*struct {
				Status            string
				Duration          time.Duration
				DurationInTraffic time.Duration
				Distance          struct {
					HumanReadable string
					Meters        int
				}
			}
		}
	}

The value m.Get(r, c) in the output m refers to the field
dmr.Rows[r].Elements[c].Distance.Meters in the input dmr,
except when r == c, where m.Get(r, c) equals constant.Infinity
to fit the Travelling Salesman problem.
*/
func GoogleMapsMatrixToMatrix(dmr *maps.DistanceMatrixResponse) (matrix.Matrix, error) {
	if dmr == nil {
		return matrix.NewSquareMatrix([][]int{}), nil
	}
	m := make([][]int, len(dmr.Rows))
	for r, row := range dmr.Rows {
		m[r] = make([]int, len(row.Elements))
		for c, element := range row.Elements {
			if element.Status != "OK" {
				switch status := element.Status; status {
				case "NOT_FOUND":
					return nil, LocationNotFoundError{dmr, r, c}
				case "ZERO_RESULTS":
					return nil, RouteNotFoundError{dmr, r, c}
				default:
					return nil, common.WrapError(fmt.Errorf("Unknown error: Status: %q", status), "f8de639fab2bc7ab65c3153df6b8ee9e")
				}
			}
			if r == c {
				m[r][c] = constant.Infinity
				continue
			}
			m[r][c] = element.Distance.Meters
		}
	}
	return matrix.NewSquareMatrix(m), nil
}

/*
CalculateTotalTime is a helper function which calculates
the estimated total time needed for driving along the given path
from the information retrieved from Google Maps API.

It takes a *maps.DistanceMatrixResponse value and a path as input,
and returns the estimated total time needed as output.
The field used to calculate the total time is
dmr.Rows[from].Elements[to].Duration,
where from and to are the indices of the locations.
*/
func CalculateTotalTime(dmr *maps.DistanceMatrixResponse, path []int) time.Duration {
	totalTime := time.Duration(0)
	size := len(path)
	for i := 1; i < size; i++ {
		totalTime += dmr.Rows[path[i-1]].Elements[path[i]].Duration
	}
	return totalTime
}
