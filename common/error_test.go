package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWrappedError(t *testing.T) {
	var _ MyError = WrappedError{}
	var _ Causer = WrappedError{}
	err0 := fmt.Errorf("test")
	err := WrapError(err0, "ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "internal server error (ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONEncodeError(t *testing.T) {
	var _ MyError = JSONEncodeError{}
	var _ Causer = JSONEncodeError{}
	err0 := fmt.Errorf("test")
	err := NewJSONEncodeError(err0, "b48b9aae-1abc-444a-9db8-89f3778d11ea")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "b48b9aae-1abc-444a-9db8-89f3778d11ea" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "internal server error (b48b9aae-1abc-444a-9db8-89f3778d11ea)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONDecodeError(t *testing.T) {
	var _ MyError = JSONDecodeError{}
	var _ Causer = JSONDecodeError{}
	err0 := fmt.Errorf("test")
	err := NewJSONDecodeError(err0, "3ced8e3b-82e9-42ab-acd5-47e585bb7f11")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "3ced8e3b-82e9-42ab-acd5-47e585bb7f11" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "invalid JSON (3ced8e3b-82e9-42ab-acd5-47e585bb7f11)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestInsufficientLocationCountError(t *testing.T) {
	var _ MyError = InsufficientLocationCountError{}
	locs := Locations{}
	err := NewInsufficientLocationCountError(locs, "9240926c-d38d-4683-aeb1-c0c76cc85479")
	if !reflect.DeepEqual(err.Locations(), locs) {
		t.Errorf("err.Locations() not expected")
	}
	if err.Hash() != "9240926c-d38d-4683-aeb1-c0c76cc85479" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "insufficient number of locations: expected >= 2, got 0 (9240926c-d38d-4683-aeb1-c0c76cc85479)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestInvalidLocationError(t *testing.T) {
	var _ MyError = InvalidLocationError{}
	locs0 := Locations{
		{"22.284419", "114.159510", "22.326442"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := NewInvalidLocationError(locs0, index0, "b10efaef-98f0-4ed6-9067-fdc443e8fafa")
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Hash() != "b10efaef-98f0-4ed6-9067-fdc443e8fafa" {
		t.Errorf("err.Hash() not expected")
	}
	if err0.Error() != "invalid route start location (b10efaef-98f0-4ed6-9067-fdc443e8fafa)" {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "114.159510", "22.326442"},
	}
	index1 := 1
	err1 := NewInvalidLocationError(locs1, index1, "7ce896d9-3c3e-48d2-ac4f-2d2950f0ab09")
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Hash() != "7ce896d9-3c3e-48d2-ac4f-2d2950f0ab09" {
		t.Errorf("err.Hash() not expected")
	}
	if err1.Error() != "invalid drop off location #1 (7ce896d9-3c3e-48d2-ac4f-2d2950f0ab09)" {
		t.Errorf("err1.Error() not expected")
	}
	err1.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestLatitudeError(t *testing.T) {
	var _ MyError = LatitudeError{}
	locs0 := Locations{
		{"90.372081", "114.107877"},
		{"22.284419", "114.159510"},
	}
	index0 := 0
	err0 := NewLatitudeError(locs0, index0, "acc75ebd-c145-40fa-9cc6-1f1337d4b8ae")
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Hash() != "acc75ebd-c145-40fa-9cc6-1f1337d4b8ae" {
		t.Errorf("err.Hash() not expected")
	}
	if err0.Error() != `invalid route start latitude: "90.372081" (acc75ebd-c145-40fa-9cc6-1f1337d4b8ae)` {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.284419", "114.159510"},
		{"90.372081", "114.107877"},
	}
	index1 := 1
	err1 := NewLatitudeError(locs1, index1, "2311e1c7-ceb9-43e6-9d02-029e8a3a6541")
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Hash() != "2311e1c7-ceb9-43e6-9d02-029e8a3a6541" {
		t.Errorf("err.Hash() not expected")
	}
	if err1.Error() != `invalid drop off latitude #1: "90.372081" (2311e1c7-ceb9-43e6-9d02-029e8a3a6541)` {
		t.Errorf("err1.Error() not expected")
	}
	err1.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestLongitudeError(t *testing.T) {
	var _ MyError = LongitudeError{}
	locs0 := Locations{
		{"22.284419", "180.159510"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := NewLongitudeError(locs0, index0, "c5727bc8-44e8-4e4e-b378-35fbe2604afc")
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if err0.Hash() != "c5727bc8-44e8-4e4e-b378-35fbe2604afc" {
		t.Errorf("err.Hash() not expected")
	}
	if err0.Error() != `invalid route start longitude: "180.159510" (c5727bc8-44e8-4e4e-b378-35fbe2604afc)` {
		t.Errorf("err0.Error() not expected")
	}
	err0.ErrorDetails() // NOTE: Just call it to mark it tested

	locs1 := Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "180.159510"},
	}
	index1 := 1
	err1 := NewLongitudeError(locs1, index1, "2a965bd4-ea7f-4351-9cb5-8c38f42178a0")
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if err1.Hash() != "2a965bd4-ea7f-4351-9cb5-8c38f42178a0" {
		t.Errorf("err.Hash() not expected")
	}
	if err1.Error() != `invalid drop off longitude #1: "180.159510" (2a965bd4-ea7f-4351-9cb5-8c38f42178a0)` {
		t.Errorf("err1.Error() not expected")
	}
	err1.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestInvalidTokenError(t *testing.T) {
	var _ MyError = InvalidTokenError{}
	token := "0939d991-f6ee-4beb-997d-2d3b610b1a4f"
	err := NewInvalidTokenError(token, "d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4")
	if err.Token() != token {
		t.Errorf("err.Token not expected")
	}
	if err.Hash() != "d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != `invalid token: "0939d991-f6ee-4beb-997d-2d3b610b1a4f" (d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4)` {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}
