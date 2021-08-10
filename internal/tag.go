package internal

import (
	"github.com/AlecAivazis/survey/v2"
)

type Tag struct {
	Name string `json:"name"`
}

func NewTag() (Tag, error) {
	var tag Tag
	prompt := &survey.Input{Message: "Enter Tag name:"}
	if err := survey.AskOne(prompt, &tag.Name); err != nil {
		return tag, err
	}
	if tag.Name == "" {
		// @todo: set random name for tag if empty.
		tag.Name = "tag-name1"
	}
	tag.Name = slug(tag.Name)
	return tag, nil
}
