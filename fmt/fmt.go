package fmt

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

// Emph takes a command documentation format string (an extremely limited
// version of Markdown that is also Godoc friendly) and transforms it as
// follows:
//
// * Initial and trailing blank lines are removed.
//
// * Indentation is removed - the number of spaces preceeding the first word of
//   the first line are ignored in every line (including raw text blocks).
//
// * Raw text ignored - any line beginning with four or more spaces (after
//   convenience indentation is removed) will be kept as it is exactly (code
//   examples, etc.) but should never exceed 80 characters (including the
//   spaces).
//
// * Blocks are unwrapped - any non-blank (without three or less initial spaces)
//   will be trimmed line following a line will be joined to the preceding line
//   recursively (unless hard break).
//
// * Hard breaks kept - like Markdown any line that ends with two or more
//   spaces will automatically force a line return.
//
// * URL links argument names and anything else within angle brackets (<url>),
//   will trigger italics in both text blocks and usage sections.
//
// * Italic, Bold, and BoldItalic inline emphasis using one, two, or three
//   stars respectivly will be observed and cannot be intermixed or intraword.
//   Each opener must be preceded by a UNICODE space (or nothing) and followed by
//   a non-space rune. Each closer must be preceded by a non-space rune and
//   followed by a UNICODE space (or nothing).
//
// For historic reasons the following environment variables will be
// observed if found (and also provide color support for the less
// pager utility):
//
// * italic - LESS_TERMCAP_so
// * bold - LESS_TERMCAP_md
// * bolditalic = LESS_TERMCAP_mb
//
func Emph(input string, indent, width int) (output string) {

	// this scanner could be waaaay more lexy
	// but suits the need and clear to read

	var strip int
	var blockbuf string

	// standard state machine approach
	inblock := false
	inraw := false
	inhard := false
	gotindent := false

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		txt := scanner.Text()
		trimmed := strings.TrimSpace(txt)

		// ignore blank lines
		if !(inraw || inblock) && len(trimmed) == 0 {
			continue
		}

		// infer the indent to strip for every line
		if !gotindent && len(trimmed) > 0 {
			for i, v := range txt {
				if v != ' ' {
					strip = i
					gotindent = true
					break
				}
			}
		}

		// strip convenience indent
		if len(txt) >= strip {
			txt = txt[strip:]
		}

		// raw block start
		if !inblock && !inraw && len(txt) > 4 && txt[0:4] == "    " && len(trimmed) > 0 {
			inraw = true
			output += "\n\n" + txt
			continue
		}

		// in raw block
		if inraw && len(txt) > 4 {
			output += "\n" + txt
			continue
		}

		// raw block end
		if inraw && len(trimmed) == 0 {
			inraw = false
			continue
		}

		// another block line, join it
		if inblock && len(trimmed) > 0 {
			if len(txt) >= 2 && txt[len(txt)-2:] == "  " {
				inhard = true
			}
			space := " "
			if inhard {
				space = "\n"
			}
			blockbuf += space + trimmed
			continue
		}

		// beginning of a new block
		if !inblock && len(trimmed) > 0 {
			inhard = false
			inblock = true
			if len(txt) >= 2 && txt[len(txt)-2:] == "  " {
				inhard = true
			}
			blockbuf = trimmed
			continue
		}

		// end block
		if inblock && len(trimmed) == 0 {
			inblock = false
			output += "\n\n" + Emphasize(Wrap(blockbuf, width-strip-4))
			continue
		}
	}

	// flush last block
	if inblock {
		output += "\n\n" + Emphasize(Wrap(blockbuf, width-strip-4))
	}
	output = Indent(strings.TrimSpace(output), indent)
	return
}

func Emphasize(buf string) string {

	// keep for debugging
	// italic = `<italic>`
	// bold = `<bold>`
	// bolditalic = `<bolditalic>`
	// reset = `<reset>`
	// underline = `<ul>`

	nbuf := []rune{}
	prev := ' '
	opentok := false
	otok := ""
	closetok := false
	ctok := ""
	for i := 0; i < len([]rune(buf)); i++ {
		r := []rune(buf)[i]

		if r == '<' {
			// TODO detect underline support
			//nbuf = append(nbuf, '<')
			nbuf = append(nbuf, []rune(underline)...)
			for {
				i++
				r = unicode.ToUpper(rune(buf[i]))
				if r == '>' {
					i++
					break
				}
				nbuf = append(nbuf, r)
			}
			nbuf = append(nbuf, []rune(reset)...)
			//nbuf = append(nbuf, '>')
			i--
			continue
		}

		if r != '*' {

			if opentok {
				tokval := " "
				if !unicode.IsSpace(r) {
					switch otok {
					case "*":
						tokval = italic
					case "**":
						tokval = bold
					case "***":
						tokval = bolditalic
					}
				} else {
					tokval = otok
				}
				nbuf = append(nbuf, []rune(tokval)...)
				opentok = false
				otok = ""
			}

			if closetok {
				nbuf = append(nbuf, []rune(reset)...) // practical, not perfect
				ctok = ""
				closetok = false
			}

			prev = r
			nbuf = append(nbuf, r)
			continue
		}

		// everything else for '*'
		if unicode.IsSpace(prev) || opentok {
			opentok = true
			otok += string(r)
			continue
		}

		// only closer conditions remain
		if !unicode.IsSpace(prev) {
			closetok = true
			ctok += string(r)
			continue
		}

		// nothing special
		closetok = false
		nbuf = append(nbuf, r)
	}

	// for tokens at the end of a block
	if closetok {
		nbuf = append(nbuf, []rune(reset)...)
	}

	return string(nbuf)
}

// Indent indents each line the set number of spaces.
func Indent(buf string, spaces int) string {
	nbuf := ""
	scanner := bufio.NewScanner(strings.NewReader(buf))
	scanner.Scan()
	for n := 0; n < spaces; n++ {
		nbuf += " "
	}
	nbuf += scanner.Text()
	for scanner.Scan() {
		nbuf += "\n"
		for n := 0; n < spaces; n++ {
			nbuf += " "
		}
		nbuf += scanner.Text()
	}
	return nbuf
}

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
		lsp := strings.Repeat(" ", side-llen)
		rsp := strings.Repeat(" ", side-rlen)
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

var reset = "\033[0m"
var italic = reset + "\033[3m"
var bold = reset + "\033[1m"
var bolditalic = reset + "\033[1m\033[3m"
var underline = reset + "\033[4m"

func init() {
	//us := os.Getenv("LESS_TERMCAP_us")
	us := os.Getenv("LESS_TERMCAP_us")
	md := os.Getenv("LESS_TERMCAP_md")
	mb := os.Getenv("LESS_TERMCAP_mb")
	if us != "" {
		italic = reset + us
	}
	if md != "" {
		bold = reset + md
	}
	if mb != "" {
		bolditalic = reset + mb
	}
}
