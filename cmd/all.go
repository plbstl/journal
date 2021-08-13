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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// allRun executes when `delete all` subcommand is run.
func allRun(cmd *cobra.Command, args []string) {
	notes, err := cmd.Flags().GetBool("notes")
	cobra.CheckErr(err)
	authors, err := cmd.Flags().GetBool("authors")
	cobra.CheckErr(err)
	tags, err := cmd.Flags().GetBool("tags")
	cobra.CheckErr(err)

	var used bool
	if notes {
		deleteFiles("notes", notesDir)
		used = true
	}
	if authors {
		deleteFiles("authors", authorsDir)
		used = true
	}
	if tags {
		deleteFiles("tags", tagsDir)
		used = true
	}
	if !used {
		color.Red("A flag must be used! \n\n")
		cmd.Usage()
	}
}

// allCmd represents the `all` subcommand
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Delete all notes, authors or tags",
	Long:  `Delete all notes, authors or tags. Deletions do not cascade.`,
	Run:   allRun,
}

func init() {
	deleteCmd.AddCommand(allCmd)
	allCmd.Flags().Bool("notes", false, "Delete all notes")
	allCmd.Flags().Bool("authors", false, "Delete all authors")
	allCmd.Flags().Bool("tags", false, "Delete all tags")
}
