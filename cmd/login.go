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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"deepgram-cli/internal/config"
	"deepgram-cli/pkg/common"
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
	loginCmd.Flags().StringVarP(&ApiKey, "api_key", "k", "", "Configure the CLI with your Deepgram API key")
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api_key"))

	// allow users to force write and skip the prompt
	loginCmd.Flags().BoolP("force-write", "f", false, "Don't prompt for confirmation when providing an API key")
}

func runLogin(cmd *cobra.Command, args []string) {
	var (
		// Get the API key from the config context
		dgKey string = viper.GetString("api_key")

		str string
		err error
	)

	fmt.Println("dgKey", dgKey)

	switch {
	case dgKey != "":
		err = cliAuth()
	default:
		str, err = webAuth(cmd)
	}

	fmt.Print(str, err)
}

func cliAuth() error {
	var err error

	if !viper.GetBool("force-write") {
		err = common.PromptBool("Do you want to write this key to config?")

		if config.ConfigFileExists() {
			err = common.PromptBool("Configuration file already exists. Overwrite?")
		}
	}

	if err != nil {
		return err
	}

	return config.WriteConfigFile()
}

func confirmAction() {
	if viper.GetBool("force-write") {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to write this key to config? (y/N): ")

	var err error

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character and convert to lower case
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if input != "y" && input != "yes" {
		os.Exit(0)
	}

	fmt.Println("Writing key to config")
}

func webAuth(cmd *cobra.Command) (string, error) {
	var (
		dgKey string = "12345"
		err   error
	)

	fmt.Print(cmd, dgKey)

	if err != nil {
		return "", err
	}

	return "Success", nil
}
