package common

import (
	"encoding/json"
	"io"
)

// Locations is a list (slice) of latitude-longitude pairs.
//
// The minimum length of a Locations is 2:
// the first element being the route start location,
// and the second element being the route end location.
//
// Each element in a Locations value is a latitude-longitude pair encoded as a slice of strings.
// Therefore, its length is exactly 2:
// the first element being the latitude,
// and the seconds element being the longitude.
//
// Example
//
// Here is an example of a valid Locations:
//     var locs Locations = Locations{
//         {"22.372081", "114.107877"},
//         {"22.284419", "114.159510"},
//         {"22.326442", "114.167811"},
//     }
type Locations [][]string

// LocationsFromJSON reads from r, parses the content as JSON, and returns the
// Locations value.
func LocationsFromJSON(r io.Reader) (locs Locations, err error) {
	err = json.NewDecoder(r).Decode(&locs)
	if err != nil {
		return nil, JSONDecodeError{err}
	}

	if len(locs) < 2 {
		return nil, InsufficientLocationCountError{locs}
	}

	for i, loc := range locs {
		if len(loc) != 2 {
			return nil, InvalidLocationError{locs, i}
		}
		if !IsLatitude(loc[0]) {
			return nil, LatitudeError{locs, i}
		}
		if !IsLongitude(loc[1]) {
			return nil, LongitudeError{locs, i}
		}
	}
	return locs, nil
}

type DrivingRoute struct {
	// Status is the status of the response.
	// It is either "success", "in progress", or "failure".
	Status string `json:"status"`

	// Path is the path of the shortest driving route.
	Path Locations `json:"path,omitempty"`

	// TotalDistance is the total driving distance (in meters) of the path above.
	// The circumference of the Earth is around 40,075km = 40,075,000m.
	// So, uint should be sufficient no matter it is 32-bit or 64-bit.
	TotalDistance uint `json:"total_distance,omitempty"`

	// TotalTime is the estimated total time (in seconds) needed for driving along the path above.
	// An unsigned 32-bit integer can represent a duration of more than 136 years.
	// So, uint should be sufficient no matter it is 32-bit or 64-bit.
	TotalTime uint `json:"total_time,omitempty"`

	// Error is the error occurred during the calculation.
	Error string `json:"error,omitempty"`
}
