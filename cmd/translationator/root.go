package translationator

import (
	"fmt"
	"github.com/spf13/cobra"
	"translationator/internal/helper"
	"translationator/internal/translib/langlib"
	"translationator/internal/translib/transclient"
	"translationator/internal/translib/transmodels"
)

var defaultApiKey = "<none specified>"

// Variable flags set by cobra
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
		if apiKey == defaultApiKey {
			helper.PrintAndExit("Please enter an apikey to run Translationator")
		}
		if len(textToTranslationate) == 0 {
			helper.PrintAndExit("Please enter a valid text to Translationate")
		}
		maxIterations := len(langlib.RandomizerLanguageCodes)
		if iterations <= 0 || iterations > maxIterations {
			helper.PrintAndExit("Invalid iteration amount provided: %d. Please use a number greater than zero and less than the maximum value of %d", iterations, maxIterations)
		}
		translationateRequest, err := transmodels.NewTranslationateRequest(apiKey, textToTranslationate, iterations, verbose)
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
	RootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", defaultApiKey, "The google cloud apikey")
	RootCmd.PersistentFlags().StringVarP(&textToTranslationate, "text", "t", "Time for a Translationator run!", "The text to Translationate")
	RootCmd.PersistentFlags().IntVarP(&iterations, "iterations", "i", 10, "The amount of iterations to run")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Displays active information about execution, including each translation made")
	if err := RootCmd.Execute(); err != nil {
		helper.PrintAndExit("Execution failed: %v", err)
	}
}
