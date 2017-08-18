package main

import (
	"fmt"

	"googlemaps.github.io/maps"
)

type LocationNotFoundError struct {
	dmr *maps.DistanceMatrixResponse
	row int
	col int
}

func (err LocationNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.dmr
}

func (err LocationNotFoundError) Row() int {
	return err.row
}

func (err LocationNotFoundError) Col() int {
	return err.col
}

func (err LocationNotFoundError) Error() string {
	return fmt.Sprintf(
		"location not found: %q or %q",
		err.dmr.OriginAddresses[err.row],
		err.dmr.OriginAddresses[err.col],
	)
}

func (err LocationNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"LocationNotFoundError (%d, %d): %#v",
		err.row, err.col, err.dmr,
	)
}

type RouteNotFoundError struct {
	dmr *maps.DistanceMatrixResponse
	row int
	col int
}

func (err RouteNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.dmr
}

func (err RouteNotFoundError) Row() int {
	return err.row
}

func (err RouteNotFoundError) Col() int {
	return err.col
}

func (err RouteNotFoundError) Error() string {
	return fmt.Sprintf(
		"route not found: from %q to %q",
		err.dmr.OriginAddresses[err.row],
		err.dmr.OriginAddresses[err.col],
	)
}

func (err RouteNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"RouteNotFoundError (%d, %d): %#v",
		err.row, err.col, err.dmr,
	)
}
