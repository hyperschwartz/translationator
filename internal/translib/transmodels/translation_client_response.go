package transmodels

import "golang.org/x/text/language"

type TranslationClientResponse struct {
	originalText        string
	translatedText      string
	originalLanguage    language.Tag
	translationLanguage language.Tag
}

func NewTranslationClientResponse(originalText string, translatedText string, originalLanguage language.Tag, translationLanguage language.Tag) TranslationClientResponse {
	return TranslationClientResponse{
		originalText:        originalText,
		translatedText:      translatedText,
		originalLanguage:    originalLanguage,
		translationLanguage: translationLanguage,
	}
}

func (t TranslationClientResponse) GetOriginalText() string {
	return t.originalText
}

func (t TranslationClientResponse) GetTranslatedText() string {
	return t.translatedText
}

func (t TranslationClientResponse) GetOriginalLanguage() language.Tag {
	return t.originalLanguage
}

func (t TranslationClientResponse) GetTranslationLanguage() language.Tag {
	return t.translationLanguage
}
