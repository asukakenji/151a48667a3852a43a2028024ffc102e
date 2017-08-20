package lib

import (
	"reflect"
	"testing"
	"time"

	"googlemaps.github.io/maps"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/matrix"
)

func TestLocationsToGoogleMapsLocations(t *testing.T) {
	cases := []struct {
		locs     common.Locations
		expected []string
	}{
		{
			nil,
			[]string{},
		},
		{
			common.Locations{},
			[]string{},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
			},
			[]string{
				"22.372081,114.107877",
			},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
			},
			[]string{
				"22.372081,114.107877",
				"22.284419,114.159510",
			},
		},
		{
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
				{"22.326442", "114.167811"},
			},
			[]string{
				"22.372081,114.107877",
				"22.284419,114.159510",
				"22.326442,114.167811",
			},
		},
	}
	for _, c := range cases {
		got := LocationsToGoogleMapsLocations(c.locs)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf(
				"LocationsToGoogleMapsLocations(%v) = %v, expected %v",
				c.locs, got, c.expected,
			)
		}
	}
}

func TestPathToLocationPath(t *testing.T) {
	cases := []struct {
		locs     common.Locations
		path     []int
		expected common.Locations
	}{
		{
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
				{"22.326442", "114.167811"},
			},
			[]int{0, 2, 1},
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.326442", "114.167811"},
				{"22.284419", "114.159510"},
			},
		},
	}
	for _, c := range cases {
		got := PathToLocationPath(c.locs, c.path)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf(
				"PathToLocationPath(%v, %v) = %v, expected %v",
				c.locs, c.path, got, c.expected,
			)
		}
	}
}

func TestGoogleMapsMatrixToMatrix(t *testing.T) {
	cases := []struct {
		resp            *maps.DistanceMatrixResponse
		expectedM       matrix.Matrix
		expectedErrType error
	}{
		{
			nil,
			matrix.NewSquareMatrix([][]int{}),
			nil,
		},
		{
			resp1,
			matrix.NewSquareMatrix([][]int{
				{constant.Infinity, 15518, 9667},
				{15223, constant.Infinity, 8333},
				{10329, 8464, constant.Infinity},
			}),
			nil,
		},
		{
			resp2,
			nil,
			RouteNotFoundError{},
		},
		{
			resp3,
			nil,
			LocationNotFoundError{},
		},
		{
			respX,
			nil,
			common.WrappedError{},
		},
	}
	for _, c := range cases {
		gotM, gotErr := GoogleMapsMatrixToMatrix(c.resp)
		if !reflect.DeepEqual(gotM, c.expectedM) || reflect.TypeOf(gotErr) != reflect.TypeOf(c.expectedErrType) {
			t.Errorf(
				"GoogleMapsMatrixToMatrix(%v) = (%v, %T), expected (%v, %T)",
				c.resp, gotM, gotErr, c.expectedM, c.expectedErrType,
			)
		}
	}
}

func TestCalculateTotalTime(t *testing.T) {
	cases := []struct {
		resp     *maps.DistanceMatrixResponse
		path     []int
		expected time.Duration
	}{
		{nil, nil, 0},
		{resp1, []int{0, 2, 1}, 1753000000000},
	}
	for _, c := range cases {
		got := CalculateTotalTime(c.resp, c.path)
		if got != c.expected {
			t.Errorf(
				"CalculateTotalTime(%v, %v) = %d, expected %d",
				c.resp, c.path, got, c.expected,
			)
		}
	}
}
