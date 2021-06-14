package valid

import (
	"unicode"
)

// Name validates a name string. Names must consist of Unicode lowercase
// letters and should be full, speakable words in order to promote
// command lines that can be directly called from conversational UIs and
// chat interfaces without modification.
func Name(name string) bool {
	if name == "" {
		return false
	}
	for _, r := range name {
		if !(unicode.IsLetter(r) && unicode.IsLower(r)) {
			return false
		}
	}
	return true
}

// Alias validates an alias string. Aliases must consist of Unicode
// letters only but do not need to be full, speakable workds. For
// performance reason, only the first rune is checked.
func Alias(alias string) bool { return Name(alias) }
