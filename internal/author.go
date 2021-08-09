package internal

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
)

type Author struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Notes []*Note `json:"notes"`
}

func NewAuthor(id string) (Author, error) {
	var author Author
	prompt := &survey.Input{Message: "Enter Author name:"}
	if err := survey.AskOne(prompt, &author.Name); err != nil {
		return author, err
	}
	if author.Name == "" {
		// @todo: set random name for author if empty.
		author.Name = "random name"
	}
	if id == "" {
		id = uuid.NewString()
	}
	author.ID = id
	return author, nil
}
