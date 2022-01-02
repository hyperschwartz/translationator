package translate

import (
	log "github.com/siruspen/logrus"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"translationator/internal/configs"
	"translationator/internal/helper"
	"translationator/internal/translate/languages"
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
	CurrentLanguage languages.LanguageCode
	TargetLanguage  languages.LanguageCode
}

type TranslationResponse struct {
	OriginalText        string
	TranslatedText      string
	OriginalLanguage    languages.LanguageCode
	TranslationLanguage languages.LanguageCode
}

func Translationate(request TranslationateRequest) TranslationateResponse {
	executedLanguages := make([]languages.LanguageCode, request.Iterations)
	remainingLanguages := languages.RandomizerLanguageCodes
	currentLanguage := languages.English
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
		remainingLanguages = languages.FilterLanguageCodes(remainingLanguages, func(code languages.LanguageCode) bool {
			return helper.ArrayContains(executedLanguages, code)
		})
		currentText = translateResponse.TranslatedText
	}
	finalResponse := TranslateTextTo(TranslationRequest{
		ApiKey:          request.ApiKey,
		Text:            currentText,
		CurrentLanguage: currentLanguage,
		TargetLanguage:  languages.English,
	})
	return TranslationateResponse{
		OriginalText:        request.Text,
		TranslationatedText: finalResponse.TranslatedText,
	}
}

func TranslateTextTo(request TranslationRequest) TranslationResponse {
	httpRequest, err := http.NewRequest("GET", configs.RapidapiUrl, nil)
	if err != nil {
		log.Panic("Failed to properly format request")
	}
	queryParams := httpRequest.URL.Query()
	queryParams.Add("text", request.Text)
	queryParams.Add("to", string(request.TargetLanguage))
	queryParams.Add("from", string(request.CurrentLanguage))
	httpRequest.URL.RawQuery = queryParams.Encode()
	httpRequest.Header.Add("x-rapidapi-host", "nlp-translation.p.rapidapi.com")
	httpRequest.Header.Add("x-rapidapi-key", request.ApiKey)
	response, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Panic("Failed to translate text and received error:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Failed to defer closure of request body:", err)
		}
	}(response.Body)
	body, _ := ioutil.ReadAll(response.Body)
	log.Info(string(body))
	return TranslationResponse{
		OriginalText:        request.Text,
		TranslatedText:      "idk",
		OriginalLanguage:    request.CurrentLanguage,
		TranslationLanguage: request.TargetLanguage,
	}
}
