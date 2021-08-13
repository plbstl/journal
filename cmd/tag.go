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
)

// tagRun executes when `create tag` subcommand is run.
func tagRun(cmd *cobra.Command, args []string) {
	tag, err := internal.NewTag()
	if err != nil {
		fmt.Println("cannot create new tag")
		return
	}

	err = os.MkdirAll(tagsDir, 0755)
	if err != nil {
		fmt.Println("cannot create tags directory", tagsDir)
		return
	}

	exists, err := internal.FileExists("tag", tag.Name, tagsDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("%s exists already \n", cyan(tag.Name))
		return
	}

	if err = os.WriteFile(filepath.Join(tagsDir, tag.Name), []byte{}, 0666); err != nil {
		fmt.Println("cannot save new tag")
		return
	}

	fmt.Printf("tag %s created \n", cyan(tag.Name))
}

// tagCmd represents the `tag` subcommand
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a new tag",
	Run:   tagRun,
}

func init() {
	createCmd.AddCommand(tagCmd)
}
