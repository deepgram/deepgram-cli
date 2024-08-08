package auth

import (
	"deepgram-cli/internal/common"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Guard(cmd *cobra.Command, args []string) {
	apiKey := viper.GetString("api_key")

	if apiKey == "" {
		fmt.Fprintln(os.Stderr, `DEEPGRAM_API_KEY is not set in the configuration file `+common.DefaultConfigText+` or environment variable.

Run "deepgram login" to configure your API key.
			`)

		log.Fatal(`DEEPGRsAM_API_KEY is not set in the configuration file or environment variable.`)
	}
}
