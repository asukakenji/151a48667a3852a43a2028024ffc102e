package lib

import (
	"encoding/json"
	"io"
	"regexp"
	"strconv"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

const (
	integralPart   = `(?:-?|-?0|-?[1-9][0-9]*)`
	fractionalPart = `[0-9]*`
)

var (
	decimalRegex = regexp.MustCompile(
		`^(?:` +
			integralPart + `|` +
			`\.` + fractionalPart + `|` +
			integralPart + `\.` + fractionalPart +
			`)$`,
	)
)

// isLatitude checks whether s is a valid decimal.
func isDecimal(s string) bool {
	switch s {
	case "":
		return false
	case "-":
		return false
	case ".":
		return false
	case "-.":
		return false
	default:
		return decimalRegex.MatchString(s)
	}
}

// isLatitude checks whether s is a valid latitude.
func isLatitude(s string) bool {
	if !isDecimal(s) {
		return false
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -90 || f > 90 {
		return false
	}
	return true
}

// isLongitude checks whether s is a valid longitude.
func isLongitude(s string) bool {
	if !isDecimal(s) {
		return false
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	if f < -180 || f > 180 {
		return false
	}
	return true
}

// LocationsFromJSON reads from r, parses the content as JSON, and returns the
// Locations value.
func LocationsFromJSON(r io.Reader) (locs common.Locations, err common.Error) {
	_err := json.NewDecoder(r).Decode(&locs)
	if _err != nil {
		hash := common.NewToken()
		return nil, common.NewJSONDecodeError(_err, hash)
	}

	if len(locs) < 2 {
		hash := common.NewToken()
		return nil, NewInsufficientLocationCountError(locs, hash)
	}

	for i, loc := range locs {
		if len(loc) != 2 {
			hash := common.NewToken()
			return nil, NewInvalidLocationError(locs, i, hash)
		}
		if !isLatitude(loc[0]) {
			hash := common.NewToken()
			return nil, NewLatitudeError(locs, i, hash)
		}
		if !isLongitude(loc[1]) {
			hash := common.NewToken()
			return nil, NewLongitudeError(locs, i, hash)
		}
	}
	return locs, nil
}
