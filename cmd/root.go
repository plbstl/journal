/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paulebose/journal/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "journal",
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default on windows=[%AppData%\\Journal\\config.yml], mac=[$HOME/Library/Application Support/Journal/config.yml], linux=[$HOME/.config/Journal/config.yml]).")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	userConfigDir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	journalConfigDir = filepath.Join(userConfigDir, "Journal")
	authorsDir = filepath.Join(journalConfigDir, "authors")
	notesDir = filepath.Join(journalConfigDir, "notes")
	tagsDir = filepath.Join(journalConfigDir, "tags")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in journalConfigDir directory with name "config" (without extension).
		viper.AddConfigPath(journalConfigDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
