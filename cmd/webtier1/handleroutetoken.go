package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/kr/beanstalk"
)

// GetShortestDrivingRoute deals with the request "GET /route/{token}".
func GetShortestDrivingRoute(w http.ResponseWriter, req *http.Request) {
	addr := "127.0.0.1:11300" // TODO: Customize: addr
	status := http.StatusOK
	var id uint64
	var dr *common.DrivingRoute
	var err error

	vars := mux.Vars(req)
	token, ok := vars["token"]
	if !ok {
		// This section deals with the logic when the request does not contain
		// {token}. It should be impossible to reach here since the engine
		// should not call this method if {token} doesn't exist.
		status = http.StatusInternalServerError
		err = common.WrapError(nil, "539cd7a5469b42ec1a53306df7fb2495")
		goto HandleError
	} else if !common.IsToken(token) {
		status = http.StatusBadRequest
		err = common.NewInvalidTokenError(token)
		goto HandleError
	}

	err = taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
		_id, _dr, _err := taskqueue.GetAnswer(conn, token)
		id, dr = _id, _dr
		return _err
	})
	if err != nil {
		status = http.StatusInternalServerError
		goto HandleError
	}

	glog.Infof("GetShortestDrivingRoute: id: %d, token: %q, driving_route: %#v", id, token, dr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK) // TODO: Should depend on the status of dr
	json.NewEncoder(w).Encode(dr)
	return

HandleError:
	if myerr, ok := err.(common.MyError); ok {
		glog.Errorf("GetShortestDrivingRoute: %s", myerr.ErrorDetails())
	} else {
		glog.Errorf("GetShortestDrivingRoute: type assertion failed for error %#v", err)
		status = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: err.Error()})
	return
}

func mock(w http.ResponseWriter) {
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
