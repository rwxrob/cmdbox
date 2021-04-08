/*
Package comp is the tab completion subpackage of CmdBox. Currently only Bash Programmable Completion and detection with `complete -C cmd cmd` is supported.
*/
package comp

import "os"

// Line is set if a completion context from the shell is detected. (For
// Bash it is COMP_LINE. See Programmable Completion in the bash man
// page.)
var Line string

func init() {
	Line = os.Getenv("COMP_LINE") // bash only for now
}
