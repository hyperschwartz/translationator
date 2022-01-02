package translations

import (
	"cloud.google.com/go/translate"
	"context"
	log "github.com/siruspen/logrus"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"math/rand"
	"translationator/internal/translations/languages"
)

type TranslationateRequest struct {
	ApiKey     string
	Text       string
	Iterations int
}

type TranslationateResponse struct {
	OriginalText        string
	TranslationatedText string
}

type TranslationRequest struct {
	ApiKey          string
	Text            string
	CurrentLanguage language.Tag
	TargetLanguage  language.Tag
}

type TranslationResponse struct {
	OriginalText        string
	TranslatedText      string
	OriginalLanguage    language.Tag
	TranslationLanguage language.Tag
}

func Translationate(request TranslationateRequest) TranslationateResponse {
	executedLanguages := make([]language.Tag, request.Iterations)
	remainingLanguages := languages.RandomizerLanguageCodes
	currentLanguage := language.English
	currentText := request.Text
	for i := 0; i < request.Iterations-1; i++ {
		nextLanguage := remainingLanguages[rand.Intn(len(remainingLanguages))]
		executedLanguages = append(executedLanguages, nextLanguage)
		translateResponse := TranslateTextTo(TranslationRequest{
			ApiKey:          request.ApiKey,
			Text:            currentText,
			CurrentLanguage: currentLanguage,
			TargetLanguage:  nextLanguage,
		})
		currentLanguage = nextLanguage
		remainingLanguages = languages.FilterLanguageCodes(remainingLanguages, func(code language.Tag) bool {
			return languages.LanguageCodeInArray(remainingLanguages, code)
		})
		currentText = translateResponse.TranslatedText
	}
	finalResponse := TranslateTextTo(TranslationRequest{
		ApiKey:          request.ApiKey,
		Text:            currentText,
		CurrentLanguage: currentLanguage,
		TargetLanguage:  language.English,
	})
	return TranslationateResponse{
		OriginalText:        request.Text,
		TranslationatedText: finalResponse.TranslatedText,
	}
}

func TranslateTextTo(request TranslationRequest) TranslationResponse {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(request.ApiKey))
	if err != nil {
		log.Fatal("Could not create cloud client: ", err)
	}
	translations, err := client.Translate(
		ctx,
		[]string{request.Text},
		request.TargetLanguage,
		&translate.Options{
			Source: request.CurrentLanguage,
			Format: translate.Text,
		},
	)
	if err != nil {
		log.Fatal("Translation failed to succeed: ", err)
	}
	if len(translations) <= 0 {
		log.Fatal("Resulting translations was somehow empty!")
	}
	return TranslationResponse{
		OriginalText:        request.Text,
		TranslatedText:      translations[0].Text,
		OriginalLanguage:    request.CurrentLanguage,
		TranslationLanguage: request.TargetLanguage,
	}
}
