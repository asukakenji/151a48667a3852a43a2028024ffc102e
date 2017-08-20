package lib

import (
	"fmt"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

// LocationsError is used to indicate an invalid Locations value.
// It extends the common.Error interface.
type LocationsError interface {
	common.Error
	Locations() common.Locations
}

// InsufficientLocationCountError indicates
// there are not enough locations in a Locations value.
// It implements the LocationsError interface.
type InsufficientLocationCountError struct {
	locs common.Locations
	common.DefaultHasher
}

func NewInsufficientLocationCountError(locs common.Locations, hash string) InsufficientLocationCountError {
	return InsufficientLocationCountError{
		locs,
		common.NewDefaultHasher(hash),
	}
}

func (err InsufficientLocationCountError) Locations() common.Locations {
	return err.locs
}

func (err InsufficientLocationCountError) Error() string {
	return fmt.Sprintf(
		"insufficient number of locations: expected >= 2, got %d (%s)",
		len(err.locs),
		err.Hash(),
	)
}

func (err InsufficientLocationCountError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InsufficientLocationCountError: %#v",
		err.Hash(),
		err.locs,
	)
}

// InvalidLocationError indicates an element (a location) in a Locations value
// does not have exactly 2 elements (latitude and longitude).
// It implements the LocationsError interface.
type InvalidLocationError struct {
	locs  common.Locations
	index int
	common.DefaultHasher
}

func NewInvalidLocationError(locs common.Locations, index int, hash string) InvalidLocationError {
	return InvalidLocationError{
		locs,
		index,
		common.NewDefaultHasher(hash),
	}
}

func (err InvalidLocationError) Locations() common.Locations {
	return err.locs
}

func (err InvalidLocationError) Index() int {
	return err.index
}

func (err InvalidLocationError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start location (%s)",
			err.Hash(),
		)
	}
	return fmt.Sprintf(
		"invalid drop off location #%d (%s)",
		err.index,
		err.Hash(),
	)
}

func (err InvalidLocationError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InvalidLocationError: %#v",
		err.Hash(),
		err.locs[err.index],
	)
}

// LatitudeError indicates an invalid latitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
type LatitudeError struct {
	locs  common.Locations
	index int
	common.DefaultHasher
}

func NewLatitudeError(locs common.Locations, index int, hash string) LatitudeError {
	return LatitudeError{
		locs,
		index,
		common.NewDefaultHasher(hash),
	}
}

func (err LatitudeError) Locations() common.Locations {
	return err.locs
}

func (err LatitudeError) Index() int {
	return err.index
}

func (err LatitudeError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start latitude: %q (%s)",
			err.locs[err.index][0],
			err.Hash(),
		)
	}
	return fmt.Sprintf(
		"invalid drop off latitude #%d: %q (%s)",
		err.index,
		err.locs[err.index][0],
		err.Hash(),
	)
}

func (err LatitudeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] LatitudeError: %#v",
		err.Hash(),
		err.locs[err.index][0],
	)
}

// LongitudeError indicates an invalid longitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
type LongitudeError struct {
	locs  common.Locations
	index int
	common.DefaultHasher
}

func NewLongitudeError(locs common.Locations, index int, hash string) LongitudeError {
	return LongitudeError{
		locs,
		index,
		common.NewDefaultHasher(hash),
	}
}

func (err LongitudeError) Locations() common.Locations {
	return err.locs
}

func (err LongitudeError) Index() int {
	return err.index
}

func (err LongitudeError) Error() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start longitude: %q (%s)",
			err.locs[err.index][1],
			err.Hash(),
		)
	}
	return fmt.Sprintf(
		"invalid drop off longitude #%d: %q (%s)",
		err.index,
		err.locs[err.index][1],
		err.Hash(),
	)
}

func (err LongitudeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] LongitudeError: %#v",
		err.Hash(),
		err.locs[err.index][1],
	)
}
