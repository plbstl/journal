/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paulebose/diary/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "diary",
	Short: "Keep important notes in your cli",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Version = internal.Version
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default on windows=[%AppData%\\Diary\\config.yml], mac=[$HOME/Library/Application Support/Diary/config.yml], linux=[$HOME/.config/Diary/config.yml]).")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	userConfigDir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	diaryConfigDir = filepath.Join(userConfigDir, "Diary")
	authorsDir = filepath.Join(diaryConfigDir, "authors")
	notesDir = filepath.Join(diaryConfigDir, "notes")
	tagsDir = filepath.Join(diaryConfigDir, "tags")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in diaryConfigDir directory with name "config" (without extension).
		viper.AddConfigPath(diaryConfigDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
