package lib

import (
	"context"

	"googlemaps.github.io/maps"
)

func GetDistanceMatrix(apiKey string, glocs []string) (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return resp, nil
}
