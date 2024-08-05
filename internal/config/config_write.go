package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"deepgram-cli/pkg/common"
)

func ConfigFileExists() bool {
	return viper.ConfigFileUsed() != ""
}

func WriteConfigFile() error {
	configFile := viper.ConfigFileUsed()

	if configFile == "" {
		fullname := common.ConfigFileName + "." + common.ConfigFileType

		if dirname, err := os.UserHomeDir(); err == nil {
			configFile = filepath.Join(dirname, common.ConfigHomeSubdir, common.ConfigDeepgramSubdir, fullname)
		}

		if configFile == "" {
			return errors.New("Failed to acquire config directory name")
		}

		configDirPath := filepath.Dir(configFile)

		if err := os.MkdirAll(configDirPath, os.ModePerm); err != nil {
			return err
		}

		fmt.Printf("Deepgram configuration file created at %s\n", configFile)
	}

	if err := viper.WriteConfigAs(configFile); err != nil {
		return err
	}

	return nil
}
