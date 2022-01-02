package translate

import (
	"fmt"
	"translationator/internal/translate/languages"
)

type TranslationRequest struct {
	Text            string
	CurrentLanguage languages.LanguageCode
	TargetLanguage  languages.LanguageCode
}

func TranslateTextTo(request TranslationRequest) {
	for _, code := range languages.RandomizerLanguageCodes {
		fmt.Println(code)
	}
}
