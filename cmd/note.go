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
