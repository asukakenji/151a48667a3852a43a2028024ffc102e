package taskqueue

import (
	"fmt"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

type ConnectionError struct {
	common.DefaultCauser
	common.DefaultHasher
}

func NewConnectionError(cause error, hash string) ConnectionError {
	return ConnectionError{
		common.NewDefaultCauser(cause),
		common.NewDefaultHasher(hash),
	}
}

func (err ConnectionError) Error() string {
	return fmt.Sprintf(
		"task queue connection error (%s)",
		err.Hash(),
	)
}

func (err ConnectionError) ErrorDetails() string {
	return fmt.Sprintf(
		"ConnectionError (%s): %#v",
		err.Hash(),
		err.Cause(),
	)
}

type NotFoundError struct {
	common.DefaultCauser
	common.DefaultHasher
}

func NewNotFoundError(cause error, hash string) NotFoundError {
	return NotFoundError{
		common.NewDefaultCauser(cause),
		common.NewDefaultHasher(hash),
	}
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf(
		"not found error (%s)",
		err.Hash(),
	)
}

func (err NotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"NotFoundError (%s): %#v",
		err.Hash(),
		err.Cause(),
	)
}

type JobTooBigError struct {
	common.DefaultCauser
	common.DefaultHasher
}

func NewJobTooBigError(cause error, hash string) JobTooBigError {
	return JobTooBigError{
		common.NewDefaultCauser(cause),
		common.NewDefaultHasher(hash),
	}
}

func (err JobTooBigError) Error() string {
	return fmt.Sprintf(
		"job too big error (%s)",
		err.Hash(),
	)
}

func (err JobTooBigError) ErrorDetails() string {
	return fmt.Sprintf(
		"JobTooBigError (%s): %#v",
		err.Hash(),
		err.Cause(),
	)
}

type UnexpectedError struct {
	common.DefaultCauser
	common.DefaultHasher
}

func NewUnexpectedError(cause error, hash string) UnexpectedError {
	return UnexpectedError{
		common.NewDefaultCauser(cause),
		common.NewDefaultHasher(hash),
	}
}

func (err UnexpectedError) Error() string {
	return fmt.Sprintf(
		"unexpected error (%s)",
		err.Hash(),
	)
}

func (err UnexpectedError) ErrorDetails() string {
	return fmt.Sprintf(
		"UnexpectedError (%s): %#v",
		err.Hash(),
		err.Cause(),
	)
}
