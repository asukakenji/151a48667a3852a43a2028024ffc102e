package common

import "fmt"

type MyError interface {
	error
	fmt.Stringer
}

func WrapError(cause error, hash string) MyError {
	return UnknownError{cause, hash}
}

type Causer interface {
	MyError
	Cause() error
}

type UnknownError struct {
	cause error
	hash  string
}

func (err UnknownError) Cause() error {
	return err.cause
}

func (err UnknownError) Hash() string {
	return err.hash
}

func (err UnknownError) String() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.hash,
	)
}

func (err UnknownError) Error() string {
	return fmt.Sprintf(
		"UnknownError (%s): %#v",
		err.hash,
		err.cause,
	)
}

type JSONDecodeError struct {
	cause error
}

func (err JSONDecodeError) Cause() error {
	return err.cause
}

func (err JSONDecodeError) String() string {
	return "invalid JSON"
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf(
		"JSONDecodeError: %#v",
		err.cause,
	)
}

type InvalidLocationsError interface {
	MyError
	Locations() Locations
}

type InsufficientLocationCountError struct {
	locs Locations
}

func (err InsufficientLocationCountError) Locations() Locations {
	return err.locs
}

func (err InsufficientLocationCountError) String() string {
	return fmt.Sprintf(
		"insufficient number of locations: expected >= 2, got %d",
		len(err.locs),
	)
}

func (err InsufficientLocationCountError) Error() string {
	return fmt.Sprintf(
		"InsufficientLocationCountError: %#v",
		err.locs,
	)
}

type InvalidLocationError struct {
	locs  Locations
	index int
}

func (err InvalidLocationError) Locations() Locations {
	return err.locs
}

func (err InvalidLocationError) Index() int {
	return err.index
}

func (err InvalidLocationError) String() string {
	if err.index == 0 {
		return "invalid route start location"
	}
	return fmt.Sprintf(
		"invalid drop off location #%d",
		err.index,
	)
}

func (err InvalidLocationError) Error() string {
	return fmt.Sprintf(
		"InvalidLocationError: %#v",
		err.locs[err.index],
	)
}

type LatitudeError struct {
	locs  Locations
	index int
}

func (err LatitudeError) Locations() Locations {
	return err.locs
}

func (err LatitudeError) Index() int {
	return err.index
}

func (err LatitudeError) String() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start latitude: %q",
			err.locs[err.index][0],
		)
	}
	return fmt.Sprintf(
		"invalid drop off latitude #%d: %q",
		err.index,
		err.locs[err.index][0],
	)
}

func (err LatitudeError) Error() string {
	return fmt.Sprintf(
		"LatitudeError: %#v",
		err.locs[err.index][0],
	)
}

type LongitudeError struct {
	locs  Locations
	index int
}

func (err LongitudeError) Locations() Locations {
	return err.locs
}

func (err LongitudeError) Index() int {
	return err.index
}

func (err LongitudeError) String() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start latitude: %q",
			err.locs[err.index][1],
		)
	}
	return fmt.Sprintf(
		"invalid drop off latitude #%d: %q",
		err.index,
		err.locs[err.index][1],
	)
}

func (err LongitudeError) Error() string {
	return fmt.Sprintf(
		"LongitudeError: %#v",
		err.locs[err.index][1],
	)
}

type InvalidTokenError struct {
	token string
}

func NewInvalidTokenError(token string) InvalidTokenError {
	return InvalidTokenError{token}
}

func (err InvalidTokenError) Token() string {
	return err.token
}

func (err InvalidTokenError) String() string {
	return fmt.Sprintf(
		"invalid token: %q",
		err.token,
	)
}

func (err InvalidTokenError) Error() string {
	return fmt.Sprintf(
		"TokenError: %#v",
		err.token,
	)
}
