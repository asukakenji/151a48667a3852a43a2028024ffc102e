package main

import (
	"reflect"
	"testing"

	"googlemaps.github.io/maps"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
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

func TestGoogleMapsMatrixToMatrix(t *testing.T) {
	cases := []struct {
		dmr      *maps.DistanceMatrixResponse
		expected matrix.Matrix
	}{
		{nil, matrix.NewSquareMatrix([][]int{})},
	}
	for _, c := range cases {
		got := GoogleMapsMatrixToMatrix(c.dmr)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf(
				"GoogleMapsMatrixToMatrix(%v) = %v, expected %v",
				c.dmr, got, c.expected,
			)
		}
	}
}

func TestCalculateTotalTime(t *testing.T) {
	cases := []struct {
		dmr      *maps.DistanceMatrixResponse
		path     []int
		expected int
	}{
		{nil, nil, 0},
	}
	for _, c := range cases {
		got := CalculateTotalTime(c.dmr, c.path)
		if got != c.expected {
			t.Errorf(
				"CalculateTotalTime(%v, %v) = %d, expected %d",
				c.dmr, c.path, got, c.expected,
			)
		}
	}
}
