package configs

import (
	log "github.com/siruspen/logrus"
	"os"
)

func ConfigureGlobalTime() {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		log.Error("Failed to alter timezone to UTC:", err)
	}
}
