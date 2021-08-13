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
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// createRun executes when `create` command is run.
func createRun(cmd *cobra.Command, args []string) {
	var item string
	prompt := &survey.Select{
		Message: "Create a new item:",
		Options: []string{Note, Author, Tag},
	}
	survey.AskOne(prompt, &item)

	switch item {
	case Note:
		noteRun(&cobra.Command{}, []string{})
	case Author:
		authorRun(&cobra.Command{}, []string{})
	case Tag:
		tagRun(&cobra.Command{}, []string{})
	default:
		survey.AskOne(prompt, &item)
	}
}

// createCmd represents the `create` command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note, author or tag",
	Run:   createRun,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
