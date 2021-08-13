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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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

// deleteCmd represents the `delete` command
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
