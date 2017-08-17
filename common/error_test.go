package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWrapError(t *testing.T) {
	err0 := fmt.Errorf("test")
	err := WrapError(err0, "d59657d4f804ac7492f5a98fbf2d3bbf")
	if err.Cause() != err0 {
		t.Errorf("err != err0")
	}
	if err.Error() != "internal server error (d59657d4f804ac7492f5a98fbf2d3bbf)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONEncodeError(t *testing.T) {
	err0 := fmt.Errorf("test")
	err := NewJSONEncodeError(err0)
	if err.Cause() != err0 {
		t.Errorf("err != err0")
	}
	if err.Error() != "internal server error (JSONEncodeError)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONDecodeError(t *testing.T) {
	err0 := fmt.Errorf("test")
	err := NewJSONDecodeError(err0)
	if err.Cause() != err0 {
		t.Errorf("err != err0")
	}
	if err.Error() != "invalid JSON" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestLocationsError(t *testing.T) {
	locs := Locations{}
	err := InsufficientLocationCountError{locs}
	if !reflect.DeepEqual(err.Locations(), locs) {
		t.Errorf("err.Locations() != locs")
	}
	if err.Error() != "insufficient number of locations: expected >= 2, got 0" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestInvalidLocationError(t *testing.T) {
	locs0 := Locations{
		{"22.284419", "114.159510", "22.326442"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := InvalidLocationError{locs0, index0}
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() != locs0")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Error() != "invalid route start location" {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "114.159510", "22.326442"},
	}
	index1 := 1
	err1 := InvalidLocationError{locs1, index1}
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() != locs1")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Error() != "invalid drop off location #1" {
		t.Errorf("err1.Error() not expected")
	}
}

func TestLatitudeError(t *testing.T) {
	locs0 := Locations{
		{"90.372081", "114.107877"},
		{"22.284419", "114.159510"},
	}
	index0 := 0
	err0 := LatitudeError{locs0, index0}
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() != locs0")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Error() != "invalid route start latitude: \"90.372081\"" {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.284419", "114.159510"},
		{"90.372081", "114.107877"},
	}
	index1 := 1
	err1 := LatitudeError{locs1, index1}
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() != locs1")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Error() != "invalid drop off latitude #1: \"90.372081\"" {
		t.Errorf("err1.Error() not expected")
	}
}

func TestLongitudeError(t *testing.T) {
	locs0 := Locations{
		{"22.284419", "180.159510"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := LongitudeError{locs0, index0}
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() != locs0")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Error() != "invalid route start longitude: \"180.159510\"" {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "180.159510"},
	}
	index1 := 1
	err1 := LongitudeError{locs1, index1}
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() != locs1")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Error() != "invalid drop off longitude #1: \"180.159510\"" {
		t.Errorf("err1.Error() not expected")
	}
}

func TestInvalidTokenError(t *testing.T) {
	token := "test"
	err := NewInvalidTokenError(token)
	if err.Token() != token {
		t.Errorf("err.Token != token")
	}
	if err.Error() != "invalid token: \"test\"" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}
