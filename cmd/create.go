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
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// createRun executes when `create` command is run.
func createRun(cmd *cobra.Command, args []string) {
	var item string
	prompt := &survey.Select{
		Message: "Create a new item:",
		Options: []string{Note, Author, Tag},
	}
	survey.AskOne(prompt, &item)

	switch item {
	case Note:
		noteRun(&cobra.Command{}, []string{})
	case Author:
		authorRun(&cobra.Command{}, []string{})
	case Tag:
		tagRun(&cobra.Command{}, []string{})
	default:
		survey.AskOne(prompt, &item)
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note, author or tag",
	Run:   createRun,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
