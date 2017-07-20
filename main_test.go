package main

import "testing"

func TestIsToken(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"01234567-0123-0123-0123-0123456789ab", true},
		{"f01234567-0123-0123-0123-0123456789abf", false},
	}
	for _, c := range cases {
		got := IsToken(c.s)
		if got != c.expected {
			t.Errorf("IsToken(%s) = %t, expected %t", c.s, got, c.expected)
		}
	}
}

func TestNewToken(t *testing.T) {
	for i := 0; i < 65536; i++ {
		got := NewToken()
		if !IsToken(got) {
			t.Errorf("NewToken() = %s, IsToken(%s) = false", got, got)
		}
	}
}
