package main

import (
	"translationator/cmd/translationator"
	"translationator/internal/configs"
)

func main() {
	configs.ConfigureApp()
	translationator.Execute()
}
