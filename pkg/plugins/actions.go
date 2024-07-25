// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package plugins

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/spf13/cobra"
)

// plugins that are currently installed
var pluginsInstalled = make(map[string]*Plugin)
var pluginDescriptors = make(map[string]*PluginDescriptor)

// ListInstalledPlugins returns a list of installed plugins
func ListInstalledPlugins() map[string]*Plugin {
	return pluginsInstalled
}

func GetPlugin(pluginName string) *Plugin {
	return pluginsInstalled[pluginName]
}

func DiscoverPlugin(pluginName string) (*PluginDescriptor, error) {
	// Get the plugin binary file
	binaryFilePath := filepath.Join("plugins", pluginName+".so")
	_, err := os.Stat(binaryFilePath)
	if err != nil {
		fmt.Printf("Failed to find plugin binary %s: %v\n", binaryFilePath, err)
		return nil, err
	}

	// Get the plugin descriptor file
	jsonFilePath := filepath.Join("plugins", pluginName+".plugin")
	_, err = os.Stat(jsonFilePath)
	if err != nil {
		fmt.Printf("Failed to find plugin descriptor %s: %v\n", jsonFilePath, err)
		return nil, err
	}

	// Open the ".plugin" file
	fileContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Printf("Failed to read file %s: %v", jsonFilePath, err)
		return nil, err
	}

	// Load the JSON according to the Plugin Descriptor
	var pluginDescriptor PluginDescriptor
	err = json.Unmarshal(fileContent, &pluginDescriptor)
	if err != nil {
		log.Printf("Failed to unmarshal plugin descriptor from file %s: %v", jsonFilePath, err)
		return nil, err
	}

	// load and save the plugin
	newPlugin, err := NewPlugin(&pluginDescriptor)
	if err != nil {
		log.Printf("Failed to load plugin from file %s: %v", jsonFilePath, err)
		return nil, err
	}

	pluginsInstalled[pluginName] = newPlugin
	pluginDescriptors[pluginName] = &pluginDescriptor
	return &pluginDescriptor, nil
}

func LoadPlugin(pluginName, commandName string) (*cobra.Command, error) {
	// Get the plugin binary file
	binaryFilePath := filepath.Join("plugins", pluginName+".so")
	// fmt.Printf("---------> binaryFilePath: %s\n", binaryFilePath)
	_, err := os.Stat(binaryFilePath)
	if err != nil {
		fmt.Printf("Failed to find plugin binary %s: %v\n", binaryFilePath, err)
		return nil, err
	}

	// get plugin command
	p, err := plugin.Open(binaryFilePath)
	if err != nil {
		fmt.Printf("Failed to Open plugin %s: %v\n", binaryFilePath, err)
		return nil, err
	}
	b, err := p.Lookup(commandName + "Cmd")
	if err != nil {
		fmt.Printf("-----> Failed to Lookup plugin %s: %v\n", binaryFilePath, err)
		return nil, err
	}
	f, err := p.Lookup("Init" + commandName)
	if err == nil {
		f.(func())()
	}

	fmt.Printf("Plugin %s installed\n", pluginName)
	return *b.(**cobra.Command), nil
}

func UninstallPlugin(pluginName string) error {
	delete(pluginsInstalled, pluginName)
	delete(pluginDescriptors, pluginName)

	// Remove the plugin file
	jsonFilePath := filepath.Join("plugins", pluginName+".plugin")
	err := os.Remove(jsonFilePath)
	if err != nil {
		log.Printf("Failed to remove plugin file %s: %v", jsonFilePath, err)
		return err
	}

	// Remove the plugin binary
	binaryFilePath := filepath.Join("plugins", pluginName)
	err = os.Remove(binaryFilePath)
	if err != nil {
		log.Printf("Failed to remove plugin %s: %v", binaryFilePath, err)
		return err
	}

	log.Printf("Plugin %s uninstalled", pluginName)
	return nil
}

func DownloadPlugin(pluginName string) error {
	// TODO: Implement this...
	return nil
}
