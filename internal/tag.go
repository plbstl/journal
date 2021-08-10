package internal

import (
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/pkg/namesgenerator"
)

type Tag struct {
	Name string `json:"name"`
}

// NewTag creates and returns a new tag.
// It returns an error if any.
func NewTag() (Tag, error) {
	var tag Tag
	prompt := &survey.Input{Message: "Enter Tag name:"}
	if err := survey.AskOne(prompt, &tag.Name); err != nil {
		return tag, err
	}
	tag.Name = slug(tag.Name)
	if tag.Name == "" {
		rand.Seed(time.Now().UnixNano())
		tag.Name = slug(namesgenerator.GetRandomName(0))
	}
	return tag, nil
}
