package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func InitConfig(cfgFile string) error {

	if "" == cfgFile {
		return fmt.Errorf("No config file specified (use --help for ... help)\n")
	}

	absCfgFile, err := filepath.Abs(cfgFile)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absCfgFile); err != nil {
		return err
	}

	dir, basename := filepath.Split(absCfgFile)
	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	viper.SetConfigType("json")
	viper.SetConfigName(name)
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func setDefaults() {

}

func validateConfig() {

}
