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
	cause error
	hash  string
}

func NewNotFoundError(cause error, hash string) NotFoundError {
	return NotFoundError{cause, hash}
}

func (err NotFoundError) Cause() error {
	return err.cause
}

func (err NotFoundError) Hash() string {
	return err.hash
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf(
		"not found error (%s)",
		err.hash,
	)
}

func (err NotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"NotFoundError (%s): %#v",
		err.hash,
		err.cause,
	)
}

type UnexpectedError struct {
	cause error
	hash  string
}

func NewUnexpectedError(cause error, hash string) UnexpectedError {
	return UnexpectedError{cause, hash}
}

func (err UnexpectedError) Cause() error {
	return err.cause
}

func (err UnexpectedError) Hash() string {
	return err.hash
}

func (err UnexpectedError) Error() string {
	return fmt.Sprintf(
		"unexpected error (%s)",
		err.hash,
	)
}

func (err UnexpectedError) ErrorDetails() string {
	return fmt.Sprintf(
		"UnexpectedError (%s): %#v",
		err.hash,
		err.cause,
	)
}
