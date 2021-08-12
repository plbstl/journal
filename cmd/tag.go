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
