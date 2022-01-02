package configs

import log "github.com/siruspen/logrus"

func ConfigureApp() {
	ConfigureGlobalTime()
	ConfigureLogger()
	log.Info("Configuration successfully completed!")
}
