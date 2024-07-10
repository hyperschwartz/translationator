package transmodels

import (
	"errors"
	"fmt"
	"github.com/hyperschwartz/translationator/internal/helper"
	"github.com/hyperschwartz/translationator/internal/translib/langlib"
	"google.golang.org/api/option"
	"strconv"
)

type TranslationateRequest struct {
	clientOption option.ClientOption
	text         string
	iterations   int
	verbose      bool
}

func NewTranslationateRequest(clientOption option.ClientOption, text string, iterations int, verbose bool) (TranslationateRequest, error) {
	if iterations < 1 {
		return EmptyTranslationateRequest(), errors.New("Must specify iteration amount of at least 1. Value specified: " + strconv.Itoa(iterations))
	}
	maxIterations := len(langlib.RandomizerLanguageCodes)
	if iterations > maxIterations {
		return EmptyTranslationateRequest(), helper.FmtErr("There are only [%d] total iterations possible", maxIterations)
	}
	return TranslationateRequest{clientOption: clientOption, text: text, iterations: iterations, verbose: verbose}, nil
}

func (t TranslationateRequest) GetClientOption() option.ClientOption {
	return t.clientOption
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
