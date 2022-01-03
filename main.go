package main

import (
	"translationator/cmd/translationator"
	"translationator/internal/configs"
)

func main() {
	// Establish all configurations before doing other operations (enable randomness, UTC timezone, etc)
	configs.ConfigureApp()
	// Attempt to parse CLI input with Cobra
	translationator.Execute()
}
