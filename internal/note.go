package internal

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID     string  `json:"id"`
	Author *Author `json:"author"`
	Text   string  `json:"text"`
}

func NewNote(id string, au *Author) Note {
	var note Note
	note.Author = au
	if id == "" {
		id = uuid.NewString()
	}
	note.ID = id
	return note
}

func OpenNote(note *Note) {
	// @todo: use a better way to load editor.
	p := &survey.Editor{FileName: "*.md"}
	survey.AskOne(p, "")
}
