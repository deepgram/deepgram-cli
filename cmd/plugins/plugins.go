// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package plugins

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/deepgram-devs/deepgram-cli/cmd"
	"github.com/deepgram-devs/deepgram-cli/pkg/plugins"
)

// pluginCmd represents the manage command
var pluginCmd = &cobra.Command{
	Use:   "plugins",
	Short: "This manages plugins in this installation.",
	Long: `This manages plugins current instaled. For example:

list - lists all plugins
install -name <PLUGIN NAME> - installs a plugin
uninstall -name <PLUGIN NAME> - uninstalls a plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plugins called")
		if len(args) == 0 {
			pluginsInstalled := plugins.ListInstalledPlugins()
			for name := range pluginsInstalled {
				fmt.Printf("Plugin: %s\n", name)
			}
		}
	},
}

func init() {
	// // load all plugins
	// pluginsInstalled := plugins.ListInstalledPlugins()
	// for name, plugin := range pluginsInstalled {
	// 	fmt.Printf("-----------------> Adding plugin: %s\n", name)
	// 	cmd.RootCmd.AddCommand(plugin.Cmd)
	// }

	// add the plugin command
	cmd.RootCmd.AddCommand(pluginCmd)

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	pluginCmd.PersistentFlags().String("name", "", "Plugin name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
