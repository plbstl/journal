package internal

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID, AuthorName, Text string
}

// NewNote creates and returns a new note.
// It returns an error if any.
func NewNote(id, authorName string) Note {
	var note Note
	note.AuthorName = authorName
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
