package comp

import (
	"os"
	"strings"
)

// Line returns the full current command line being evaluated for this
// executable being run in completion context. (For Bash it is when
// COMP_LINE is set. See Programmable Completion in the bash man page.)
//
// Only Bash is supported.
func Line() string { return os.Getenv("COMP_LINE") }

// LineArgs returns Line but as a slice of strings.
func LineArgs() []string { return strings.Split(Line(), " ") }
