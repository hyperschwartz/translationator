package main

import (
	"translationator/internal/configs"
	"translationator/internal/translate"
	"translationator/internal/translate/languages"
)

func main() {
	configs.ConfigureApp()
	//translationator.Execute()
	translate.TranslateTextTo(translate.TranslationRequest{
		ApiKey:          "insert-rapidapi-key-here",
		Text:            "Hello",
		CurrentLanguage: languages.English,
		TargetLanguage:  languages.Spanish,
	})
}
