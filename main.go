package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

var (
	tokenRegex = regexp.MustCompile("^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}$")
)

func IsToken(s string) bool {
	return tokenRegex.MatchString(s)
}

func NewToken() string {
	return uuid.NewV4().String()
}

// TODO: Check whether NaN, Inf, etc. are accepted in ParseFloat.
// TODO: Check exactly 6 digits after decimal point.
// -90 (South Pole) <= latitude <= +90 (North Pole)
func IsLatitude(s string) bool {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -90 || f > 90 {
		return false
	}
	return true
}

// TODO: Check whether NaN, Inf, etc. are accepted in ParseFloat.
// TODO: Check exactly 6 digits after decimal point.
// -180 (West) <= longitude <= +180 (East)
func IsLongitude(s string) bool {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -180 || f > 180 {
		return false
	}
	return true
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/route", SubmitStartPointAndDropOffLocations).Methods("POST")
	router.HandleFunc("/route/{token}", GetShortestDrivingRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

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

// SubmitStartPointAndDropOffLocations deals with the request "POST /route".
func SubmitStartPointAndDropOffLocations(w http.ResponseWriter, req *http.Request) {
	var locations Locations
	// --- Copied from http.Request.Body ---
	// For server requests the Request Body is always non-nil
	// but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP
	// Handler does not need to.
	err := json.NewDecoder(req.Body).Decode(&locations)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: "invalid input body (invalid JSON)"})
		return
	}

	if len(locations) < 2 {
		// At least 2 locations are required:
		// - route start location
		// - route end location
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: "insufficient number of locations"})
		return
	}
	for i, loc := range locations {
		if len(loc) != 2 {
			var msg string
			if i == 0 {
				msg = "invalid route start location"
			} else {
				msg = fmt.Sprintf("invalid drop off location #%d", i)
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct {
				Error string `json:"error"`
			}{Error: msg})
			return
		}
		if !IsLatitude(loc[0]) {
			var msg string
			if i == 0 {
				msg = fmt.Sprintf("invalid route start latitude: %s", loc[0])
			} else {
				msg = fmt.Sprintf("invalid drop off latitude #%d: %s", i, loc[0])
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct {
				Error string `json:"error"`
			}{Error: msg})
			return
		}
		if !IsLongitude(loc[1]) {
			var msg string
			if i == 0 {
				msg = fmt.Sprintf("invalid route start longitude: %s", loc[1])
			} else {
				msg = fmt.Sprintf("invalid drop off longitude #%d: %s", i, loc[1])
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct {
				Error string `json:"error"`
			}{Error: msg})
			return
		}
	}
	fmt.Println(locations)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: NewToken()})
}

// GetShortestDrivingRoute deals with the request "GET /route/{token}".
func GetShortestDrivingRoute(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	token, ok := params["token"]
	if !ok {
		// This section deals with the logic when the request does not contain
		// {token}. It should be impossible to reach here since the engine
		// should not call this method if {token} doesn't exist.
		// TODO: Should log this!
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: "internal server error (539cd7a5469b42ec1a53306df7fb2495)"})
		return
	}
	if !IsToken(token) {
		fmt.Printf("Bad Token: %s\n", token)
		// Return error
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: "invalid token"})
		return
	}
	fmt.Printf("Token: %s\n", token)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	r := rand.Float64()
	if r < 1.0/3.0 {
		json.NewEncoder(w).Encode(DrivingRoute{
			Status: "success",
			Path: [][]string{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
				{"22.326442", "114.167811"},
			},
			TotalDistance: 123,
			TotalTime:     456,
		})
	} else if r < 2.0/3.0 {
		json.NewEncoder(w).Encode(DrivingRoute{
			Status: "in progress",
		})
	} else {
		json.NewEncoder(w).Encode(DrivingRoute{
			Status: "failure",
			Error:  "unknown error",
		})
	}
}
