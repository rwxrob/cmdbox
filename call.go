package cmdbox

import "github.com/rwxrob/cmdbox/fmt"

// Call allows any Command in the Register to be called directly by
// name. Avoid abusing this method since it creates very tight coupling
// dependencies between Commands. The key parameter must be the key name
// of an existing Command in the Register. The caller points to the
// Command that is doing the calling. If the key name is not found an
// Unimplemented error will be returned. By convention, passing a nil
// as the caller indicates the Command was called from something besides
// another Command, usually the cmdbox package itself.
//
// Call uses the following algorithm to determine what to run:
//
// 1. If a Command.Method is found, run and return
// 2. If first argument in Command.Commands, run and return
// 3. If Command.Default assume arguments are for it, run and return
// 4. Return a UsageError
//
func Call(key string, args []string, caller *Command) error {
	defer TrapPanic()
	x, has := Register[key]
	if !has {
		return Unimplemented(key)
	}
	x.Caller = caller

	// if has own method assume it can handle itself and args
	if x.Method != nil {
		return x.Method(args)
	}

	// if no method and no args try default
	if args == nil || len(args) == 0 {
		if x.Default != nil {
			return Call(fmt.String(x.Default), args, caller)
		}
		return fmt.Errorf("empty Call() arguments with no Command.Default")
	}

	// if first arg is in list of Commands call it
	first := args[0]
	for _, name := range x.Commands {
		if name == first {
			return Call(name, args[1:], caller)
		}
	}

	// assume arguments are for default
	if len(args) > 0 && x.Default != nil {
		return Call(fmt.String(x.Default), args, caller)
	}

	return x.UsageError()
}
