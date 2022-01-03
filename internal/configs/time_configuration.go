package configs

import (
	"fmt"
	"os"
)

func ConfigureGlobalTime() {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		fmt.Printf("Failed to alter timezone to UTC: %v", err)
	}
}
