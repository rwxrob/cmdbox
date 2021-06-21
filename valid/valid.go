/* The valid package simply and efficiently minimally validates command
 * line components passed into CmdBox commands, aliases, and such.
 */
package valid

import (
	"strings"
	"unicode"
)

// Name validates a name string. Names must consist of one or two words
// separated by a single space and composed of Unicode lowercase letters
// and should be full, speakable words in order to promote command lines
// that can be directly called from conversational UIs and chat
// interfaces without modification.
func Name(name string) bool {
	if name == "" {
		return false
	}
	words := strings.Split(name, " ")
	for _, word := range words {
		//	fmt.Printf("%q\n", word)
		if !IsWord(word) {
			return false
		}
	}
	return true
}

// IsWord returns true if ever rune is a lowercase letter.
func IsWord(w string) bool {
	if len(w) == 0 {
		return false
	}
	for _, r := range w {
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
