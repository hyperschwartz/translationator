package configs

import log "github.com/siruspen/logrus"

func ConfigureApp() {
	ConfigureLogger()
	ConfigureGlobalTime()
	log.Debug("Configuration successfully completed!")
}
