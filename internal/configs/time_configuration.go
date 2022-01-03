package configs

import (
	"fmt"
	"os"
)

// ConfigureGlobalTime Ensures that the application is set to UTC time for standardization
func ConfigureGlobalTime() {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		fmt.Printf("Failed to alter timezone to UTC: %v", err)
	}
}
