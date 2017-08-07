package common

type Locations [][]string

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

type Question struct {
	Token     string    `json:"token"`
	Locations Locations `json:"locations"`
}

type Answer struct {
	Timestamp    int64         `json:"timestamp"`
	DrivingRoute *DrivingRoute `json:"driving_route"`
}
