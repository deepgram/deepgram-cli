/*
Copyright © 2024 Deepgram

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"deepgram-cli/internal/auth"
	"deepgram-cli/internal/common"
	"deepgram-cli/internal/config"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in a user",
	Long:  `Logs a user into Deepgram with browser-based authentication.`,
	Run:   runLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)
	// Allow users to provide an API key directly during login.
	// This also removes the global scope api_key flag from the --help menu
	// to reduce confusion and provide a different description in this context.
	// This isn't bound to the loaded config "api_key" key, so we can tell if a
	// new key is being provided.
	loginCmd.Flags().StringVarP(&ApiKey, "api_key", "k", "", "Configure the CLI with your Deepgram API key")

	// allow users to force write and skip the prompt
	loginCmd.Flags().BoolP("force-write", "f", false, "Don't prompt for confirmation when providing an API key")
}

func runLogin(cmd *cobra.Command, args []string) {
	var (
		// Get the API key from the config context
		key string = cmd.Flags().Lookup("api_key").Value.String()
	)

	switch {
	case key != "":
		cobra.CheckErr(cliAuth())
	default:
		cobra.CheckErr(webAuth())
	}
}

func cliAuth() error {
	if !viper.GetBool("force-write") {
		cobra.CheckErr(common.PromptBool("Do you want to write this key to config?"))

		if config.ConfigFileExists() {
			cobra.CheckErr(common.PromptBool("Configuration file already exists. Overwrite?"))
		}
	}

	return config.WriteConfigFile()
}

func randomClientID(length int) string {
	const urlFriendlyChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
	deviceCode := make([]byte, length)

	for i := range deviceCode {
		deviceCode[i] = urlFriendlyChars[rand.Intn(len(urlFriendlyChars))]
	}

	return string(deviceCode)
}

func webAuth() error {
	var (
		key string = viper.GetString("api_key")

		verificationURI string
		deviceCode      auth.DeviceCodeResponse
		newKey          *auth.AccessTokenResponse
		err             error
	)

	// check if the user is already logged in - if so, prompt to ensure they want to log in again
	if key != "" {
		err = common.PromptBool("You're already logged in. Do you want to login again?")
		if err != nil {
			return err
		}
	}

	// use a hostname to identify the device
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	// create a random client ID that will be used for both ends of the device auth flow
	clientID := randomClientID(40)

	// get the device code from our community auth endpoint
	deviceCode, err = auth.RequestDeviceCode(clientID, hostname, []string{"usage:write"})
	if err != nil {
		return err
	}

	// format the verification URI with a device code
	verificationURI = fmt.Sprintf(
		"%s?device_code=%s&client_id=%s&hostname=%s",
		deviceCode.VerificationURI,
		url.QueryEscape(deviceCode.DeviceCode),
		url.QueryEscape(clientID),
		url.QueryEscape(hostname),
	)

	// format the prompt message for our user
	prompt := fmt.Sprintf(
		"%s\n\n%s%s\n\n",
		"Hello from Deepgram! Press Enter to open browser and login automatically.",
		common.MutedMessage("Here is your login link in case browser did not open: "),
		common.MutedMessage(verificationURI),
	)

	// start polling in a goroutine
	pollDone := make(chan struct{})
	var pollErr error
	go func() {
		defer close(pollDone)
		newKey, pollErr = auth.PollForAccessToken(clientID, hostname, deviceCode.DeviceCode, deviceCode.Interval)
	}()

	// wait for the user to press Enter in a separate goroutine
	enterPressed := make(chan struct{})
	var enterErr error
	var openErr error
	go func() {
		defer close(enterPressed)
		if enterErr = common.PromptEnter(prompt); enterErr == nil {
			openErr = open.Run(verificationURI)
		}
	}()

	// Wait for either the polling to finish or the user to press Enter
	select {
	case <-pollDone:
		if pollErr != nil {
			return pollErr
		}
	case <-enterPressed:
		if enterErr != nil {
			return enterErr
		}

		if openErr != nil {
			return openErr
		}

		// The user pressed Enter and potentially opened the browser, but we still wait for polling
		<-pollDone
		if pollErr != nil {
			return pollErr
		}
	}

	// proceed with using newKey...
	viper.Set("api_key", newKey)
	// write the new key to the config file
	if err = config.WriteConfigFile(); err != nil {
		return err
	}

	fmt.Print("Key created and stored successfully.\n\nYou are now logged in. Happy coding!")

	return nil
}
