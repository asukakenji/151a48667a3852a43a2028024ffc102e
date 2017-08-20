package common

import "fmt"

// Causer is a super-interface for types having a cause.
type Causer interface {
	Cause() error
}

type DefaultCauser struct {
	cause error
}

func NewDefaultCauser(cause error) DefaultCauser {
	return DefaultCauser{cause}
}

func (causer DefaultCauser) Cause() error {
	return causer.cause
}

// Hasher is a super-interface for types having a hash.
type Hasher interface {
	Hash() string
}

type DefaultHasher struct {
	hash string
}

func NewDefaultHasher(hash string) DefaultHasher {
	return DefaultHasher{hash}
}

func (hasher DefaultHasher) Hash() string {
	return hasher.hash
}

/*
Error is the root type of all errors defined in this project.
It extends the error interfaces.

The Error function required by the error interface
is used for returning an external error message for a client:

	Error() string

The ErrorDetails function
is used for returning an internal error message for logging:

	ErrorDetails() string
*/
type Error interface {
	Hasher
	error
	ErrorDetails() string
}

type CauserError interface {
	Causer
	Error
}

// WrappedError is used to implement the WrapError function.
// It implements the Causer interface.
type WrappedError struct {
	DefaultCauser
	DefaultHasher
}

// WrapError wraps a (generic) error value to a Error value.
// Doing so allows the logic of the API client to be unified to deal with only Error.
//
// The returned error implements the Causer interface.
// The original error could be found by calling Cause.
func WrapError(cause error, hash string) CauserError {
	return WrappedError{
		DefaultCauser{cause},
		DefaultHasher{hash},
	}
}

func (err WrappedError) Error() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.Hash(),
	)
}

func (err WrappedError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] WrappedError: %#v",
		err.Hash(),
		err.Cause(),
	)
}

// JSONEncodeError is used to indicate an error occurred
// when a value is encoded to a JSON message.
// It implements the Causer interface.
type JSONEncodeError struct {
	DefaultCauser
	DefaultHasher
}

func NewJSONEncodeError(err error, hash string) JSONEncodeError {
	return JSONEncodeError{
		DefaultCauser{err},
		DefaultHasher{hash},
	}
}

func (err JSONEncodeError) Error() string {
	return fmt.Sprintf(
		"internal server error (%s)",
		err.Hash(),
	)
}

func (err JSONEncodeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] JSONEncodeError: %#v",
		err.Hash(),
		err.Cause(),
	)
}

// JSONDecodeError is used to indicate an error occurred
// when a value is decoded from a JSON message.
// It implements the Causer interface.
type JSONDecodeError struct {
	DefaultCauser
	DefaultHasher
}

func NewJSONDecodeError(err error, hash string) JSONDecodeError {
	return JSONDecodeError{
		DefaultCauser{err},
		DefaultHasher{hash},
	}
}

func (err JSONDecodeError) Error() string {
	return fmt.Sprintf(
		"invalid JSON (%s)",
		err.Hash(),
	)
}

func (err JSONDecodeError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] JSONDecodeError: %#v",
		err.Hash(),
		err.Cause(),
	)
}

// InvalidTokenError indicates an invalid token.
// It implements the Error interface.
type InvalidTokenError struct {
	token string
	DefaultHasher
}

func NewInvalidTokenError(token, hash string) InvalidTokenError {
	return InvalidTokenError{
		token,
		DefaultHasher{hash},
	}
}

func (err InvalidTokenError) Token() string {
	return err.token
}

func (err InvalidTokenError) Error() string {
	return fmt.Sprintf(
		"invalid token: %q (%s)",
		err.token,
		err.Hash(),
	)
}

func (err InvalidTokenError) ErrorDetails() string {
	return fmt.Sprintf(
		"[%s] InvalidTokenError: %q",
		err.Hash(),
		err.token,
	)
}
