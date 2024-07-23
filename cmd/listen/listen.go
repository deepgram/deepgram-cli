// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package listen

import (
	"github.com/spf13/cobra"

	"github.com/deepgram-devs/deepgram-cli/cmd"
	rest "github.com/deepgram-devs/deepgram-cli/cmd/listen/rest"
)

// listenCmd represents the listen command
var ListenCmd = &cobra.Command{
	Use:   "listen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	ListenCmd.AddCommand(rest.SttRestCmd)

	cmd.RootCmd.AddCommand(ListenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
