package common

/*
Locations is a list (slice) of latitude-longitude pairs.

Description

The minimum length of a Locations is 2:
the first element being the route start location,
and the second element being the route end location.

Each element in a Locations value is a latitude-longitude pair encoded as a slice of strings.
Therefore, its length is exactly 2:
the first element being the latitude,
and the seconds element being the longitude.

A latitude value of 0 represents the Equator.
A latitude value of -90 represents the South Pole,
while +90 represents the North Pole.
Values smaller than -90 or larger than +90 are invalid.

A longitude value of 0 represents the prime meridian at Greenwich.
A negative value represents locations to the West of the prime meridian,
while a positive value represents locations to the East of the prime meridian.

Example

Here is an example of a valid Locations:

	var locs Locations = Locations{
		{"22.372081", "114.107877"}, // "11 Hoi Shing Rd, Tsuen Wan, Hong Kong"
		{"22.284419", "114.159510"}, // "Laguna City, Central, Hong Kong"
		{"22.326442", "114.167811"}, // "789 Nathan Rd, Mong Kok, Hong Kong"
	}
*/
type Locations [][]string

/*
DrivingRoute contains a driving route response.

Description

The field Status is mandatory.
Whether other fields exist depends on Status.
See "Examples" below for valid combinations.

The field Status is the status of the response.
It is either "success", "in progress", or "failure".

The field Path is the path of the shortest driving route.

The field TotalDistance is the total driving distance (in meters) of the path above.
Since the circumference of the Earth is around 40,000km = 40,000,000m at the Equator,
and a signed 32-bit integer can represent 2,147,483,647m,
more than 50,000 times the circumference of the Earth,
int should be sufficient no matter it is 32-bit or 64-bit.

The field TotalTime is the estimated total time (in seconds) needed for driving along the path above.
Since a signed 32-bit integer can represent a duration of more than 68 years,
int should be sufficient no matter it is 32-bit or 64-bit.

The field Error is the error occurred during the process.

Examples

Here is an example of a "success" DrivingRoute:

	var dr0 DrivingRoute = &DrivingRoute{
		Status: "success",
		Path: Locations{
			{"22.372081", "114.107877"}, // "11 Hoi Shing Rd, Tsuen Wan, Hong Kong"
			{"22.326442", "114.167811"}, // "789 Nathan Rd, Mong Kok, Hong Kong"
			{"22.284419", "114.159510"}, // "Laguna City, Central, Hong Kong"
		},
		TotalDistance: 20000,
		TotalTime:     1800,
	}

Here is an example of a "in progress" DrivingRoute:

	var dr1 DrivingRoute = &DrivingRoute{
		Status: "in progress",
	}

Here is an example of a "failure" DrivingRoute:

	var dr2 DrivingRoute = &DrivingRoute{
		Status: "failure",
		Error:  "internal server error (539cd7a5469b42ec1a53306df7fb2495)",
	}
*/
type DrivingRoute struct {
	Status        string    `json:"status"`
	Path          Locations `json:"path,omitempty"`
	TotalDistance int       `json:"total_distance,omitempty"`
	TotalTime     int       `json:"total_time,omitempty"`
	Error         string    `json:"error,omitempty"`
}
