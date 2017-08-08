package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// GetShortestDrivingRoute deals with the request "GET /route/{token}".
func GetShortestDrivingRoute(w http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	var myerr common.MyError

	vars := mux.Vars(req)
	token, ok := vars["token"]
	if !ok {
		// This section deals with the logic when the request does not contain
		// {token}. It should be impossible to reach here since the engine
		// should not call this method if {token} doesn't exist.
		status = http.StatusInternalServerError
		myerr = common.WrapError(nil, "539cd7a5469b42ec1a53306df7fb2495")
	} else if !common.IsToken(token) {
		status = http.StatusBadRequest
		myerr = common.NewInvalidTokenError(token)
	}
	if status != http.StatusOK {
		glog.Errorf("GetShortestDrivingRoute: %s", myerr.Error())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: myerr.String()})
		return
	}

	glog.Infof("GetShortestDrivingRoute: token: %q", token)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO: Implement this!
	r := rand.Float64()
	if r < 1.0/3.0 {
		json.NewEncoder(w).Encode(common.DrivingRoute{
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
		json.NewEncoder(w).Encode(common.DrivingRoute{
			Status: "in progress",
		})
	} else {
		json.NewEncoder(w).Encode(common.DrivingRoute{
			Status: "failure",
			Error:  "unknown error",
		})
	}
}
