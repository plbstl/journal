package internal

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var (
	// The "safe" set of characters.
	alphanum = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005A, 1}, // A-Z
			{0x0061, 0x007A, 1}, // a-z
		},
	}
	// Characters in these ranges will be ignored.
	nop = []*unicode.RangeTable{
		unicode.Mark,
		unicode.Sk, // Symbol - modifier
		unicode.Lm, // Letter - modifier
		unicode.Cc, // Other - control
		unicode.Cf, // Other - format
	}
)

// slug replaces each run of characters which are not ASCII letters or numbers
// with the Replacement character, except for leading or trailing runs. Letters
// will be stripped of diacritical marks and lowercased. Letter or number
// codepoints that do not have combining marks or a lower-cased variant will be
// passed through unaltered.
func slug(s string) string {
	buf := make([]rune, 0, len(s))
	replacement := false

	for _, r := range norm.NFKD.String(s) {
		switch {
		case unicode.In(r, alphanum):
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
