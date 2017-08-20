package main

import (
	"encoding/json"
	"net/http"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// GetShortestDrivingRoute deals with the request "GET /route/{token}".
func GetShortestDrivingRoute(w http.ResponseWriter, req *http.Request) {
	addr := "127.0.0.1:11300" // TODO: Customize: addr
	status := http.StatusOK
	var id uint64
	var dr *common.DrivingRoute
	var err common.Error

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
		hash := common.NewToken()
		err = common.NewInvalidTokenError(token, hash)
		goto HandleError
	}

	err = taskqueue.WithConnection(addr, func(conn *taskqueue.Connection) common.Error {
		_id, _dr, _err := taskqueue.GetAnswer1(conn, token)
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
	glog.Errorf("[%s] GetShortestDrivingRoute: %s", err.Hash(), err.ErrorDetails())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: err.Error()})
	return
}
