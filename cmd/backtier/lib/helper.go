package lib

import (
	"fmt"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
	"github.com/golang/glog"
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

func PathToLocationPath(locs common.Locations, path []int) common.Locations {
	locationPath := make([][]string, len(locs))
	for i, index := range path {
		locationPath[i] = locs[index]
	}
	return locationPath
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
resp.Rows[r].Elements[c].Distance.Meters in the input resp,
except when r == c, where m.Get(r, c) equals constant.Infinity
to fit the Travelling Salesman problem.
*/
func GoogleMapsMatrixToMatrix(resp *maps.DistanceMatrixResponse) (matrix.Matrix, common.Error) {
	if resp == nil {
		return matrix.NewSquareMatrix([][]int{}), nil
	}
	m := make([][]int, len(resp.Rows))
	for r, row := range resp.Rows {
		m[r] = make([]int, len(row.Elements))
		for c, element := range row.Elements {
			if element.Status != "OK" {
				switch status := element.Status; status {
				case "NOT_FOUND":
					hash := common.NewToken()
					glog.Errorf("[%s] GoogleMapsMatrixToMatrix: NOT_FOUND", hash)
					return nil, NewLocationNotFoundError(resp, r, c, hash)
				case "ZERO_RESULTS":
					hash := common.NewToken()
					glog.Errorf("[%s] GoogleMapsMatrixToMatrix: ZERO_RESULTS", hash)
					return nil, NewRouteNotFoundError(resp, r, c, hash)
				default:
					fixedHash := "f8de639fab2bc7ab65c3153df6b8ee9e"
					glog.Errorf("[%s] GoogleMapsMatrixToMatrix: Unknown error: Status: %q", fixedHash, status)
					return nil, common.WrapError(fmt.Errorf("Unknown error: Status: %q", status), fixedHash)
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
resp.Rows[from].Elements[to].Duration,
where from and to are the indices of the locations.
*/
func CalculateTotalTime(resp *maps.DistanceMatrixResponse, path []int) time.Duration {
	totalTime := time.Duration(0)
	size := len(path)
	for i := 1; i < size; i++ {
		totalTime += resp.Rows[path[i-1]].Elements[path[i]].Duration
	}
	return totalTime
}
