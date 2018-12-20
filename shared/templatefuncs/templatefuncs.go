package templatefuncs

import (
	"os"
	"strconv"
	"time"

	"go.isomorphicgo.org/go/isokit"
)

// RubyDate returns the current time in Ruby date format.
func RubyDate(t time.Time) string {

	layout := time.RubyDate
	return t.Format(layout)

}

// UnixTime returns the current time in Unix format.
func UnixTime(t time.Time) string {

	return strconv.FormatInt(t.Unix(), 10)

}

// IsProduction tells if the app is running in production mode.
func IsProduction() bool {

	if isokit.OperatingEnvironment() == isokit.ServerEnvironment {
		return os.Getenv("GOINGS_APP_MODE") == "production"
	} else {
		return false
	}

}
