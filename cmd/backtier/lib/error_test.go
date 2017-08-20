package lib

import (
	"fmt"
	"strings"
	"testing"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

func TestTokenCollisionError(t *testing.T) {
	name := "TokenCollisionError"
	hash := "4ccb3fc4-db1e-46b2-b8c7-d11bebc8816a"
	err := NewTokenCollisionError("d781ecd3-5bd8-48f7-a848-94d25804fead", 123, 42, hash)
	var _ common.Error = err
	if err.NewQuestionID() != 123 {
		t.Errorf("err.NewQuestionID() not expected")
	}
	if err.OldQuestionID() != 42 {
		t.Errorf("err.OldQuestionID() not expected")
	}
	if err.Token() != "d781ecd3-5bd8-48f7-a848-94d25804fead" {
		t.Errorf("err.Token() not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestRetryCountLimitExceededError(t *testing.T) {
	name := "RetryCountLimitExceededError"
	hash := "276c8014-6151-4e43-a18a-105073e77005"
	err := NewRetryCountLimitExceededError("8a14c7cf-c838-4586-a5a8-35c61dedbf9b", 123, 3, hash)
	var _ common.Error = err
	if err.QuestionID() != 123 {
		t.Errorf("err.QuestionID() not expected")
	}
	if err.MaxRetryCount() != 3 {
		t.Errorf("err.MaxRetryCount() not expected")
	}
	if err.Token() != "8a14c7cf-c838-4586-a5a8-35c61dedbf9b" {
		t.Errorf("err.Token() not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestLocationNotFoundError(t *testing.T) {
	name := "LocationNotFoundError"
	hash := "b2cf8674-4d78-46fd-951d-38fa8cb221b9"
	err := NewLocationNotFoundError(resp1, 1, 2, hash)
	var _ common.Error = err
	if err.DistanceMatrixResponse() != resp1 {
		t.Errorf("err.DistanceMatrixResponse() not expected")
	}
	if err.Row() != 1 {
		t.Errorf("err.Row() not expected")
	}
	if err.Col() != 2 {
		t.Errorf("err.Col() not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestRouteNotFoundError(t *testing.T) {
	name := "RouteNotFoundError"
	hash := "5b023d0e-08f2-466f-8ce6-9243048d28e0"
	err := NewRouteNotFoundError(resp2, 0, 1, hash)
	var _ common.Error = err
	if err.DistanceMatrixResponse() != resp2 {
		t.Errorf("err.DistanceMatrixResponse() not expected")
	}
	if err.Row() != 0 {
		t.Errorf("err.Row() not expected")
	}
	if err.Col() != 1 {
		t.Errorf("err.Col() not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestConnectionError(t *testing.T) {
	name := "ConnectionError"
	hash := "73e192e2-ccc3-4fa3-a399-2cfc637ee0a8"
	cause := fmt.Errorf("test")
	err := NewConnectionError(cause, hash)
	var _ common.Error = err
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestExternalAPIError(t *testing.T) {
	name := "ExternalAPIError"
	hash := "aa737aad-8ba4-4855-8644-61a27e3cf6f8"
	cause := fmt.Errorf("test")
	err := NewExternalAPIError(cause, hash)
	var _ common.Error = err
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}
