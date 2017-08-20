package lib

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

func TestInsufficientLocationCountError(t *testing.T) {
	name := "InsufficientLocationCountError"
	hash := "9240926c-d38d-4683-aeb1-c0c76cc85479"
	locs := common.Locations{}
	err := NewInsufficientLocationCountError(locs, hash)
	var _ common.Error = err
	if !reflect.DeepEqual(err.Locations(), locs) {
		t.Errorf("err.Locations() not expected")
	}
	if !strings.HasSuffix(err.Error(), fmt.Sprintf("(%s)", hash)) {
		t.Errorf("err.Error() not expected")
	}
	if !strings.HasPrefix(err.ErrorDetails(), fmt.Sprintf("[%s] %s", hash, name)) {
		t.Errorf("err.ErrorDetails() not expected")
	}
}

func TestInvalidLocationError(t *testing.T) {
	name := "InvalidLocationError"

	hash0 := "b10efaef-98f0-4ed6-9067-fdc443e8fafa"
	locs0 := common.Locations{
		{"22.284419", "114.159510", "22.326442"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := NewInvalidLocationError(locs0, index0, hash0)
	var _ common.Error = err0
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if !strings.HasSuffix(err0.Error(), fmt.Sprintf("(%s)", hash0)) {
		t.Errorf("err0.Error() not expected")
	}
	if !strings.HasPrefix(err0.ErrorDetails(), fmt.Sprintf("[%s] %s", hash0, name)) {
		t.Errorf("err0.ErrorDetails() not expected")
	}

	hash1 := "7ce896d9-3c3e-48d2-ac4f-2d2950f0ab09"
	locs1 := common.Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "114.159510", "22.326442"},
	}
	index1 := 1
	err1 := NewInvalidLocationError(locs1, index1, hash1)
	var _ common.Error = err1
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if !strings.HasSuffix(err1.Error(), fmt.Sprintf("(%s)", hash1)) {
		t.Errorf("err1.Error() not expected")
	}
	if !strings.HasPrefix(err1.ErrorDetails(), fmt.Sprintf("[%s] %s", hash1, name)) {
		t.Errorf("err1.ErrorDetails() not expected")
	}
}

func TestLatitudeError(t *testing.T) {
	name := "LatitudeError"

	hash0 := "acc75ebd-c145-40fa-9cc6-1f1337d4b8ae"
	locs0 := common.Locations{
		{"90.372081", "114.107877"},
		{"22.284419", "114.159510"},
	}
	index0 := 0
	err0 := NewLatitudeError(locs0, index0, hash0)
	var _ common.Error = err0
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if !strings.HasSuffix(err0.Error(), fmt.Sprintf("(%s)", hash0)) {
		t.Errorf("err0.Error() not expected")
	}
	if !strings.HasPrefix(err0.ErrorDetails(), fmt.Sprintf("[%s] %s", hash0, name)) {
		t.Errorf("err0.ErrorDetails() not expected")
	}

	hash1 := "2311e1c7-ceb9-43e6-9d02-029e8a3a6541"
	locs1 := common.Locations{
		{"22.284419", "114.159510"},
		{"90.372081", "114.107877"},
	}
	index1 := 1
	err1 := NewLatitudeError(locs1, index1, hash1)
	var _ common.Error = err1
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if !strings.HasSuffix(err1.Error(), fmt.Sprintf("(%s)", hash1)) {
		t.Errorf("err1.Error() not expected")
	}
	if !strings.HasPrefix(err1.ErrorDetails(), fmt.Sprintf("[%s] %s", hash1, name)) {
		t.Errorf("err1.ErrorDetails() not expected")
	}
}

func TestLongitudeError(t *testing.T) {
	name := "LongitudeError"

	hash0 := "c5727bc8-44e8-4e4e-b378-35fbe2604afc"
	locs0 := common.Locations{
		{"22.284419", "180.159510"},
		{"22.372081", "114.107877"},
	}
	index0 := 0
	err0 := NewLongitudeError(locs0, index0, hash0)
	var _ common.Error = err0
	if !reflect.DeepEqual(err0.Locations(), locs0) {
		t.Errorf("err0.Locations() not expected")
	}
	if err0.Index() != 0 {
		t.Errorf("err0.Index() not expected")
	}
	if !strings.HasSuffix(err0.Error(), fmt.Sprintf("(%s)", hash0)) {
		t.Errorf("err0.Error() not expected")
	}
	if !strings.HasPrefix(err0.ErrorDetails(), fmt.Sprintf("[%s] %s", hash0, name)) {
		t.Errorf("err0.ErrorDetails() not expected")
	}

	hash1 := "2a965bd4-ea7f-4351-9cb5-8c38f42178a0"
	locs1 := common.Locations{
		{"22.372081", "114.107877"},
		{"22.284419", "180.159510"},
	}
	index1 := 1
	err1 := NewLongitudeError(locs1, index1, hash1)
	var _ common.Error = err1
	if !reflect.DeepEqual(err1.Locations(), locs1) {
		t.Errorf("err1.Locations() not expected")
	}
	if err1.Index() != 1 {
		t.Errorf("err1.Index() not expected")
	}
	if !strings.HasSuffix(err1.Error(), fmt.Sprintf("(%s)", hash1)) {
		t.Errorf("err1.Error() not expected")
	}
	if !strings.HasPrefix(err1.ErrorDetails(), fmt.Sprintf("[%s] %s", hash1, name)) {
		t.Errorf("err1.ErrorDetails() not expected")
	}
}
