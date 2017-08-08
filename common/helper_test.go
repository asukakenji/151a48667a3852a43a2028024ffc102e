package common

import "testing"

func TestIsToken(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"1234567-89ab-cdef-0123-456789abcdef", false},
		{"01234567-89ab-cdef-0123-456789abcde", false},
		{"01234567-89ab-cdef-0123-456789abcdef", true},
		{"01234567-89ab-cdef-0123-456789abcdeff", false},
		{"f01234567-89ab-cdef-0123-456789abcdeff", false},
		{"f01234567-89ab-cdef-0123-456789abcdeff", false},
	}
	for _, c := range cases {
		got := IsToken(c.s)
		if got != c.expected {
			t.Errorf("IsToken(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestNewToken(t *testing.T) {
	for i := 0; i < 65536; i++ {
		got := NewToken()
		if !IsToken(got) {
			t.Errorf("NewToken() = %#v, IsToken(%#v) = false", got, got)
		}
	}
}

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
		got := IsLatitude(c.s)
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
		got := IsLongitude(c.s)
		if got != c.expected {
			t.Errorf("IsLatitude(%#v) = %t, expected %t", c.s, got, c.expected)
		}
	}
}
