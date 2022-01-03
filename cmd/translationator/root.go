package translationator

import (
	"fmt"
	"github.com/spf13/cobra"
	"translationator/internal/helper"
	"translationator/internal/translib"
	"translationator/internal/translib/langlib"
	"translationator/internal/translib/transmodels"
)

var defaultApiKey = "<none specified>"

var RootCmd = &cobra.Command{
	Use:   "translationator",
	Short: "A crazy translation app that converts your words into art",
	Long: `This application will translib a given phrase in the English
language into many other langlib, and then back into English.
This will result in some zany output!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, err := cmd.Flags().GetString("apikey")
		if err != nil || apiKey == defaultApiKey {
			helper.PrintAndExit("Please enter an apikey to run Translationator")
		}
		textToTranslationate, err := cmd.Flags().GetString("text")
		if err != nil || len(textToTranslationate) == 0 {
			helper.PrintAndExit("Please enter a valid text to Translationate")
		}
		iterations, err := cmd.Flags().GetInt("iterations")
		if err != nil {
			helper.PrintAndExit("Could not get a proper value for iterations")
		}
		maxIterations := len(langlib.RandomizerLanguageCodes)
		if iterations <= 0 || iterations > maxIterations {
			helper.PrintAndExit("Invalid iteration amount provided: %d. Please use a number greater than zero and less than the maximum value of %d", iterations, maxIterations)
		}
		translationateRequest, err := transmodels.NewTranslationateRequest(apiKey, textToTranslationate, iterations)
		if err != nil {
			helper.PrintAndExit("Failed to create a translationate request: %v", err)
		}
		resp, err := translib.Translationate(translationateRequest)
		if err != nil {
			helper.PrintAndExit("Error occurred during translationate: %v", err)
		}
		fmt.Println(resp.GetTranslationatedText())
		return nil
	},
}

func Execute() {
	RootCmd.PersistentFlags().StringP("apikey", "a", defaultApiKey, "The google cloud apikey")
	RootCmd.PersistentFlags().StringP("text", "t", "Time for a Translationator run!", "The text to Translationate")
	RootCmd.PersistentFlags().IntP("iterations", "i", 10, "The amount of iterations to run")
	if err := RootCmd.Execute(); err != nil {
		helper.PrintAndExit("Execution failed: %v", err)
	}
}
