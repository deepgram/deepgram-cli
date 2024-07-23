// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package plugins

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/deepgram-devs/deepgram-cli/pkg/plugins/interfaces"
)

// NewPlugin creates an instance of Plugin.
func NewPlugin(descriptor *PluginDescriptor) (*Plugin, error) {
	p := &Plugin{
		Cmd: &cobra.Command{
			Use:     descriptor.Name,
			Short:   descriptor.Description,
			Aliases: descriptor.Aliases,
		},
	}

	p.Cmd.AddCommand(
		newDescribeCmd(descriptor.Description),
		newVersionCmd(descriptor.Version),
		newInfoCmd(descriptor),
	)

	return p, nil
}

// NewTestFor creates a plugin descriptor for a test plugin.
func NewTestFor(pluginName string) *PluginDescriptor {
	return &PluginDescriptor{
		Name:        fmt.Sprintf("%s-test", pluginName),
		Description: fmt.Sprintf("test for %s", pluginName),
		Version:     "v0.0.1",
		// BuildSHA:    SHA,
		Group:   interfaces.TestCmdGroup,
		Aliases: []string{fmt.Sprintf("%s-alias", pluginName)},
	}
}

func newDescribeCmd(description string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "describe",
		Short:  "Describes the plugin",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(description)
			return nil
		},
	}

	return cmd
}

func newVersionCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Short:  "Version the plugin",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(version)
			return nil
		},
	}

	return cmd
}

func newInfoCmd(desc *PluginDescriptor) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "info",
		Short:  "Plugin info",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := json.Marshal(desc)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
			return nil
		},
	}

	return cmd
}

// AddCommands adds commands to the plugin.
func (p Plugin) AddCommands(commands ...*cobra.Command) error {
	p.Cmd.AddCommand(commands...)
	return nil
}

// Execute executes the plugin.
func (p Plugin) Execute() error {
	return p.Cmd.Execute()
}
