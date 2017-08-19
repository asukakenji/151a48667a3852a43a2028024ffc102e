package lib

import (
	"context"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"

	"googlemaps.github.io/maps"
)

// TODO: Error handling not yet updated
func GetDistanceMatrix(apiKey string, glocs []string) (*maps.DistanceMatrixResponse, common.MyError) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		//TODO: return nil, err
		return nil, nil
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
	resp, err := c.DistanceMatrix(ctx, r)
	if err != nil {
		//TODO: return nil, err
		return nil, nil
	}

	return resp, nil
}
