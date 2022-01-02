package configs

import log "github.com/siruspen/logrus"

func ConfigureApp() {
	ConfigureGlobalTime()
	log.Debug("Configuration successfully completed!")
}
