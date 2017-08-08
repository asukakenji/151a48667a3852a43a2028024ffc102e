package common

import (
	"regexp"
	"strconv"

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

// TODO: Check whether NaN, Inf, etc. are accepted in ParseFloat.
// TODO: Check exactly 6 digits after decimal point.
// -90 (South Pole) <= latitude <= +90 (North Pole)
func IsLatitude(s string) bool {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -90 || f > 90 {
		return false
	}
	return true
}

// TODO: Check whether NaN, Inf, etc. are accepted in ParseFloat.
// TODO: Check exactly 6 digits after decimal point.
// -180 (West) <= longitude <= +180 (East)
func IsLongitude(s string) bool {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -180 || f > 180 {
		return false
	}
	return true
}
