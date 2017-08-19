package lib

import (
	"fmt"

	"googlemaps.github.io/maps"
)

type TokenCollisionError struct {
	newQuestionID uint64
	oldQuestionID uint64
	token         string
}

func NewTokenCollisionError(token string, newQuestionId, oldQuestionId uint64) TokenCollisionError {
	return TokenCollisionError{newQuestionId, oldQuestionId, token}
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
}

func NewRetryCountLimitExceededError(token string, questionID uint64, maxRetryCount int) RetryCountLimitExceededError {
	return RetryCountLimitExceededError{questionID, maxRetryCount, token}
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
	dmr *maps.DistanceMatrixResponse
	row int
	col int
}

func (err LocationNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.dmr
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
		err.dmr.OriginAddresses[err.row],
		err.dmr.OriginAddresses[err.col],
	)
}

func (err LocationNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"LocationNotFoundError (%d, %d): %#v",
		err.row, err.col, err.dmr,
	)
}

type RouteNotFoundError struct {
	dmr *maps.DistanceMatrixResponse
	row int
	col int
}

func (err RouteNotFoundError) DistanceMatrixResponse() *maps.DistanceMatrixResponse {
	return err.dmr
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
		err.dmr.OriginAddresses[err.row],
		err.dmr.OriginAddresses[err.col],
	)
}

func (err RouteNotFoundError) ErrorDetails() string {
	return fmt.Sprintf(
		"RouteNotFoundError (%d, %d): %#v",
		err.row, err.col, err.dmr,
	)
}
