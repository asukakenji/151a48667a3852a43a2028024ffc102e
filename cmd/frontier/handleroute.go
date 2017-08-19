package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
)

// SubmitStartPointAndDropOffLocations deals with the request "POST /route".
func SubmitStartPointAndDropOffLocations(w http.ResponseWriter, req *http.Request) {
	addr := "127.0.0.1:11300"    // TODO: Customize: addr
	timeLimit := 5 * time.Second // TODO: Customize: timeLimit
	status := http.StatusOK
	var token string
	var id uint64
	var err error

	// --- Copied from http.Request.Body ---
	// For server requests the Request Body is always non-nil but will return
	// EOF immediately when no body is present. The Server will close the
	// request body. The ServeHTTP Handler does not need to.
	locs, err := common.LocationsFromJSON(req.Body)
	if err != nil {
		status = http.StatusBadRequest
		goto HandleError
	}

	token = common.NewToken()

	err = taskqueue.WithConnection(addr, func(conn *taskqueue.Connection) error {
		_id, _err := taskqueue.AddQuestion(conn, token, locs, timeLimit)
		id = _id
		return _err
	})
	if err != nil {
		status = http.StatusInternalServerError
		goto HandleError
	}

	glog.Infof("SubmitStartPointAndDropOffLocations: id: %d, token: %q, locations: %#v", id, token, locs)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: token})
	return

HandleError:
	if myerr, ok := err.(common.MyError); ok {
		glog.Errorf("SubmitStartPointAndDropOffLocations: %s", myerr.ErrorDetails())
	} else {
		glog.Errorf("SubmitStartPointAndDropOffLocations: type assertion failed for error %#v", err)
		status = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}
