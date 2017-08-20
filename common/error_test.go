package common

import (
	"fmt"
	"strings"
	"testing"
)

func TestCauser(t *testing.T) {
	cause := fmt.Errorf("test")
	causer := NewDefaultCauser(cause)
	if causer.Cause() != cause {
		t.Errorf("causer.Cause() not expected")
	}
}

func TestHasher(t *testing.T) {
	hash := "e7e7f7b6-5654-45a0-94b3-4a2215b62739"
	hasher := NewDefaultHasher(hash)
	if hasher.Hash() != hash {
		t.Errorf("hasher.Hash() not expected")
	}
}

func TestWrappedError(t *testing.T) {
	name := "WrappedError"
	hash := "ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b"
	cause := fmt.Errorf("test")
	err := WrapError(cause, hash)
	var _ Error = err
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestJSONEncodeError(t *testing.T) {
	name := "JSONEncodeError"
	hash := "b48b9aae-1abc-444a-9db8-89f3778d11ea"
	cause := fmt.Errorf("test")
	err := NewJSONEncodeError(cause, hash)
	var _ Error = err
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestJSONDecodeError(t *testing.T) {
	name := "JSONDecodeError"
	hash := "3ced8e3b-82e9-42ab-acd5-47e585bb7f11"
	cause := fmt.Errorf("test")
	err := NewJSONDecodeError(cause, hash)
	var _ Error = err
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestInvalidTokenError(t *testing.T) {
	name := "InvalidTokenError"
	hash := "d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4"
	token := "0939d991-f6ee-4beb-997d-2d3b610b1a4f"
	err := NewInvalidTokenError(token, hash)
	var _ Error = err
	if err.Token() != token {
		t.Errorf("err.Token not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}
