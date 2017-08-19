package taskqueue

import "fmt"

type ConnectionError struct {
	cause error
	hash  string
}

func NewConnectionError(cause error, hash string) ConnectionError {
	return ConnectionError{cause, hash}
}

func (err ConnectionError) Cause() error {
	return err.cause
}

func (err ConnectionError) Hash() string {
	return err.hash
}

func (err ConnectionError) Error() string {
	return fmt.Sprintf(
		"task queue connection error (%s)",
		err.hash,
	)
}

func (err ConnectionError) ErrorDetails() string {
	return fmt.Sprintf(
		"ConnectionError (%s): %#v",
		err.hash,
		err.cause,
	)
}

type NotFoundError struct {
}

func NewNotFoundError() NotFoundError {
	return NotFoundError{}
}

func (err NotFoundError) Error() string {
	return ""
}

func (err NotFoundError) ErrorDetails() string {
	return ""
}
