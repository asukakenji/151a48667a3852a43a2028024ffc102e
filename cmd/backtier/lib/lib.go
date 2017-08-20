package lib

import (
	"context"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"

	"googlemaps.github.io/maps"
)

func GetDistanceMatrix(apiKey string, glocs []string) (resp *maps.DistanceMatrixResponse, err common.Error) {
	c, _err := maps.NewClient(maps.WithAPIKey(apiKey))
	if _err != nil {
		hash := common.NewToken()
		glog.Errorf("[%s] GetDistanceMatrix: NewClient()", hash)
		return nil, NewConnectionError(_err, hash)
	}

	ctx := context.Background()
	r := &maps.DistanceMatrixRequest{
		Origins:                  glocs,
		Destinations:             glocs,
		Mode:                     maps.TravelModeDriving,
		Language:                 "",
		Avoid:                    maps.Avoid(""),
		Units:                    maps.UnitsMetric,
		DepartureTime:            "now",
		ArrivalTime:              "",
		TrafficModel:             maps.TrafficModel(""),
		TransitMode:              []maps.TransitMode(nil),
		TransitRoutingPreference: maps.TransitRoutingPreference(""),
	}
	resp, _err = c.DistanceMatrix(ctx, r)
	if _err != nil {
		hash := common.NewToken()
		glog.Errorf("[%s] GetDistanceMatrix: DistanceMatrix()", hash)
		return nil, NewExternalAPIError(_err, hash)
	}

	return resp, nil
}
