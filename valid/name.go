package valid

import (
	"unicode"
	"unicode/utf8"
)

// Name validates a name string. Names must consist of Unicode letters
// only and need to be full, speakable words. For performance reasons,
// only the first rune in the string is checked.
func Name(name string) bool {
	if name == "" {
		return false
	}
	r, _ := utf8.DecodeRuneInString(name)
	return unicode.IsLetter(r)
}

// Alias validates an alias string. Aliases must consist of Unicode
// letters only but do not need to be full, speakable workds. For
// performance reason, only the first rune is checked.
func Alias(alias string) bool { return Name(alias) }
