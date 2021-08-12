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
