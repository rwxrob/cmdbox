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

// Execute traps all panics, detects completion context and completes,
// or looks up the Command pointer for name from cmdbox.Register sets
// cmdbox.Main to it, adds the 'help' and 'version' prefabs then Calls
// it passing all but the first argument from os.Args. Execute is
// gauranteed to always exit the program cleanly. See Register, Main,
// TrapPanic(). If no name is passed will infer from the name of the
// executable in multicall fashion (akin to BusyBox, see ExecutedAs).
func Execute(a ...string) {
	defer TrapPanic()
	var name string
	if len(a) > 0 {
		name = a[0]
	} else {
		name = executedAs
	}
	command, has := Register[name]
	if !has {
		ExitUnimplemented(name)
	}
	Main = command
	if comp.Yes() {
		Main.Complete()
		Exit()
	}
	err := command.Call(os.Args[1:])
	if err != nil {
		ExitError(err)
	}
	Exit()
}

// Call allows any Command in the Register to be called directly by
// name. Avoid abusing this method since it creates very tight coupling
// dependencies between Commands.
func Call(name string, args []string) error {
	defer TrapPanic()
	command, has := Register[name]
	if !has {
		return Unimplemented(name)
	}
	return command.Call(args)
}
