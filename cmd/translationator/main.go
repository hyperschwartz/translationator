package main

import (
	"github.com/hyperschwartz/translationator/internal/configs"
	"github.com/hyperschwartz/translationator/internal/translationator"
)

func main() {
	// Establish all configurations before doing other operations (enable randomness, UTC timezone, etc)
	configs.ConfigureApp()
	// Attempt to parse CLI input with Cobra
	translationator.Execute()
}
