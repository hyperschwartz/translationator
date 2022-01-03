package helper

import (
	"errors"
	"fmt"
)

// FmtErr Helps in using fmt.Sprintf to format the values contained within an error.  Just a little shortcut for ease
// of use.
func FmtErr(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}
