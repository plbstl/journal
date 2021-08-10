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
	"github.com/spf13/cobra"
)

var red = color.New(color.FgRed).SprintFunc()

// deleteRun executes when `delete` command is run.
func deleteRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		color.Red("An ID must be passed in! \n\n")
		cmd.Usage()
		return
	}

	id := args[0]
	lists := []struct {
		label string
		dir   string
	}{
		{Note, notesDir},
		{Author, authorsDir},
		{Tag, tagsDir},
	}
	for _, list := range lists {
		files, err := listOf(list.dir)
		if err != nil {
			fmt.Printf("cannot read %ss directory: %s \n", list.label, red(list.dir))
			return
		}
		for _, ID := range files {
			if id == ID {
				if err = os.Remove(filepath.Join(list.dir, id)); err != nil {
					fmt.Println("error removing item with ID", id)
				}
				fmt.Printf("deleted %s %s \n", list.label, cyan(id))
				return
			}
		}
	}
	fmt.Printf("cannot find ID %s \n", cyan(id))
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a note, tag or author with the specified ID",
	Aliases: []string{"del"},
	Run:     deleteRun,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// deleteFiles removes an item (notes, authors or tags)
// by deleting all the contents of its folder.
func deleteFiles(item, dir string) {
	files, err := listOf(dir)
	if err != nil {
		fmt.Printf("cannot read %s directory: %s \n", item, red(dir))
		return
	}

	if err := os.RemoveAll(dir); err != nil {
		fmt.Printf("cannot complete deletion of %s \n", cyan(item))
		return
	}

	for _, file := range files {
		color.Red(file)
	}
	fmt.Printf("%d %s deleted \n", len(files), item)
}
