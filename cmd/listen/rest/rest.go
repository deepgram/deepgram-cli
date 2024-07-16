// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package rest

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sttRestCmd represents the sttRest command
var SttRestCmd = &cobra.Command{
	Use:     "rest",
	Aliases: []string{"sttrest", "stt-rest", "speech-to-text-rest"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Speech-to-Text Rest called")
	},
}
