// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/deepgram-devs/deepgram-cli/pkg/plugins"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Get the list of files in the "plugins" subdirectory
	files, err := os.ReadDir("plugins")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the files and create a new plugin for each one
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".plugin") {
			// remove the ".plugin" suffix
			pluginName := strings.TrimSuffix(file.Name(), ".plugin")

			desc, err := plugins.DiscoverPlugin(pluginName)
			if err != nil {
				log.Printf("Failed to install plugin %s: %v", pluginName, err)
				continue
			}

			pluginCmd, err := plugins.LoadPlugin(pluginName, desc.EntryPoint)
			if err != nil {
				log.Printf("Failed to install plugin %s: %v", pluginName, err)
				continue
			}

			// add command
			RootCmd.AddCommand(pluginCmd)
		}
	}

	// load all plugins
	pluginsInstalled := plugins.ListInstalledPlugins()
	for _, plugin := range pluginsInstalled {
		// fmt.Printf("-----------------> Adding plugin: %s\n", name)
		RootCmd.AddCommand(plugin.Cmd)
	}

	// TODO: disable completion command?
	RootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
