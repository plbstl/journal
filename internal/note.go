package internal

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID     string
	Author *Author
	Text   string
}

// NewNote creates and returns a new note.
// It returns an error if any.
func NewNote(id string, au *Author) Note {
	var note Note
	note.Author = au
	if id == "" {
		id = uuid.NewString()
	}
	note.ID = id
	return note
}

// OpenNote opens a note in a special terminal editor.
func OpenNote(note *Note) {
	// @todo: use a better way to load editor.
	p := &survey.Editor{FileName: "*.md"}
	survey.AskOne(p, "")
}
