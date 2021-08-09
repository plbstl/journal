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

	"github.com/paulebose/diary/internal"
	"github.com/spf13/cobra"
)

// tagRun executes when `create tag` command is run.
func tagRun(cmd *cobra.Command, args []string) {
	tag, err := internal.NewTag()
	if err != nil {
		fmt.Println("cannot create new tag")
		return
	}
	fmt.Printf("tag %s created \n", cyan(tag.Name))
	// @todo: save tag persistently.
}

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a new tag",
	Run:   tagRun,
}

func init() {
	createCmd.AddCommand(tagCmd)
}
