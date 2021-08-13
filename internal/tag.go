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
package internal

import (
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/pkg/namesgenerator"
)

type Tag struct {
	Name string
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
