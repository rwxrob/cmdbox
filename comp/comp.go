/*
Package comp is the tab completion subpackage of CmdBox that implements the common completion methods needed by most CmdBox command modules.

Currently only Bash Programmable Completion and detection with `complete -C cmd cmd` is supported.
*/
package comp

import "os"

// Func defines a tab completion function that can sense its completion
// context and return a list of completion strings. Also see Line and
// LineArgs and Bash Programmable Completion.
type Func func() []string

// Yes returns true if the current executable is being called in
// a completion context, usually somone tapping tab. Currently, this is
// detected only by the presence of the Bash COMP_LINE environment
// variable. Eventually, other shell completion methods will be added.
func Yes() bool { return len(os.Getenv("COMP_LINE")) > 0 }
