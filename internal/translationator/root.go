package translationator

import (
	"errors"
	"fmt"
	"github.com/hyperschwartz/translationator/internal/helper"
	"github.com/hyperschwartz/translationator/internal/translib/langlib"
	"github.com/hyperschwartz/translationator/internal/translib/transclient"
	"github.com/hyperschwartz/translationator/internal/translib/transmodels"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"os"
)

var defaultApiKey = "<none specified>"
var translationCredJson = "translation-credentials.json"
var defaultFilePath = fmt.Sprintf("~/translationator/%s", translationCredJson)

// Variable flags set by cobra
var filePath string
var apiKey string
var textToTranslationate string
var iterations int
var verbose bool

var RootCmd = &cobra.Command{
	Use:   "translationator",
	Short: "A crazy translation app that converts your words into art",
	Long: `This application will translib a given phrase in the English
language into many other langlib, and then back into English.
This will result in some zany output!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var clientOption option.ClientOption
		// Prefer an API key if provided. This is a deliberate action and should be honored
		if apiKey != defaultApiKey {
			clientOption = option.WithAPIKey(apiKey)
			// If the filepath is not provided, auto-resolve the user home directory and use it as the path
		} else if filePath == defaultFilePath {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return helper.FmtErr("failed to resolve user home directory: %v", err)
			}
			clientOption = option.WithCredentialsFile(fmt.Sprintf("%s/translationator/%s", homeDir, translationCredJson))
			// Otherwise, attempt to find a credentials file at a provided file path
		} else {
			clientOption = option.WithCredentialsFile(filePath)
		}
		if len(textToTranslationate) == 0 {
			return errors.New("please enter a valid text to Translationate")
		}
		maxIterations := len(langlib.RandomizerLanguageCodes)
		if iterations <= 0 || iterations > maxIterations {
			return helper.FmtErr("invalid iteration amount provided: %d. please use a number greater than zero and less than the maximum value of %d", iterations, maxIterations)
		}
		translationateRequest, err := transmodels.NewTranslationateRequest(clientOption, textToTranslationate, iterations, verbose)
		if err != nil {
			return helper.FmtErr("failed to create a translationate request: %v", err)
		}
		resp, err := transclient.Translationate(translationateRequest)
		if err != nil {
			return helper.FmtErr("error occurred during translationate: %v", err)
		}
		fmt.Println(resp.GetTranslationatedText())
		return nil
	},
}

func Execute() {
	// Establish all flag definitions on the RootCmd before executing it
	RootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", defaultApiKey, "The google cloud apikey. Alternative to translation-credentials.json")
	RootCmd.PersistentFlags().StringVarP(&filePath, "filepath", "p", defaultFilePath, "The location of your Google Translate credential JSON")
	RootCmd.PersistentFlags().StringVarP(&textToTranslationate, "text", "t", "Time for a Translationator run!", "The text to Translationate")
	RootCmd.PersistentFlags().IntVarP(&iterations, "iterations", "i", 10, "The amount of iterations to run")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Displays active information about execution, including each translation made")
	// Run the app, and blow up if any error is found
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
