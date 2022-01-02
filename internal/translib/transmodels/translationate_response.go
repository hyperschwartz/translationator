package transmodels

type TranslationateResponse struct {
	originalText        string
	translationatedText string
}

func NewTranslationateResponse(originalText string, translationatedText string) TranslationateResponse {
	return TranslationateResponse{originalText: originalText, translationatedText: translationatedText}
}

func (t TranslationateResponse) GetOriginalText() string {
	return t.originalText
}

func (t TranslationateResponse) GetTranslationatedText() string {
	return t.translationatedText
}
