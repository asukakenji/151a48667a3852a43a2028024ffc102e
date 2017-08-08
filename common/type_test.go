package common

import "testing"

func TestIsLatitude(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
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
	}
	for _, c := range cases {
		got := isLatitude(c.s)
		if got != c.expected {
			t.Errorf("IsLatitude(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestIsLongitude(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
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
	}
	for _, c := range cases {
		got := isLongitude(c.s)
		if got != c.expected {
			t.Errorf("IsLatitude(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}
