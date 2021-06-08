package cmdbox

import (
	"strings"

	"github.com/rwxrob/cmdbox/fmt"
)

// Call allows any Command in the Register to be called directly by
// name. If the caller is not nil and the name does not already have
// a space in it then the caller.Name+" "+name will first be
// tried before just the name as is. By convention, passing a nil as the
// caller indicates the Command was called from something besides
// another Command, usually the cmdbox package itself.
//
// Delegation resolves as follows (for x as *Command):
//
//     1. If x.Method is not nil, delegate to it:
//            return x.Method(args)
//     2. If first arg in x.Commands, shift args and delegate:
//            return Call(caller, x.Name+" "+first, args[1:])
//     3. If x.Default, assume all args intended for default:
//            return Call(caller,x.Name+" "+x.Default,args)
//     4. Return x.UsageError
//
// Note the recursive call to Call itself does not change the original
// caller. (By design, there is no implementation of call stack tracking
// of any kind.)
//
func Call(caller *Command, name string, args []string) error {
	defer TrapPanic()

	// most common case
	x, has := Register[name]

	if caller != nil && !strings.ContainsRune(name, ' ') {
		x, has = Register[caller.Name+" "+name]
	}

	if !has {
		return Unimplemented(name)
	}

	x.Caller = caller

	if x.Method != nil {
		return x.Method(args)
	}

	// if first arg is in map of Commands call it
	if len(args) > 0 {
		first := args[0]
		for k, name := range x.Commands {
			if k == first {
				return Call(caller, name, args[1:])
			}
		}
	}

	if x.Default != nil {
		return Call(caller, fmt.String(x.Default), args)
	}

	return x.UsageError()
}
