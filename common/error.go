package common

import "fmt"

/*
MyError is the root type of all errors defined in this project.
It extends the error interfaces.

The Error function required by the error interface
is used for returning an external error message for a client:

	Error() string

The ErrorDetails function
is used for returning an internal error message for logging:

	ErrorDetails() string
*/
type MyError interface {
	Hash() string
	error
	ErrorDetails() string
}

// WrapError wraps a (generic) error value to a MyError value.
// Doing so allows the logic of the API client to be unified to deal with only MyError.
//
// The returned error implements the Causer interface.
// The original error could be found by calling Cause.
func WrapError(cause error, hash string) Causer {
	return WrappedError{cause, hash}
}

// WrappedError is used to implement the WrapError function.
// It implements the Causer interface.
type WrappedError struct {
	cause error
	hash  string
}

func (err WrappedError) Cause() error {
	return err.cause
}

func (err WrappedError) Hash() string {
	return err.hash
}

func (err WrappedError) Error() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.hash,
	)
}

func (err WrappedError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] WrappedError: %#v",
		err.hash,
		err.cause,
	)
}

// Causer is an interface for errors having a cause.
// It extends the MyError interface.
type Causer interface {
	MyError
	Cause() error
}

// JSONEncodeError is used to indicate an error occurred
// when a value is encoded to a JSON message.
// It implements the Causer interface.
type JSONEncodeError struct {
	cause error
	hash  string
}

func NewJSONEncodeError(err error, hash string) JSONEncodeError {
	return JSONEncodeError{err, hash}
}

func (err JSONEncodeError) Cause() error {
	return err.cause
}

func (err JSONEncodeError) Hash() string {
	return err.hash
}

func (err JSONEncodeError) Error() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.hash,
	)
}

func (err JSONEncodeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] JSONEncodeError: %#v",
		err.hash,
		err.cause,
	)
}

// JSONDecodeError is used to indicate an error occurred
// when a value is decoded from a JSON message.
// It implements the Causer interface.
type JSONDecodeError struct {
	cause error
	hash  string
}

func NewJSONDecodeError(err error, hash string) JSONDecodeError {
	return JSONDecodeError{err, hash}
}

func (err JSONDecodeError) Cause() error {
	return err.cause
}

func (err JSONDecodeError) Hash() string {
	return err.hash
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf(
		"invalid JSON (%s)",
		err.hash,
	)
}

func (err JSONDecodeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] JSONDecodeError: %#v",
		err.hash,
		err.cause,
	)
}

// InvalidTokenError indicates an invalid token.
// It implements the MyError interface.
type InvalidTokenError struct {
	token string
	hash  string
}

func NewInvalidTokenError(token, hash string) InvalidTokenError {
	return InvalidTokenError{token, hash}
}

func (err InvalidTokenError) Token() string {
	return err.token
}

func (err InvalidTokenError) Hash() string {
	return err.hash
}

func (err InvalidTokenError) Error() string {
	return fmt.Sprintf(
		"invalid token: %q (%s)",
		err.token,
		err.hash,
	)
}

func (err InvalidTokenError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InvalidTokenError: %q",
		err.hash,
		err.token,
	)
}
