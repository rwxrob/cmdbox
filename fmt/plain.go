package fmt

import (
	"bufio"
	"strings"
)

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
