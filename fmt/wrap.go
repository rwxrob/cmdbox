/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fmt

import (
	"unicode"
)

// peekWord returns the runes up to the next space.
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
