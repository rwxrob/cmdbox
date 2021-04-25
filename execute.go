package cmdbox

import (
	"github.com/rwxrob/cmdbox/comp"
)

// Execute traps all panics, detects completion context and completes,
// or looks up the Command pointer for name from cmdbox.Register sets
// cmdbox.Main to it, adds the 'help' and 'version' prefabs (if they are
// not yet added), then Calls it passing cmd.Args.  Execute is
// gauranteed to always exit the program cleanly. See Register, Main,
// TrapPanic().
func Execute(name string) {
	defer TrapPanic()
	command, has := Register[name]
	if !has {
		ExitUnimplemented(name)
	}
	Main = command
	if comp.Line != "" {
		Complete()
		Exit()
	}
	err := command.Call(Args)
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
