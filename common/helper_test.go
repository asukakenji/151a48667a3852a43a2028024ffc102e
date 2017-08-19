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
