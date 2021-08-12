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

	"github.com/AlecAivazis/survey/v2"
	"github.com/paulebose/journal/internal"
	"github.com/spf13/cobra"
)

// noteRun executes when `create note` subcommand is run.
func noteRun(cmd *cobra.Command, args []string) {
	author, err := cmd.Flags().GetString("author")
	cobra.CheckErr(err)
	// text, err := cmd.Flags().GetString("text")
	// cobra.CheckErr(err)
	id, err := cmd.Flags().GetString("id")
	cobra.CheckErr(err)
	shouldOpenNote, err := cmd.Flags().GetBool("open")
	cobra.CheckErr(err)

	if author == "" {
		// @todo: fetch default author from somewhere.
		const defaultAuthor = "Authy Prof"
		author = defaultAuthor
	}
	note := internal.NewNote(id, author)
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

	err = os.MkdirAll(notesDir, 0755)
	if err != nil {
		fmt.Println("cannot create notes directory", notesDir)
		return
	}

	exists, err := internal.FileExists("note", note.ID, notesDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Printf("note with id %s exists already \n", cyan(note.ID))
		return
	}

	if err = os.WriteFile(filepath.Join(notesDir, note.ID), []byte{}, 0666); err != nil {
		fmt.Println("cannot save new note")
		return
	}
}

// noteCmd represents the `note` subcommand
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create a new note",
	Run:   noteRun,
}

func init() {
	createCmd.AddCommand(noteCmd)
	initNoteFlags(noteCmd)
}

func initNoteFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("open", "o", false, "Open note after creating it")
	cmd.Flags().StringP("text", "t", "", "Add initial text to created note")
	cmd.Flags().String("id", "", "Associate a custom ID")
	cmd.Flags().String("author", "", "Use custom author to create note")
}
