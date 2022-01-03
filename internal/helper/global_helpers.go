package helper

import (
	"errors"
	"fmt"
	"os"
)

func PrintAndExit(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("FATAL ERROR: "+format, a...))
	os.Exit(1)
}

func FmtErr(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}
