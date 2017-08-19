package common

import (
	"fmt"
	"testing"
)

func TestWrappedError(t *testing.T) {
	var _ MyError = WrappedError{}
	var _ Causer = WrappedError{}
	err0 := fmt.Errorf("test")
	err := WrapError(err0, "ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "internal server error (ad7ae720-b6e2-4eb9-8e9b-c9dbd009758b)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONEncodeError(t *testing.T) {
	var _ MyError = JSONEncodeError{}
	var _ Causer = JSONEncodeError{}
	err0 := fmt.Errorf("test")
	err := NewJSONEncodeError(err0, "b48b9aae-1abc-444a-9db8-89f3778d11ea")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "b48b9aae-1abc-444a-9db8-89f3778d11ea" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "internal server error (b48b9aae-1abc-444a-9db8-89f3778d11ea)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestJSONDecodeError(t *testing.T) {
	var _ MyError = JSONDecodeError{}
	var _ Causer = JSONDecodeError{}
	err0 := fmt.Errorf("test")
	err := NewJSONDecodeError(err0, "3ced8e3b-82e9-42ab-acd5-47e585bb7f11")
	if err.Cause() != err0 {
		t.Errorf("err.Cause() not expected")
	}
	if err.Hash() != "3ced8e3b-82e9-42ab-acd5-47e585bb7f11" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != "invalid JSON (3ced8e3b-82e9-42ab-acd5-47e585bb7f11)" {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}

func TestInvalidTokenError(t *testing.T) {
	var _ MyError = InvalidTokenError{}
	token := "0939d991-f6ee-4beb-997d-2d3b610b1a4f"
	err := NewInvalidTokenError(token, "d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4")
	if err.Token() != token {
		t.Errorf("err.Token not expected")
	}
	if err.Hash() != "d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4" {
		t.Errorf("err.Hash() not expected")
	}
	if err.Error() != `invalid token: "0939d991-f6ee-4beb-997d-2d3b610b1a4f" (d9d26801-f7f9-4d1f-bddb-6e9d2b1706b4)` {
		t.Errorf("err.Error() not expected")
	}
	err.ErrorDetails() // NOTE: Just call it to mark it tested
}
