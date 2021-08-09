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

	"github.com/AlecAivazis/survey/v2"
	"github.com/paulebose/diary/internal"
	"github.com/spf13/cobra"
)

var (
	shouldOpenNote bool
	text           string
)

// noteRun executes when `create note` command is run.
func noteRun(cmd *cobra.Command, args []string) {
	// @todo: fetch author from somewhere.
	au := &internal.Author{
		ID:   "test-123",
		Name: "authy",
	}
	note := internal.NewNote(id, au)
	fmt.Printf("note created with id %s \n", cyan(note.ID))

	if shouldOpenNote {
		internal.OpenNote(&note)
	} else {
		prompt := &survey.Confirm{
			Message: "Open created note?",
		}
		survey.AskOne(prompt, &shouldOpenNote)
		if shouldOpenNote {
			internal.OpenNote(&note)
		}
	}
	// @todo: save note persistently.
}

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create a new note",
	Run:   noteRun,
}

func init() {
	createCmd.AddCommand(noteCmd)
	noteCmd.Flags().BoolVarP(&shouldOpenNote, "open", "o", false, "Open note after creating it")
	noteCmd.Flags().StringVarP(&text, "text", "t", "", "Add initial text to created note")
}
