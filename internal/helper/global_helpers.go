package helper

import (
	"fmt"
	"os"
)

func PrintAndExit(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("FATAL ERROR: "+format, a...))
	os.Exit(1)
}
