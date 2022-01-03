package translationator

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"translationator/internal/helper"
	"translationator/internal/translib"
	"translationator/internal/translib/langlib"
	"translationator/internal/translib/transmodels"
)

var RootCmd = &cobra.Command{
	Use:   "translationator",
	Short: "Translationator",
	Long: `This application will translib a given phrase in the English
language into many other langlib, and then back into English.
This will result in some zany output!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			helper.PrintAndExit("Expected at least two arguments as input!")
		}
		apiKey := args[0]
		textToTranslationate := args[1]
		var iterations int
		if len(args) > 2 {
			it, err := strconv.Atoi(args[2])
			if err != nil {
				helper.PrintAndExit("Expected input for iterations: %v", err)
			}
			iterations = it
		} else {
			iterations = 10
		}
		maxIterations := len(langlib.RandomizerLanguageCodes)
		if iterations < 0 || iterations > maxIterations {
			helper.PrintAndExit("Max iterations allowed: %d. iterations provided: %d", maxIterations, iterations)
		}
		translationateRequest, err := transmodels.NewTranslationateRequest(apiKey, textToTranslationate, iterations)
		if err != nil {
			helper.PrintAndExit("Failed to create a translationate request: %v", err)
		}
		resp := translib.Translationate(translationateRequest)
		fmt.Println(resp.GetTranslationatedText())
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		helper.PrintAndExit("Execution failed: %v", err)
	}
}
