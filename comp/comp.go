/*
Package comp is the tab completion subpackage of CmdBox that implements the common completion methods needed by most CmdBox command modules.

Currently only Bash Programmable Completion and detection with `complete -C cmd cmd` is supported.
*/
package comp

import (
	"os"
	"strings"
)

// This can be set to force a completion context no matter what the
// shell or situation. This is useful mostly for testing and should
// usually never be modified by any Command of subcommand, which can
// create very unpredictable results for the completions of other
// Commands. (And, as a reminder, always vet any CmdBox command modules
// you import to be sure they do not do insecure things with the cmdbox
// state, such as cmdbox.This.) See Line(), Args(), Word() as well.
var This string

// Func defines a tab completion function that can sense its completion
// context and return a list of completion strings. Also see Line and
// LineArgs and Bash Programmable Completion.
type Func func() []string

// Yes returns true if the current executable is being called in
// a completion context, usually somone tapping tab. Currently, this is
// detected only by the presence of the Bash COMP_LINE environment
// variable. Eventually, other shell completion methods will be added.
// See Line() (which, if length > 0, means true).
func Yes() bool { return len(Line()) > 0 }

// Line returns the full current command line being evaluated for this
// executable being run in completion context. (For Bash it is when
// COMP_LINE is set. See Programmable Completion in the bash man page.)
//
// Only Bash is supported.
func Line() string {
	if This != "" {
		return This
	}
	return os.Getenv("COMP_LINE")
}

// Args returns Line but as a slice of strings. If the Line() has one or
// more spaces at the end include an space (" ") as the last item. This
// is to distinquish between users wanting to tab on prefixes versus
// all the possibilities for a command.
func Args() []string {
	args := []string{}
	line := Line()
	if line == "" {
		return args
	}
	args = strings.Split(line, " ")
	if line[len(line)-1] == ' ' {
		args = append(args, " ")
	}
	return args
}

// Word returns the last of Args or empty string. This includes the
// special single space (" ") string indicating there were trailing
// spaces when completion was invoked. This should not be confused with
// the empty string indicating Args was empty.
func Word() string {
	args := Args()
	if len(args) > 0 {
		return args[len(args)-1]
	}
	return ""
}
