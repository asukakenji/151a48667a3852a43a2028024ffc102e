package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
