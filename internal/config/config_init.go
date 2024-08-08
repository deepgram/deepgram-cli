package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"deepgram-cli/internal/common"
)

func ConfigInit(cmd *cobra.Command, args []string) error {
	viper.SetConfigName(common.ConfigFileName)
	viper.SetConfigType(common.ConfigFileType)

	if dirname, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(filepath.Join(dirname, common.ConfigHomeSubdir))
		viper.AddConfigPath(filepath.Join(dirname, common.ConfigHomeSubdir, common.ConfigDeepgramSubdir))
	}

	viper.SetEnvPrefix("deepgram")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return nil
}
