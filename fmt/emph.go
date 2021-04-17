package fmt

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/rwxrob/cmdbox/term/esc"
)

// NOTE: These are consistent with esc package but are repeated here to
// enable them to be used through the fmt package in short form. They
// are also not constants because they can be overriden by environment
// variables (per ancient termcap conventions).

var reset = esc.Reset
var italic = esc.Italic
var bold = esc.Bold
var bolditalic = esc.BoldItalic
var underline = esc.Underline

func init() {
	var x string
	x = os.Getenv("LESS_TERMCAP_us")
	if x != "" {
		italic = x
	}
	x = os.Getenv("LESS_TERMCAP_md")
	if x != "" {
		bold = x
	}
	x = os.Getenv("LESS_TERMCAP_mb")
	if x != "" {
		bolditalic = x
	}
	x = os.Getenv("LESS_TERMCAP_us")
	if x != "" {
		underline = x
	}
}

// Emph takes a command documentation format string (an extremely
// limited version of Markdown that is also Godoc friendly) and
// transforms it as follows:
//
// * Initial and trailing blank lines are removed.
//
// * Indentation is removed - the number of spaces preceeding the first
//   word of the first line are ignored in every line (including raw text
//   blocks).
//
// * Raw text ignored - any line beginning with four or more spaces
//   (after convenience indentation is removed) will be kept as it is
//   exactly (code examples, etc.) but should never exceed 80 characters
//   (including the spaces).
//
// * Blocks are unwrapped - any non-blank (without three or less initial
//   spaces) will be trimmed line following a line will be joined to the
//   preceding line recursively (unless hard break).
//
// * Hard breaks kept - like Markdown any line that ends with two or
//   more spaces will automatically force a line return.
//
// * URL links argument names and anything else within angle brackets
//   (<url>), will trigger uppercase with underline in both text blocks
//   and usage sections.
//
// * Italic, Bold, and BoldItalic inline emphasis using one, two, or
//   three stars respectivly will be observed and cannot be intermixed or
//   intraword.  Each opener must be preceded by a UNICODE space (or
//   nothing) and followed by a non-space rune. Each closer must be
//   preceded by a non-space rune and followed by a UNICODE space (or
//   nothing).
//
// For historic reasons the following environment variables will be
// observed if found (and also provide color support for the less pager
// utility):
//
//   * Italic      LESS_TERMCAP_so
//   * Bold        LESS_TERMCAP_md
//   * BoldItalic  LESS_TERMCAP_mb
//   * Underline   LESS_TERMCAP_us
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

// Emphasize replaces minimal Markdown-like syntax with *Italic*,
// **Bold**, ***BoldItalic***, and <UNDERLINED_UPPER>
func Emphasize(buf string) string {

	// italic = `<italic>`
	// bold = `<bold>`
	// bolditalic = `<bolditalic>`
	// reset = `<reset>`

	nbuf := []rune{}
	prev := ' '
	opentok := false
	otok := ""
	closetok := false
	ctok := ""
	for i := 0; i < len([]rune(buf)); i++ {
		r := []rune(buf)[i]

		if r == '<' {
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

// Plain is the same as Format but without any Emphasis.
func Plain(input string, indent, width int) (output string) {

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
			output += "\n\n" + Wrap(blockbuf, width-strip-4)
			continue
		}
	}

	// flush last block
	if inblock {
		output += "\n\n" + Wrap(blockbuf, width-strip-4)
	}
	output = Indent(strings.TrimSpace(output), indent)
	return
}
