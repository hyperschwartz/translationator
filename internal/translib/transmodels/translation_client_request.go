package transmodels

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/hyperschwartz/translationator/internal/helper"
	"golang.org/x/text/language"
)

type TranslationClientRequest struct {
	client          *translate.Client
	ctx             context.Context
	text            string
	currentLanguage language.Tag
	targetLanguage  language.Tag
}

func NewTranslationClientRequest(
	client *translate.Client,
	ctx context.Context,
	text string,
	currentLanguage language.Tag,
	targetLanguage language.Tag,
) (TranslationClientRequest, error) {
	if currentLanguage == targetLanguage {
		return EmptyTranslationClientRequest(), helper.FmtErr(
			"Invalid request. Current language [%s] must differ from target language [%s]",
			currentLanguage.String(),
			targetLanguage.String(),
		)
	}
	return TranslationClientRequest{
		client:          client,
		ctx:             ctx,
		text:            text,
		currentLanguage: currentLanguage,
		targetLanguage:  targetLanguage,
	}, nil
}

func (t TranslationClientRequest) GetGoogleClient() *translate.Client {
	return t.client
}

func (t TranslationClientRequest) GetGoogleContext() context.Context {
	return t.ctx
}

func (t TranslationClientRequest) GetText() string {
	return t.text
}

func (t TranslationClientRequest) GetCurrentLanguage() language.Tag {
	return t.currentLanguage
}

func (t TranslationClientRequest) GetTargetLanguage() language.Tag {
	return t.targetLanguage
}

func EmptyTranslationClientRequest() TranslationClientRequest {
	return TranslationClientRequest{}
}
