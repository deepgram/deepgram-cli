// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package speak

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/deepgram-devs/deepgram-cli/cmd"
	"github.com/deepgram-devs/deepgram-cli/cmd/speak/rest"
)

// speakCmd represents the speak command
var SpeakCmd = &cobra.Command{
	Use:   "speak",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("speak (Text-to-Speech) called")
		fmt.Println("Available subcommands:")
		if len(args) == 0 {
			for _, subCmd := range cmd.Commands() {
				fmt.Printf("- %s: %s\n", subCmd.Use, subCmd.Short)
			}
		}
	},
}

func init() {
	SpeakCmd.AddCommand(rest.TtsRestCmd)

	cmd.RootCmd.AddCommand(SpeakCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// speakCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// speakCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
