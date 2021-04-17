package fmt

import (
	"unicode"
)

func peekWord(buf []rune, start int) []rune {
	word := []rune{}
	for _, r := range buf[start:] {
		if unicode.IsSpace(r) {
			break
		}
		word = append(word, r)
	}
	return word
}

// Wrap wraps the string to the given width using spaces to separate
// words. If passed a negative width will effectively join all words in
// the buffer into a single line with no wrapping.
func Wrap(buf string, width int) string {
	if width == 0 {
		return buf
	}
	nbuf := ""
	curwidth := 0
	for i, r := range []rune(buf) {
		// hard breaks always as is
		if r == '\n' {
			nbuf += "\n"
			curwidth = 0
			continue
		}
		if unicode.IsSpace(r) {
			// FIXME: don't peek every word, only after passed width
			// change the space to a '\n' in the buffer slice directly
			next := peekWord([]rune(buf), i+1)
			if width > 0 && (curwidth+len(next)+1) > width {
				nbuf += "\n"
				curwidth = 0
				continue
			}
		}
		nbuf += string(r)
		curwidth++
	}
	return nbuf
}
