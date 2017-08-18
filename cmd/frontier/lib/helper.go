package lib

import (
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var (
	tokenRegex = regexp.MustCompile("^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}$")
)

// IsToken checks whether s is a UUID.
func IsToken(s string) bool {
	return tokenRegex.MatchString(s)
}

// NewToken generates a Version 4 UUID.
func NewToken() string {
	return uuid.NewV4().String()
}
