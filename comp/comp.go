/*
Package comp is the tab completion subpackage of CmdBox that implements the common completion methods needed by most CmdBox command modules.

Currently only Bash Programmable Completion and detection with `complete -C cmd cmd` is supported.
*/
package comp

// Func defines a tab completion function that can sense its completion
// context and return a list of completion strings. Also see Line and
// LineArgs and Bash Programmable Completion.
type Func func() []string
