// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package plugins

import (
	"github.com/spf13/cobra"

	"github.com/deepgram-devs/deepgram-cli/pkg/plugins/interfaces"
)

// Plugin is a Tanzu CLI plugin.
type Plugin struct {
	Cmd *cobra.Command
}

// PluginDescriptor describes a plugin binary.
type PluginDescriptor struct {
	// Name is the name of the plugin.
	Name string `json:"name" yaml:"name"`

	// Description is the plugin's description.
	Description string `json:"description" yaml:"description"`

	// Version of the plugin. Must be a valid semantic version https://semver.org/
	Version string `json:"version" yaml:"version"`

	// EntryPoint for the Cobra CMD and Init function.
	EntryPoint string `json:"entrypoint" yaml:"entrypoint"`

	// BuildSHA is the git commit hash the plugin was built with.
	// TODO(stmcginnis): Update Makefile to set build info with LDFLAG.
	BuildSHA string `json:"buildSHA,omitempty" yaml:"buildSHA,omitempty"`

	// Command group for the plugin.
	Group interfaces.CmdGroup `json:"group,omitempty" yaml:"group,omitempty"`

	// DocURL for the plugin.
	DocURL string `json:"docURL,omitempty" yaml:"docURL,omitempty"`

	// Aliases are other text strings used to call this command
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`
}
