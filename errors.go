package cmdbox

import (
	"fmt"
	"os"
)

// Exit just exits with 0 return value but observes the KeepAlive
// variable.
func Exit() {
	if !KeepAlive {
		os.Exit(0)
	}
}

// ExitError prints err and exits with 1 return value.
func ExitError(err ...interface{}) {
	switch e := err[0].(type) {
	case string:
		if len(err) > 1 {
			fmt.Printf(e+"\n", err[1:])
		}
		fmt.Println(e)
	case error:
		fmt.Printf("%v\n", e)
	}
	if !KeepAlive {
		os.Exit(1)
	}
}

// ExitUnimplemented calls Unimplemented and calls ExitError().
func ExitUnimplemented(thing interface{}) {
	ExitError(Unimplemented(thing))
}

// TrapPanic recovers from any panic and more gracefully displays the
// error as an exit message. It can be redefined to behave differently
// or set to an empty func() to allow the panic to blow up with its full
// trace log.
var TrapPanic = func() {
	if r := recover(); r != nil {
		ExitError(r)
	}
}

// Unimplemented just returns an unimplemented error for the thing passed.
func Unimplemented(thing interface{}) error {
	return fmt.Errorf("unimplemented: %v", thing)
}
