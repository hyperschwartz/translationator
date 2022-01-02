package configs

import log "github.com/siruspen/logrus"

func ConfigureApp() {
	ConfigureGlobalTime()
	ConfigureGlobalRandom()
	log.Debug("Configuration successfully completed!")
}
