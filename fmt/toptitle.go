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
	"strings"
)

// TopTitle is takes three strings for the left, center, and right and
// the number of cols wide (in runes, not bytes) that the entire line
// should be returning a string with its heading elements properly
// spaced. The name comes from the position of such a title header line
// at the top of traditional UNIX man pages.
func TopTitle(left, center, right string, cols int) string {
	llen := len([]rune(left))
	clen := len([]rune(center))
	rlen := len([]rune(right))
	if llen+clen+rlen <= cols {
		side := (cols - clen) / 2
		var rep int
		rep = side - llen
		if rep < 0 {
			rep = 0
		}
		lsp := strings.Repeat(" ", rep)
		rep = side - rlen
		if rep < 0 {
			rep = 0
		}
		rsp := strings.Repeat(" ", rep)
		return left + lsp + center + rsp + right
	}
	if clen+rlen <= cols {
		pad := cols - clen - rlen
		sp := strings.Repeat(" ", pad)
		return center + sp + right
	}
	if clen <= cols {
		side := (cols - clen) / 2
		sp := strings.Repeat(" ", side)
		return sp + center + sp
	}
	return center[0:cols]
}
