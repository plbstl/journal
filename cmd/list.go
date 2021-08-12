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

	"github.com/fatih/color"
	"github.com/paulebose/journal/internal"
	"github.com/spf13/cobra"
)

// listRun executes when `list` command is run.
func listRun(cmd *cobra.Command, args []string) {
	// authors, err := cmd.Flags().GetBool("authors")
	// cobra.CheckErr(err)
	// tags, err := cmd.Flags().GetBool("tags")
	// cobra.CheckErr(err)
	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)

	// fetch notes
	files, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Printf("cannot read directory: %s \n", notesDir)
	}
	filepaths := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			filepaths = append(filepaths, filepath.Join(notesDir, file.Name()))
		}
	}

	// list notes
	for i, fp := range filepaths {
		if i >= limit {
			break
		}

		var note internal.Note
		if err = internal.UnmarshalNote(&note, fp); err != nil {
			fmt.Println(err)
			continue
		}

		if note.Title == "" {
			color.Blue("[No Title]")
		} else {
			color.Blue(note.Title)
		}
		fmt.Println(note.TruncatedBody(69))
		fmt.Println()
	}
}

// listCmd represents the `list` command.
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Print out a list of created notes",
	Aliases: []string{"ls"},
	Run:     listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("authors", "a", false, "Print out a list of authors")
	listCmd.Flags().BoolP("tags", "t", false, "Print out a list of created tags")
	listCmd.Flags().IntP("limit", "l", 10, "Max number of items to print out")
}

// listOf returns a slice of filenames found in the directory.
func listOf(dir string) ([]string, error) {
	var filenames []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}
