package translationator

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"os"
	"translationator/internal/helper"
	"translationator/internal/translib/langlib"
	"translationator/internal/translib/transclient"
	"translationator/internal/translib/transmodels"
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
				helper.PrintAndExit("Failed to resolve user home directory: %v", err)
			}
			clientOption = option.WithCredentialsFile(fmt.Sprintf("%s/translationator/%s", homeDir, translationCredJson))
			// Otherwise, attempt to find a credentials file at a provided file path
		} else {
			clientOption = option.WithCredentialsFile(filePath)
		}
		if len(textToTranslationate) == 0 {
			helper.PrintAndExit("Please enter a valid text to Translationate")
		}
		maxIterations := len(langlib.RandomizerLanguageCodes)
		if iterations <= 0 || iterations > maxIterations {
			helper.PrintAndExit("Invalid iteration amount provided: %d. Please use a number greater than zero and less than the maximum value of %d", iterations, maxIterations)
		}
		translationateRequest, err := transmodels.NewTranslationateRequest(clientOption, textToTranslationate, iterations, verbose)
		if err != nil {
			helper.PrintAndExit("Failed to create a translationate request: %v", err)
		}
		resp, err := transclient.Translationate(translationateRequest)
		if err != nil {
			helper.PrintAndExit("Error occurred during translationate: %v", err)
		}
		fmt.Println(resp.GetTranslationatedText())
		return nil
	},
}

func Execute() {
	RootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", defaultApiKey, "The google cloud apikey. Alternative to translation-credentials.json")
	RootCmd.PersistentFlags().StringVarP(&filePath, "filepath", "p", defaultFilePath, "The location of your Google Translate credential JSON")
	RootCmd.PersistentFlags().StringVarP(&textToTranslationate, "text", "t", "Time for a Translationator run!", "The text to Translationate")
	RootCmd.PersistentFlags().IntVarP(&iterations, "iterations", "i", 10, "The amount of iterations to run")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Displays active information about execution, including each translation made")
	if err := RootCmd.Execute(); err != nil {
		helper.PrintAndExit("Execution failed: %v", err)
	}
}
