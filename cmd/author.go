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

// authorRun executes when `create author` subcommand is run.
func authorRun(cmd *cobra.Command, args []string) {
	author, err := internal.NewAuthor()
	if err != nil {
		fmt.Println("cannot create new author")
		return
	}

	err = os.MkdirAll(authorsDir, 0755)
	if err != nil {
		fmt.Println("cannot create authors directory", authorsDir)
		return
	}

	exists, err := internal.FileExists("author", author.Name, authorsDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("%s exists already \n", cyan(author.Name))
		return
	}

	if err = os.WriteFile(filepath.Join(authorsDir, author.Name), []byte{}, 0666); err != nil {
		fmt.Println("cannot save new author")
		return
	}

	fmt.Printf("author %s created \n", cyan(author.Name))
}

// authorCmd represents the `author` subcommand.
var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Create a new author",
	Run:   authorRun,
}

func init() {
	createCmd.AddCommand(authorCmd)
}
