package cmdbox

import (
	"os"
	"path/filepath"

	"github.com/rwxrob/cmdbox/comp"
)

var executedAs = filepath.Base(os.Args[0])

// ExecutedAs returns the multicall inferred name of the executable as
// it was called during the init() phase. The multicall approach (akin
// to BusyBox) allows the binary to be renamed, hard or soft linked, or
// copied, effectively changing the behavior simply by changing the
// resulting changed name. For security reasons this name may never be
// changed at runtime (even though some applications in the UNIX past
// have employed such methods to communicate information through the
// changed name of a running executable and the resulting ps command
// output). When the Execute function is called without any arguments
// the ExecutedAs value is inferred automatically.
func ExecutedAs() string { return executedAs }

// Execute first determines the name of the command to be executed,
// traps all panics, detects completion context and completes, or looks
// up the Command pointer for name from cmdbox.Register setting
// cmdbox.Main to it. Execute is gauranteed to always exit the program
// cleanly. See Register, Main, TrapPanic(). If no name is passed will
// infer from the name of the executable in multicall fashion (akin to
// BusyBox, see ExecutedAs).
func Execute(a ...string) {
	defer TrapPanic()
	var name string
	if len(a) > 0 {
		name = a[0]
	} else {
		name = executedAs
	}
	x, has := Register[name]
	if !has {
		ExitUnimplemented(name)
	}
	Main = x
	if comp.Yes() {
		Main.Complete()
		Exit()
	}
	err := Call(name, os.Args[1:], nil)
	if err != nil {
		ExitError(err)
	}
	Exit()
}
