package main

import (
	"encoding/json"
	"net/http"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/cmd/frontier/lib"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
)

// SubmitStartPointAndDropOffLocations deals with the request "POST /route".
func SubmitStartPointAndDropOffLocations(config *Config, w http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	var token string
	var id uint64
	var err common.Error

	// --- Copied from http.Request.Body ---
	// For server requests the Request Body is always non-nil but will return
	// EOF immediately when no body is present. The Server will close the
	// request body. The ServeHTTP Handler does not need to.
	locs, err := lib.LocationsFromJSON(req.Body)
	if err != nil {
		status = http.StatusBadRequest
		goto HandleError
	}

	token = common.NewToken()

	err = taskqueue.WithConnection(config.TaskQueueAddress, func(conn *taskqueue.Connection) common.Error {
		_id, _err := taskqueue.AddQuestion(conn, token, locs, config.TimeLimitForFindingAnswer)
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
	glog.Errorf("[%s] SubmitStartPointAndDropOffLocations: %s", err.Hash(), err.ErrorDetails())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}
