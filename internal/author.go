package internal

import (
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/pkg/namesgenerator"
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

	author.Name = normalize(author.Name)
	if author.Name == "" {
		rand.Seed(time.Now().UnixNano())
		au := strings.Split(namesgenerator.GetRandomName(0), "_")
		// capitalize first letters.
		au[0] = strings.ToUpper(string(au[0][0])) + au[0][1:]
		au[1] = strings.ToUpper(string(au[1][0])) + au[1][1:]
		author.Name = strings.Join(au, " ")
	}

	if id == "" {
		id = uuid.NewString()
	}
	author.ID = id
	return author, nil
}

// normalize returns the string with extra spaces
// and non-alphanumeric characters removed.
func normalize(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	buf := make([]rune, 0, len(s))
	addSpace := false

	// must be alpha numeric.
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buf = append(buf, r)
			addSpace = true
			continue
		}
		if addSpace {
			buf = append(buf, ' ')
			addSpace = false
		}
	}

	s = strings.TrimSpace(string(buf))
	return s
}
