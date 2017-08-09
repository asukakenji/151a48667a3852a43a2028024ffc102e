package common

import "fmt"

/*
MyError is the root type of all errors defined in this project.
It extends 2 interfaces: error and fmt.Stringer.

The Error function required by the error interface
is used for returning an internal error message for logging:

	Error() string

The String function required by the fmt.Stringer interface
is used for returning an external error message for a client:

	String() string
*/
type MyError interface {
	error
	fmt.Stringer
}

// WrapError wraps a (generic) error value to a MyError value.
// Doing so allows the logic of the API client to be unified to deal with only MyError.
//
// The returned error implements the Causer interface.
// The original error could be found by calling Cause.
func WrapError(cause error, hash string) Causer {
	return unknownError{cause, hash}
}

// unknownError is used to implement the WrapError function.
// It implements the Causer interface.
type unknownError struct {
	cause error
	hash  string
}

func (err unknownError) Cause() error {
	return err.cause
}

func (err unknownError) Error() string {
	return fmt.Sprintf(
		"UnknownError (%s): %#v",
		err.hash,
		err.cause,
	)
}

func (err unknownError) String() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.hash,
	)
}

// Causer is an interface for errors having a cause.
// It extends the MyError interface.
type Causer interface {
	MyError
	Cause() error
}

// JSONDecodeError is used to indicate an error occurred
// when a value is encoded to a JSON message.
// It implements the Causer interface.
type JSONEncodeError struct {
	cause error
}

func (err JSONEncodeError) Cause() error {
	return err.cause
}

func (err JSONEncodeError) Error() string {
	return fmt.Sprintf(
		"JSONEncodeError: %#v",
		err.cause,
	)
}

func (err JSONEncodeError) String() string {
	return "internal server error (JSONEncodeError)"
}

// JSONDecodeError is used to indicate an error occurred
// when a value is decoded from a JSON message.
// It implements the Causer interface.
type JSONDecodeError struct {
	cause error
}

func (err JSONDecodeError) Cause() error {
	return err.cause
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf(
		"JSONDecodeError: %#v",
		err.cause,
	)
}

func (err JSONDecodeError) String() string {
	return "invalid JSON"
}

// LocationsError is used to indicate an invalid Locations value.
// It extends the MyError interface.
type LocationsError interface {
	MyError
	Locations() Locations
}

// InsufficientLocationCountError indicates
// there are not enough locations in a Locations value.
// It implements the LocationsError interface.
type InsufficientLocationCountError struct {
	locs Locations
}

func (err InsufficientLocationCountError) Locations() Locations {
	return err.locs
}

func (err InsufficientLocationCountError) Error() string {
	return fmt.Sprintf(
		"InsufficientLocationCountError: %#v",
		err.locs,
	)
}

func (err InsufficientLocationCountError) String() string {
	return fmt.Sprintf(
		"insufficient number of locations: expected >= 2, got %d",
		len(err.locs),
	)
}

// InvalidLocationError indicates an element (a location) in a Locations value
// does not have exactly 2 elements (latitude and longitude).
// It implements the LocationsError interface.
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

func (err InvalidLocationError) Error() string {
	return fmt.Sprintf(
		"InvalidLocationError: %#v",
		err.locs[err.index],
	)
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

// LatitudeError indicates an invalid latitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
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

func (err LatitudeError) Error() string {
	return fmt.Sprintf(
		"LatitudeError: %#v",
		err.locs[err.index][0],
	)
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

// LongitudeError indicates an invalid longitude of
// an element (a location) in a Locations value.
// It implements the LocationsError interface.
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

func (err LongitudeError) Error() string {
	return fmt.Sprintf(
		"LongitudeError: %#v",
		err.locs[err.index][1],
	)
}

func (err LongitudeError) String() string {
	if err.index == 0 {
		return fmt.Sprintf(
			"invalid route start longitude: %q",
			err.locs[err.index][1],
		)
	}
	return fmt.Sprintf(
		"invalid drop off longitude #%d: %q",
		err.index,
		err.locs[err.index][1],
	)
}

// InvalidTokenError indicates an invalid token.
// It implements the MyError interface.
type InvalidTokenError struct {
	token string
}

func NewInvalidTokenError(token string) InvalidTokenError {
	return InvalidTokenError{token}
}

func (err InvalidTokenError) Token() string {
	return err.token
}

func (err InvalidTokenError) Error() string {
	return fmt.Sprintf(
		"InvalidTokenError: %#v",
		err.token,
	)
}

func (err InvalidTokenError) String() string {
	return fmt.Sprintf(
		"invalid token: %q",
		err.token,
	)
}
