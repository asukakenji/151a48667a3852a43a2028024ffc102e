package lib

import (
	"reflect"
	"strings"
	"testing"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

func TestIsDecimal(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		// -ve and +ve integers
		{"-11", true},
		{"-10", true},
		{"-9", true},
		{"-8", true},
		{"-7", true},
		{"-6", true},
		{"-5", true},
		{"-4", true},
		{"-3", true},
		{"-2", true},
		{"-1", true},
		{"-0", true},
		{"0", true},
		{"1", true},
		{"2", true},
		{"3", true},
		{"4", true},
		{"5", true},
		{"6", true},
		{"7", true},
		{"8", true},
		{"9", true},
		{"10", true},
		{"11", true},
		// Decimal point: none, left, right, middle
		{"321", true},
		{".123", true},
		{"321.", true},
		{"321.123", true},
		// Decimal point: more than one
		{"321..123", false},
		{"321.0.123", false},
		{"32.1.1.23", false},
		// Zero related
		{"012", false},
		{".210", true},
		{"012.", false},
		{"012.210", false},
		{"0.", true},
		{".0", true},
		{"0.0", true},
		// Minus, zero, and decimal point
		{"-0.0", true},
		{"-0.1", true},
		{"-.0", true},
		{"-.1", true},
		// Special cases
		{"", false},
		{"-", false},
		{".", false},
		{"-.", false},
	}
	for _, c := range cases {
		got := isDecimal(c.s)
		if got != c.expected {
			t.Errorf("isDecimal(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestIsLatitude(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"-180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", false}, // -Inf
		{"-1.8e308", false}, // -Inf
		{"-90.000001", false},
		{"-90.000000", true},
		{"22.372081", true}, // Example 1.1
		{"22.284419", true}, // Example 1.2
		{"22.326442", true}, // Example 1.3
		{"90.000000", true},
		{"90.000001", false},
		{"1e1", false},     // 10
		{"1.8e308", false}, // +Inf
		{"180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", false}, // +Inf
	}
	for _, c := range cases {
		got := isLatitude(c.s)
		if got != c.expected {
			t.Errorf("isLatitude(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestIsLongitude(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"-180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", false}, // -Inf
		{"-1.8e308", false}, // -Inf
		{"-180.000001", false},
		{"-180.000000", true},
		{"114.107877", true}, // Example 1.1
		{"114.159510", true}, // Example 1.2
		{"114.167811", true}, // Example 1.3
		{"180.000000", true},
		{"180.000001", false},
		{"1e1", false},     // 10
		{"1.8e308", false}, // +Inf
		{"180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", false}, // +Inf
	}
	for _, c := range cases {
		got := isLongitude(c.s)
		if got != c.expected {
			t.Errorf("isLongitude(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestLocationFromJSON(t *testing.T) {
	cases := []struct {
		s               string
		expectedLocs    common.Locations
		expectedErrType error
	}{
		{`{}`, nil, common.JSONDecodeError{}},
		{`[]`, nil, InsufficientLocationCountError{}},
		{`[["22.372081","114.107877"],["22.284419","114.159510","22.326442"]]`, nil, InvalidLocationError{}},
		{`[["90.372081","114.107877"],["22.284419","114.159510"]]`, nil, LatitudeError{}},
		{`[["22.372081","114.107877"],["22.284419","180.159510"]]`, nil, LongitudeError{}},
		{
			`[["22.372081","114.107877"],["22.284419","114.159510"],["22.326442","114.167811"]]`,
			common.Locations{
				{"22.372081", "114.107877"},
				{"22.284419", "114.159510"},
				{"22.326442", "114.167811"},
			},
			nil,
		},
	}
	for _, c := range cases {
		gotLocs, gotErr := LocationsFromJSON(strings.NewReader(c.s))
		if !reflect.DeepEqual(gotLocs, c.expectedLocs) || reflect.TypeOf(gotErr) != reflect.TypeOf(c.expectedErrType) {
			t.Errorf("LocationsFromJSON(strings.NewReader(%q)) = (%#v, %T), expected (%#v, %T)",
				c.s, gotLocs, gotErr,
				c.expectedLocs, c.expectedErrType,
			)
		}
	}
}
