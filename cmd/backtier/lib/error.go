package lib

import (
	"fmt"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"

	"googlemaps.github.io/maps"
)

type TokenCollisionError struct {
	newQuestionID uint64
	oldQuestionID uint64
	token         string
	common.DefaultHasher
}

func NewTokenCollisionError(token string, newQuestionId, oldQuestionId uint64, hash string) TokenCollisionError {
	return TokenCollisionError{
		newQuestionId,
		oldQuestionId,
		token,
		common.NewDefaultHasher(hash),
	}
}

func (err TokenCollisionError) NewQuestionID() uint64 {
	return err.newQuestionID
}

func (err TokenCollisionError) OldQuestionID() uint64 {
	return err.oldQuestionID
}

func (err TokenCollisionError) Token() string {
	return err.token
}

func (err TokenCollisionError) Error() string {
	return fmt.Sprintf(
		"token collision: %q",
		err.token,
	)
}

func (err TokenCollisionError) ErrorDetails() string {
	return fmt.Sprintf(
		"TokenCollisionError: Token: %s, New Question ID: %d, Old Question ID: %d",
		err.token, err.newQuestionID, err.oldQuestionID,
	)
}

type RetryCountLimitExceededError struct {
	questionID    uint64
	maxRetryCount int
	token         string
	common.DefaultHasher
}

func NewRetryCountLimitExceededError(token string, questionID uint64, maxRetryCount int, hash string) RetryCountLimitExceededError {
	return RetryCountLimitExceededError{
		questionID,
		maxRetryCount,
		token,
		common.NewDefaultHasher(hash),
	}
}

func (err RetryCountLimitExceededError) QuestionID() uint64 {
	return err.questionID
}

func (err RetryCountLimitExceededError) MaxRetryCount() int {
	return err.maxRetryCount
}

func (err RetryCountLimitExceededError) Token() string {
	return err.token
}

func (err RetryCountLimitExceededError) Error() string {
	return fmt.Sprintf(
		"retry limit exceeded: %d",
		err.maxRetryCount,
	)
}

func (err RetryCountLimitExceededError) ErrorDetails() string {
	return fmt.Sprintf(
		"RetryCountLimitExceededError: Token: %q, Question ID: %d, Max Retry Count: %d",
		err.token, err.questionID, err.maxRetryCount,
	)
}

type LocationNotFoundError struct {
	resp *maps.DistanceMatrixResponse
	row  int
	col  int
	common.DefaultHasher
}

func NewLocationNotFoundError(resp *maps.DistanceMatrixResponse, row, col int, hash string) LocationNotFoundError {
	return LocationNotFoundError{
		resp,
		row,
		col,
		common.NewDefaultHasher(hash),
	}
}

func (err LocationNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.resp
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
		err.resp.OriginAddresses[err.row],
		err.resp.OriginAddresses[err.col],
	)
}

func (err LocationNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"LocationNotFoundError (%d, %d): %#v",
		err.row, err.col, err.resp,
	)
}

type RouteNotFoundError struct {
	resp *maps.DistanceMatrixResponse
	row  int
	col  int
	common.DefaultHasher
}

func NewRouteNotFoundError(resp *maps.DistanceMatrixResponse, row, col int, hash string) RouteNotFoundError {
	return RouteNotFoundError{
		resp,
		row,
		col,
		common.NewDefaultHasher(hash),
	}
}

func (err RouteNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.resp
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
		err.resp.OriginAddresses[err.row],
		err.resp.OriginAddresses[err.col],
	)
}

func (err RouteNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"RouteNotFoundError (%d, %d): %#v",
		err.row, err.col, err.resp,
	)
}

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
		"map provider connection error (%s)",
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

type ExternalAPIError struct {
	common.DefaultCauser
	common.DefaultHasher
}

func NewExternalAPIError(cause error, hash string) ExternalAPIError {
	return ExternalAPIError{
		common.NewDefaultCauser(cause),
		common.NewDefaultHasher(hash),
	}
}

func (err ExternalAPIError) Error() string {
	return fmt.Sprintf(
		"external API error (%s)",
		err.Hash(),
	)
}

func (err ExternalAPIError) ErrorDetails() string {
	return fmt.Sprintf(
		"ExternalAPIError (%s): %#v",
		err.Hash(),
		err.Cause(),
	)
}
