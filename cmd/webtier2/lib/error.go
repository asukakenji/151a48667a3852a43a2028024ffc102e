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

type TrialCountLimitExceededError struct {
	questionID    uint64
	maxTrialCount int
	token         string
}

func NewTrialCountLimitExceededError(token string, questionID uint64, maxTrialCount int) TrialCountLimitExceededError {
	return TrialCountLimitExceededError{questionID, maxTrialCount, token}
}

func (err TrialCountLimitExceededError) QuestionID() uint64 {
	return err.questionID
}

func (err TrialCountLimitExceededError) MaxTrialCount() int {
	return err.maxTrialCount
}

func (err TrialCountLimitExceededError) Token() string {
	return err.token
}

func (err TrialCountLimitExceededError) Error() string {
	return fmt.Sprintf(
		"trial limit exceeded: %d",
		err.maxTrialCount,
	)
}

func (err TrialCountLimitExceededError) ErrorDetails() string {
	return fmt.Sprintf(
		"TrialCountLimitExceededError: Token: %q, Question ID: %d, Max Trial Count: %d",
		err.token, err.questionID, err.maxTrialCount,
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
