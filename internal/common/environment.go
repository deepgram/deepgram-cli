package common

import (
	"os"
)

// base url defaults to community-local.deepgram.com:3000
var BaseUrl = func() string {
	envBaseUrl := os.Getenv("DEEPGRAM_CLI_BASE_URL")
	if envBaseUrl != "" {
		return envBaseUrl
	}
	return "https://community.deepgram.com"
}()
