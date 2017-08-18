package lib

import "testing"

func TestLocationNotFoundError(t *testing.T) {
	err := LocationNotFoundError{dmr1, 1, 2}
	if err.DistanceMatrixResponse() != dmr1 {
		t.Errorf("err.DistanceMatrixResponse() not expected")
	}
	if err.Row() != 1 {
		t.Errorf("err.Row() not expected")
	}
	if err.Col() != 2 {
		t.Errorf("err.Col() not expected")
	}
	if err.Error() != `location not found: "Laguna City, Central, Hong Kong" or "789 Nathan Rd, Mong Kok, Hong Kong"` {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestRouteNotFoundError(t *testing.T) {
	err := RouteNotFoundError{dmr2, 0, 1}
	if err.DistanceMatrixResponse() != dmr2 {
		t.Errorf("err.DistanceMatrixResponse() not expected")
	}
	if err.Row() != 0 {
		t.Errorf("err.Row() not expected")
	}
	if err.Col() != 1 {
		t.Errorf("err.Col() not expected")
	}
	if err.Error() != `route not found: from "142 Prince Edward Rd W, Mong Kok, Hong Kong" to "22.147034,113.559835"` {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}
