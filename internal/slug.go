package internal

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Characters in these ranges will be ignored.
var nop = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk, // Symbol - modifier
	unicode.Lm, // Letter - modifier
	unicode.Cc, // Other - control
	unicode.Cf, // Other - format
}

// slug transforms:
// Æther -> æther
// Mørdag -> mørdag
// Être en été est à mi-chemin de noël -> etre-en-ete-est-a-mi-chemin-de-noel

// slug replaces each run of characters which are not ASCII letters with
// the '-' character, except for leading or trailing runs. Letters will
// be stripped of diacritical marks and lowercased. Letter codepoints
// that do not have combining marks or a lower-cased variant will be
// passed through unaltered.
func slug(s string) string {
	buf := make([]rune, 0, len(s))
	replacement := false

	for _, r := range norm.NFKD.String(s) {
		switch {
		case unicode.In(r, unicode.Letter):
			buf = append(buf, unicode.ToLower(r))
			replacement = true
		case unicode.IsOneOf(nop, r):
			// skip
		case replacement:
			buf = append(buf, '-')
			replacement = false
		}
	}

	// Strip trailing '-' byte
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}

	return string(buf)
}
