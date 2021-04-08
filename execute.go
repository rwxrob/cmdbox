package cmdbox

import "github.com/rwxrob/cmdbox/comp"

// Execute traps all panics (see Panic), detects completion and does it, or
// sets Main to the command name passed, injects the Builtin subcommands
// (unless OmitBuiltins is true), looks up the named command from the
// internal command Index and calls it passing cmd.Args. Execute alway exits
// the program.
func Execute(name string) {
	defer TrapPanic()
	command, has := commands[name]
	if !has {
		ExitUnimplemented(name)
	}
	Main = command
	if !OmitAllBuiltins {
		for _, name := range builtins {
			if OmitBuiltins && name[0:1] == "_" {
				continue
			}
			Main.Add(name)
		}
	}
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

// Call allows any indexed subcommand to be called directly by name. Avoid
// using this method as much as possible since it creates very tight coupling
// dependencies between commands. It is included primarily publicly so that
// builtin commands like help, usage, and version can be wrapped with
// internationalized aliases.
func Call(name string, args []string) error {
	defer TrapPanic()
	command, has := commands[name]
	if !has {
		return Unimplemented(name)
	}
	return command.Call(args)
}
