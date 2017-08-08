package common

import (
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var (
	tokenRegex = regexp.MustCompile("^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}$")
)

func IsToken(s string) bool {
	return tokenRegex.MatchString(s)
}

func NewToken() string {
	return uuid.NewV4().String()
}
