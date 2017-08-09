package main

import (
	"encoding/json"
	"net/http"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
)

// SubmitStartPointAndDropOffLocations deals with the request "POST /route".
func SubmitStartPointAndDropOffLocations(w http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	var myerr common.MyError

	// --- Copied from http.Request.Body ---
	// For server requests the Request Body is always non-nil but will return
	// EOF immediately when no body is present. The Server will close the
	// request body. The ServeHTTP Handler does not need to.
	locs, err := common.LocationsFromJSON(req.Body)
	if err != nil {
		switch err := err.(type) {
		case common.JSONDecodeError:
			status = http.StatusBadRequest
			myerr = err
		case common.LocationsError:
			status = http.StatusBadRequest
			myerr = err
		default:
			status = http.StatusInternalServerError
			myerr = common.WrapError(err, "a95d9878a129695be71d2f9abdd0d828")
		}
	}
	if status != http.StatusOK {
		glog.Errorf("SubmitStartPointAndDropOffLocations: %s", err.Error())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: myerr.String()})
		return
	}

	glog.Infof("SubmitStartPointAndDropOffLocations: locations: %#v", locs)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO: Implement this!
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{Token: common.NewToken()})
}
