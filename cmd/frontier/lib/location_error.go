package lib

import (
	"fmt"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

// LocationsError is used to indicate an invalid Locations value.
// It extends the MyError interface.
type LocationsError interface {
	common.MyError
	Locations() common.Locations
}

// InsufficientLocationCountError indicates
// there are not enough locations in a Locations value.
// It implements the LocationsError interface.
type InsufficientLocationCountError struct {
	locs common.Locations
	hash string
}

func NewInsufficientLocationCountError(locs common.Locations, hash string) InsufficientLocationCountError {
	return InsufficientLocationCountError{locs, hash}
}

func (err InsufficientLocationCountError) Locations() common.Locations {
	return err.locs
}

func (err InsufficientLocationCountError) Hash() string {
	return err.hash
}

func (err InsufficientLocationCountError) Error() string {
	return fmt.Sprintf(
		"insufficient number of locations: expected >= 2, got %d (%s)",
		len(err.locs),
		err.hash,
	)
}

func (err InsufficientLocationCountError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InsufficientLocationCountError: %#v",
		err.hash,
		err.locs,
	)
}

// InvalidLocationError indicates an element (a location) in a Locations value
// does not have exactly 2 elements (latitude and longitude).
// It implements the LocationsError interface.
type InvalidLocationError struct {
	locs  common.Locations
	index int
	hash  string
}

func NewInvalidLocationError(locs common.Locations, index int, hash string) InvalidLocationError {
	return InvalidLocationError{locs, index, hash}
}

func (err InvalidLocationError) Locations() common.Locations {
	return err.locs
}

func (err InvalidLocationError) Index() int {
	return err.index
}

func (err InvalidLocationError) Hash() string {
	return err.hash
}

func (err InvalidLocationError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start location (%s)",
			err.hash,
		)
	}
	return fmt.Sprintf(
		"invalid drop off location #%d (%s)",
		err.index,
		err.hash,
	)
}

func (err InvalidLocationError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InvalidLocationError: %#v",
		err.hash,
		err.locs[err.index],
	)
}

// LatitudeError indicates an invalid latitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
type LatitudeError struct {
	locs  common.Locations
	index int
	hash  string
}

func NewLatitudeError(locs common.Locations, index int, hash string) LatitudeError {
	return LatitudeError{locs, index, hash}
}

func (err LatitudeError) Locations() common.Locations {
	return err.locs
}

func (err LatitudeError) Index() int {
	return err.index
}

func (err LatitudeError) Hash() string {
	return err.hash
}

func (err LatitudeError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start latitude: %q (%s)",
			err.locs[err.index][0],
			err.hash,
		)
	}
	return fmt.Sprintf(
		"invalid drop off latitude #%d: %q (%s)",
		err.index,
		err.locs[err.index][0],
		err.hash,
	)
}

func (err LatitudeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] LatitudeError: %#v",
		err.hash,
		err.locs[err.index][0],
	)
}

// LongitudeError indicates an invalid longitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
type LongitudeError struct {
	locs  common.Locations
	index int
	hash  string
}

func NewLongitudeError(locs common.Locations, index int, hash string) LongitudeError {
	return LongitudeError{locs, index, hash}
}

func (err LongitudeError) Locations() common.Locations {
	return err.locs
}

func (err LongitudeError) Index() int {
	return err.index
}

func (err LongitudeError) Hash() string {
	return err.hash
}

func (err LongitudeError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start longitude: %q (%s)",
			err.locs[err.index][1],
			err.hash,
		)
	}
	return fmt.Sprintf(
		"invalid drop off longitude #%d: %q (%s)",
		err.index,
		err.locs[err.index][1],
		err.hash,
	)
}

func (err LongitudeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] LongitudeError: %#v",
		err.hash,
		err.locs[err.index][1],
	)
}
