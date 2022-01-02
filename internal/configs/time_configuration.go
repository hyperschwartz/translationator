package configs

import "os"

func ConfigureGlobalTime() {
	os.Setenv("TZ", "UTC")
}
