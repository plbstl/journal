package internal

import (
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID, AuthorName, Title, Body string
}

// TruncatedBody truncates the note's body
// to a given length and returns it.
func (n *Note) TruncatedBody(length int) string {
	if len(n.Body) < length {
		return strings.TrimSpace(n.Body[:])
	}
	return strings.TrimSpace(n.Body[:length]) + "..."
}

// NewNote creates and returns a new note.
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

// UnmarshalNote parses the data in the specified file path
// and stores the result in the Note pointed to by n.
func UnmarshalNote(n *Note, fp string) error {
	// @todo: validate marshalled note.
	lines, err := parseLines(fp)
	if err != nil {
		return err
	}
	n.ID = filepath.Base(fp)
	n.AuthorName = lines[0]
	n.Title = lines[1]
	if len(lines) == 3 {
		n.Body = lines[2]
	}
	return nil
}
