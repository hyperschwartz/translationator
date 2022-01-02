package transmodels

import (
	"errors"
	"strconv"
	"translationator/internal/translib/langlib"
)

type TranslationateRequest struct {
	apiKey     string
	text       string
	iterations int
}

func NewTranslationateRequest(apiKey string, text string, iterations int) (TranslationateRequest, error) {
	if iterations < 1 {
		return TranslationateRequest{}, errors.New("Must specify iteration amount of at least 1. Value specified: " + strconv.Itoa(iterations))
	}
	maxIterations := len(langlib.RandomizerLanguageCodes)
	if iterations > maxIterations {
		return TranslationateRequest{}, errors.New("There are only [" + strconv.Itoa(maxIterations) + "] total iterations possible")
	}
	return TranslationateRequest{apiKey: apiKey, text: text, iterations: iterations}, nil
}

func (t TranslationateRequest) GetApiKey() string {
	return t.apiKey
}

func (t TranslationateRequest) GetText() string {
	return t.text
}

func (t TranslationateRequest) GetIterations() int {
	return t.iterations
}
