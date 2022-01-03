package transclient

import (
	"cloud.google.com/go/translate"
	"context"
	"errors"
	"golang.org/x/text/language"
	"math/rand"
	"translationator/internal/helper"
	"translationator/internal/translib/langlib"
	"translationator/internal/translib/transmodels"
)

// translateTextTo takes a TranslationClientRequest and attempts to invoke the Google Translate API to translate
// a phrase from one language to another.  This process is made very simple due to Google providing such an easy to use
// client for their API.
func translateTextTo(request transmodels.TranslationClientRequest) (transmodels.TranslationClientResponse, error) {
	client := request.GetGoogleClient()
	translations, err := client.Translate(
		request.GetGoogleContext(),
		[]string{request.GetText()},
		request.GetTargetLanguage(),
		&translate.Options{
			Source: request.GetCurrentLanguage(),
			Format: translate.Text,
		},
	)
	if err != nil {
		return transmodels.EmptyTranslationClientResponse(), helper.FmtErr("Translation failed: %v", err)
	}
	// The API responds with one translation per value in the array provided to "input."  The expected response is always
	// to include a single value, so we just need to ensure that at least one value is returned
	if len(translations) <= 0 {
		return transmodels.EmptyTranslationClientResponse(), errors.New("resulting translation array was empty")
	}
	return transmodels.NewTranslationClientResponse(request.GetText(), translations[0].Text, request.GetCurrentLanguage(), request.GetTargetLanguage()), nil
}

// failTranslation is a simple helper to return an error when a translation from one language to another fails for any
// reason.
func failTranslation(currentLanguage language.Tag, targetLanguage language.Tag, err error) (transmodels.TranslationateResponse, error) {
	return transmodels.EmptyTranslationateResponse(),
		helper.FmtErr("Failed to translate language [%s] to [%s]: %v", currentLanguage.String(), targetLanguage.String(), err)
}

// Translationate The core function for Translationator.  Takes all input specified via CLI and runs through as many
// iterations as requested until arriving at a final English re-translation.
func Translationate(request transmodels.TranslationateRequest) (transmodels.TranslationateResponse, error) {
	// Before running anything, contruct a Google Translate Client. This speeds up the multiple requests
	ctx := context.Background()
	client, err := translate.NewClient(ctx, request.GetClientOption())
	if err != nil {
		return transmodels.EmptyTranslationateResponse(), helper.FmtErr("Failed to construct Google Translate client: %v", err)
	}
	request.PrintIfVerbose("Successfully established Google Translate Client")
	remainingLanguages := langlib.RandomizerLanguageCodes
	currentLanguage := language.English
	currentText := request.GetText()
	request.FmtIfVerbose("Initial Text: [%s]", currentText)
	for i := 0; i < request.GetIterations(); i++ {
		nextLanguage := remainingLanguages[rand.Intn(len(remainingLanguages))]
		transClientReq, err := transmodels.NewTranslationClientRequest(client, ctx, currentText, currentLanguage, nextLanguage)
		if err != nil {
			return transmodels.EmptyTranslationateResponse(), helper.FmtErr("Failed to construct translation client request: %v", err)
		}
		translateResponse, err := translateTextTo(transClientReq)
		if err != nil {
			return failTranslation(currentLanguage, nextLanguage, err)
		}
		currentLanguage = nextLanguage
		// Remove the target language from the remaining languages array, ensuring languages are only used once
		remainingLanguages = langlib.FilterLanguageCodes(remainingLanguages, func(code language.Tag) bool {
			return code != currentLanguage
		})
		currentText = translateResponse.GetTranslatedText()
		request.FmtIfVerbose("Iteration [%d], Language: [%s], Result: [%s]", i+1, currentLanguage.String(), currentText)
	}
	finalRequest, err := transmodels.NewTranslationClientRequest(client, ctx, currentText, currentLanguage, language.English)
	if err != nil {
		return transmodels.EmptyTranslationateResponse(), helper.FmtErr("Failed to make final re-tralsnation to English: %v", err)
	}
	finalResponse, err := translateTextTo(finalRequest)
	if err != nil {
		return failTranslation(currentLanguage, language.English, err)
	}
	return transmodels.NewTranslationateResponse(request.GetText(), finalResponse.GetTranslatedText()), nil
}
