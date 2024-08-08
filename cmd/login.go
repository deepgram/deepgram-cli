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

		str string
		err error
	)

	switch {
	case key != "":
		err = cliAuth()
	default:
		err = webAuth()
	}

	fmt.Print(str, err)
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

func webAuth() error {
	var (
		key string = viper.GetString("api_key")

		newKey string
		err    error
	)

	if key != "" {
		cobra.CheckErr(common.PromptBool("You're already logged in. Do you want to login again?"))
	}

	cobra.CheckErr(err)

	ppid := os.Getppid()
	hostname, err := os.Hostname()

	cobra.CheckErr(err)

	auth, err := auth.RequestDeviceCode(ppid, hostname, []string{"usage:write"})
	if err != nil {
		return err
	}

	fmt.Println(auth)

	if err := open.Run(auth.VerificationURI); err != nil {
		cobra.CheckErr(err)
	}

	// start session
	// open browser
	// wait for response

	fmt.Println(auth)

	viper.Set("api_key", newKey)
	newKey = "123456"

	return config.WriteConfigFile()
}

// func openbrowser(url string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		cobra.CheckErr(err)
// 	}

// }
