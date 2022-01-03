package translib

import (
	"cloud.google.com/go/translate"
	"context"
	"errors"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
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
		return transmodels.TranslationClientResponse{}, errors.New("Translation failed: " + err.Error())
	}
	if len(translations) <= 0 {
		return transmodels.TranslationClientResponse{}, errors.New("Resulting translation array was empty")
	}
	return transmodels.NewTranslationClientResponse(request.GetText(), translations[0].Text, request.GetCurrentLanguage(), request.GetTargetLanguage()), nil
}

func Translationate(request transmodels.TranslationateRequest) transmodels.TranslationateResponse {
	// Before running anything, contruct a Google Translate Client. This speeds up the multiple requests
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(request.GetApiKey()))
	if err != nil {
		helper.PrintAndExit("Failed to contstruct Google Translate client: %v", err)
	}
	executedLanguages := make([]language.Tag, request.GetIterations())
	remainingLanguages := langlib.RandomizerLanguageCodes
	currentLanguage := language.English
	currentText := request.GetText()
	for i := 0; i < request.GetIterations()-1; i++ {
		nextLanguage := remainingLanguages[rand.Intn(len(remainingLanguages))]
		executedLanguages = append(executedLanguages, nextLanguage)
		transClientReq, err := transmodels.NewTranslationClientRequest(*client, ctx, currentText, currentLanguage, nextLanguage)
		if err != nil {
			helper.PrintAndExit("Failed to construct translation client request: %v", err)
		}
		translateResponse, err := translateTextTo(transClientReq)
		if err != nil {
			failTranslation(currentLanguage, nextLanguage, err)
		}
		currentLanguage = nextLanguage
		remainingLanguages = langlib.FilterLanguageCodes(remainingLanguages, func(code language.Tag) bool {
			return langlib.LanguageCodeInArray(remainingLanguages, code)
		})
		currentText = translateResponse.GetTranslatedText()
	}
	finalRequest, err := transmodels.NewTranslationClientRequest(*client, ctx, currentText, currentLanguage, language.English)
	if err != nil {
		helper.PrintAndExit("Failed to make final re-translation to English: %v", err)
	}
	finalResponse, err := translateTextTo(finalRequest)
	if err != nil {
		failTranslation(currentLanguage, language.English, err)
	}
	return transmodels.NewTranslationateResponse(request.GetText(), finalResponse.GetTranslatedText())
}

func failTranslation(currentLanguage language.Tag, targetLanguage language.Tag, err error) {
	helper.PrintAndExit("Failed to translate language [%s] to [%s]: %v", currentLanguage.String(), targetLanguage.String(), err)
}
