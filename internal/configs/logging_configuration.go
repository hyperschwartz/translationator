package configs

import log "github.com/siruspen/logrus"

func ConfigureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
}
