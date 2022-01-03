package transmodels

import (
	"errors"
	"fmt"
	"strconv"
	"translationator/internal/helper"
	"translationator/internal/translib/langlib"
)

type TranslationateRequest struct {
	apiKey     string
	text       string
	iterations int
	verbose    bool
}

func NewTranslationateRequest(apiKey string, text string, iterations int, verbose bool) (TranslationateRequest, error) {
	if iterations < 1 {
		return EmptyTranslationateRequest(), errors.New("Must specify iteration amount of at least 1. Value specified: " + strconv.Itoa(iterations))
	}
	maxIterations := len(langlib.RandomizerLanguageCodes)
	if iterations > maxIterations {
		return EmptyTranslationateRequest(), helper.FmtErr("There are only [%d] total iterations possible", maxIterations)
	}
	return TranslationateRequest{apiKey: apiKey, text: text, iterations: iterations, verbose: verbose}, nil
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

func (t TranslationateRequest) GetVerbose() bool {
	return t.verbose
}

func (t TranslationateRequest) PrintIfVerbose(text string) {
	if t.verbose {
		fmt.Println(text)
	}
}

func (t TranslationateRequest) FmtIfVerbose(format string, a ...interface{}) {
	if t.verbose {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

func EmptyTranslationateRequest() TranslationateRequest {
	return TranslationateRequest{}
}
