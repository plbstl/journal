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
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// nop contains character ranges to be ignored.
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
	dash := false

	for _, r := range norm.NFKD.String(s) {
		switch {
		case unicode.In(r, unicode.Letter):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(nop, r):
			// skip
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}

	// Strip trailing '-' byte
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}

	return string(buf)
}
