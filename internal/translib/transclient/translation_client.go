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
	if len(translations) <= 0 {
		return transmodels.EmptyTranslationClientResponse(), errors.New("resulting translation array was empty")
	}
	return transmodels.NewTranslationClientResponse(request.GetText(), translations[0].Text, request.GetCurrentLanguage(), request.GetTargetLanguage()), nil
}

func Translationate(request transmodels.TranslationateRequest) (transmodels.TranslationateResponse, error) {
	// Before running anything, contruct a Google Translate Client. This speeds up the multiple requests
	ctx := context.Background()
	client, err := translate.NewClient(ctx, request.GetClientOption())
	if err != nil {
		return transmodels.EmptyTranslationateResponse(), helper.FmtErr("Failed to contstruct Google Translate transclient: %v", err)
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
			return transmodels.EmptyTranslationateResponse(), helper.FmtErr("failed to construct translation transclient request: %v", err)
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

func failTranslation(currentLanguage language.Tag, targetLanguage language.Tag, err error) (transmodels.TranslationateResponse, error) {
	return transmodels.EmptyTranslationateResponse(),
		helper.FmtErr("Failed to translate language [%s] to [%s]: %v", currentLanguage.String(), targetLanguage.String(), err)
}
