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
