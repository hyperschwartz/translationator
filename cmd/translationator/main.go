package main

import (
	log "github.com/siruspen/logrus"
	"translationator/internal/configs"
)

func main() {
	configs.ConfigureApp()
	log.Info("Welcome to Translationator!")
}
