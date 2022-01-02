package translationator

import (
	log "github.com/siruspen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "translationator",
	Short: "Translationator",
	Long: `This application will translate a given phrase in the English
language into many other languages, and then back into English.
This will result in some zany output!`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Args found", args)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
