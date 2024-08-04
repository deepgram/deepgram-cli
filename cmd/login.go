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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in a user",
	Long:  `Logs a user into Deepgram with browser-based authentication.`,
	Run:   login,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().StringP("dg-key", "k", "", "Configure the CLI with your Deepgram API key")
	// viper.BindPFlag("dg-key", loginCmd.Flags().Lookup("dg-key"))
}

func login(cmd *cobra.Command, args []string) {
	var (
		dgKey string = viper.GetString("dg-key")

		str string
		err error
	)

	cmd.Context()

	switch {
	case dgKey != "":
		str, err = configureAuth(cmd, dgKey)
	default:
		str, err = webAuth(cmd)
	}

	fmt.Print(str, err)
}

func configureAuth(cmd *cobra.Command, dgKey string) (string, error) {
	var (
		err error
	)

	fmt.Print(cmd, dgKey)

	if err != nil {
		return "", err
	}

	return dgKey, nil
}

func webAuth(cmd *cobra.Command) (string, error) {
	var (
		url   string = "https://community.deepgram.com/api/auth/cli?"
		dgKey string = "12345"
		err   error
	)

	fmt.Print(url, cmd, dgKey)

	if err != nil {
		return "", err
	}

	return configureAuth(cmd, dgKey)
}
