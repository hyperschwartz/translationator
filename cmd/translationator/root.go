package translationator

import (
	log "github.com/siruspen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"translationator/internal/translations"
)

var RootCmd = &cobra.Command{
	Use:   "translationator",
	Short: "Translationator",
	Long: `This application will translations a given phrase in the English
language into many other languages, and then back into English.
This will result in some zany output!`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("Expected at least two arguments as input!")
		}
		apiKey := args[0]
		textToTranslationate := args[1]
		var iterations int
		if len(args) > 2 {
			it, err := strconv.Atoi(args[2])
			if err != nil {
				log.Fatal("Unexpected input for iterations:", err)
			}
			iterations = it
		} else {
			iterations = 10
		}
		if iterations < 0 || iterations > 25 {
			log.Fatal("Max iterations allowed: 25. Iterations provided:", iterations)
		}
		resp := translations.Translationate(translations.TranslationateRequest{
			ApiKey:     apiKey,
			Text:       textToTranslationate,
			Iterations: iterations,
		})
		log.Info("Translationated: ", resp.TranslationatedText)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
