/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/* Package cmdbox is a multicall, modular commander with embedded tab
completion and locale-driven documentation, that prioritizes modern,
speakable human-computer interactions from the command line.
*/
package cmdbox

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rwxrob/cmdbox/comp"
)

const (
	m_invalid_name   = "invalid name (must be lowercase word): %v"
	m_syntax_error   = "syntax error: %v"
	m_unimplemented  = "unimplemented: %v"
	m_bad_type       = "unsupported type: %T"
	m_missing_arg    = "missing argument for %v"
	m_unexpected_arg = "unexpected argument: %v"
	m_missing_caller = "requires caller"
)

// Main is always set to the main command that was used for Execute.
// This can be useful from certain subcommands to query or call directly.
var Main *Command

// Color sets the default output mode for interactive terminals. Set to
// false to force uncolored output for testing, etc. Non-interactive
// terminals have color disabled by default (unless ForceColor is set).
var Color = true

// ForceColor forces color output no matter what the default Color value
// is. This can be used for testing or associating with configuration
// parameters (for example, when a user has a pager that supports color
// output).
var ForceColor = false

// Reg contains the Commands register. See CommandMap and Add.
var Reg = NewCommandMap()

// JSON serializes the current internal package register of commands as
// JSON which can then be used to present documentation of the composite
// command in different forms. Also see YAML. Empty values are always omitted.
func JSON() string { return Reg.JSON() }

// YAML serializes the current internal package register of commands as
// YAML which can then be used to present documentation of the composite
// command in different forms. Empty values are always omitted.
func YAML() string { return Reg.YAML() }

// Print is shortcut for fmt.Println(cmdbox.YAML()) which is mostly only
// useful during testing.
func Print() { Reg.Print() }

// Init initializes (or re-initialized) the package status and empties
// the internal commands register (without changing its reference).
// Init is primarily intended for testing to reset the cmdbox package.
func Init() {
	Reg.Init()
	addHelp()
	addVersion()
}

// Add creates a new Command, adds it to the internal register, and
// returns a pointer to it (assigned to 'x' by convention).  The Add
// function is guaranteed to never return nil but will panic if invalid
// arguments are passed (see Validation below).  Further initialization
// can be done with direct assignments to fields of x from within the
// init() function. By convention only one init() function with a single
// Add call is allowed per file to maintain command modularity.
//
// First Argument is Command Name
//
// The first argument is *always* the main Name by which it will be
// called and becomes x.Name.  This uniquely identifies the Command and
// becomes the key used by Call to lookup the Command from the internal
// register for completion and execution.  By convention, these must be
// speakable, complete words with absolutely no punctuation whatsoever.
// (For performance reasons, no validation is performed on the Name.)
//
// Command Names May Contain Two Words
//
// The Name may contain two complete words separated by a single space.
// This is to avoid collisions and facilitate default tab completion. It
// also removes indirection when called from Execute.
//
//    x := cmdbox.Add("foo help")
//    x.Summary = `output help information for foo`
//
// Using two-word names is common when packaging subcommands with
// commands in such a way as to disambiguate which subcommand is wanted
// -- particularly when common words are used.
//
//    x := cmdbox.New("foo help")
//
// Commands May Have Subcommands
//
// Any variadic arguments that follow will be directly passed to
// x.Add(). This provides a succinct summary of how the command may be
// called. The h|help and version commands are added automatically to
// all Commands with help being set to the x.Default (which can be
// overriden).
//
//    x := cmdbox.New("foo", "bar")
//    // x.Default == "bar"
//
// Subcommands May Have Aliases
//
// Each argument passed in the list may be in signature form (rather
// than just a name) meaning it may have one or more aliases prefixed
// and bar-delimited which are added to the x.Commands Map:
//
//    x := cmdbox.New("foo", "b|bar")
//
// Default Usage Automatically Inferred
//
// Since it is so common to declare everything up front for a new
// Command the x.Default will be set to an optional joined list of all
// commands and aliases as if the following where explicitly assigned:
//
//    x.Usage = `[b|bar]`
//
// Of course, this can always be overriden explicitly by Command
// authors using the same assignment. Optional ([]) is the default since
// any base command since help is automatically the default.
//
// Command Method Has Priority
//
// When the Command (x) is called from cmdbox.Call the x.Commands Map is
// used to delegate the call to a matching Command in the internal
// register if and only if the Command itself does not have a Method
// defined. See Call for more about this delegation on how it finds key
// name matches in the internal register.
//
// All but top-level Commands will usually assign a x.Method to
// handle the work of the Command. By convention the arguments should be
// named "args" and no name given to the error returned:
//
//    x.Method = func(args []string) error {
//        fmt.Println("would do something")
//        return nil
//    }
//
// If a Command has a Method, then Call will pass all arguments as-is
// allowing the Method to decide if they just arguments or keywords for
// actions to be handled within that x.Method (usually within
// a switch/case block). The Method may still cmdbox.Call() to delegate to
// other registered Commands.
//
// No Command Method Will Trigger Default Delegation
//
// If the Command does not have an x.Method of its own, then the list of
// arguments passed to Add is assumed to be the signatures for other
// registered Commands that must eventually be populated by other
// Command init() functions including subcommands of the given Command.
//
// No Assertion of Command Registration
//
// Add does not validate that a potential command has been registered
// because the state of the internal register cannot be predicted at the
// specific time any init function is called. Not all Commands may yet
// have been registered before any other Add is called. This means
// runtime testing is required to check for errant calls to unregistered
// Commands (which otherwise produce a relatively harmless
// "unimplemented" error.)
//
// Duplicate Names Append Underscore
//
// Because CmdBox composite commands may be composed of packages
// imported from a rich eco-system of command module packages it is
// possible that two CmdBox modules might use conflicting names and need
// some resolution by the composite developer who is importing them.
//
// Rather than override any Command previously added with an identical
// Name, Add simply adds an underscore to the name allowing it to be
// identified with Dups. Developer will know of such conflicts in advance
// and be able to easily correct them by calling the Rename function
// before Execute.
//
func Add(name string, a ...string) *Command {
	var x *Command
	for {
		x = Reg.Get(name)
		if x == nil {
			break
		}
		name = name + "_"
	}
	x = NewCommand(name, a...)
	Reg.Set(name, x)
	return x
}

// Names returns a sorted list of all Command names in the internal
// register.
func Names() []string { return Reg.Names() }

// Dups returns key strings of duplicates (which can then be easily
// renamed). Keys are sorted in lexicographic order. See Rename.
func Dups() []string { return Reg.Dups() }

// Rename renames a Command in the Register by adding the
// new name with the same *Command and deleting the old one. This is
// useful when a name conflict causes New to append and underscore (_)
// to the duplicate's name. Rename can be called from init() at any
// point after the duplicate has been added to resolve the conflict.
// Note the order of init() execution --- while predictable --- is not
// always apparent.  When in doubt do Rename from main() to be sure.
// Rename is safe for concurrency.
func Rename(from, to string) { Reg.Rename(from, to) }

// Get returns the *Command for key name if found.
func Get(name string) *Command { return Reg.Get(name) }

// Slice returns a slice of *Command pointers and fetched from the
// internal register that match the key names passed.  If an entry is
// not found it is simply skipped. Will return an empty slice if none
// found.
func Slice(names ...string) []*Command { return Reg.Slice(names...) }

// Set the internal register for the given key to the given *Command
// pointer in a way that is safe for concurrency. Replaces entries that
// already exist. Note that although this allows register key names to
// refer to commands that have an actual x.Name that differs from the
// key this is discouraged, which is why Add and Rename should generally
// be used instead. Also see Add and Get.
func Set(name string, x *Command) { Reg.Set(name, x) }

// Delete deletes one or more commands from the internal register.
func Delete(names ...string) { Reg.Delete(names...) }

// ---------------------- exit and error handling ---------------------

// Exit calls os.Exit(0).
func Exit() { os.Exit(0) }

// ExitError prints err and exits with 1 return value.
func ExitError(err ...interface{}) {
	switch e := err[0].(type) {
	case string:
		if len(e) > 1 {
			fmt.Printf(e+"\n", err[1:])
		}
		fmt.Println(e)
	case error:
		out := fmt.Sprintf("%v", e)
		if len(out) > 0 {
			fmt.Println(out)
		}
	}
	os.Exit(1)
}

// ExitUnimplemented calls Unimplemented and calls ExitError().
func ExitUnimplemented(a string) { ExitError(Unimplemented(a)) }

// TrapPanic recovers from any panic and more gracefully displays the
// error as an exit message. It is used to gaurantee that no CmdBox
// composite command will ever panic (exiting instead). It can be
// redefined to behave differently or set to an empty func() to allow
// the panic to blow up with its full trace log.
var TrapPanic = func() {
	if r := recover(); r != nil {
		ExitError(r)
	}
}

// Unimplemented returns an unimplemented error for the Command passed.
// This function may be overriden by CmdBox command modules from their
// init and main methods to change behavior for everthing in the
// composite command. See "unimplemented" in Messages.
var Unimplemented = func(a string) error {
	return fmt.Errorf(m_unimplemented, a)
}

// UsageError returns an error containing the usage string suitable for
// printing directly.  This function may be overriden by CmdBox command
// modules from their init and main methods to change behavior for
// everthing in the composite command.
var UsageError = func(x *Command) error {
	return fmt.Errorf("usage: %v %v", x.Name, x.Usage)
}

// BadType returns an error containing the bad type attempted.
var BadType = func(v interface{}) error {
	return fmt.Errorf(m_bad_type, v)
}

// Harmless returns an error that is mostly designed to trigger an error
// exit status. This is useful for help and commands like it to help the
// user disambiguate significant output from just help and other error
// output.
var Harmless = func(msg ...string) error {
	if len(msg) > 0 {
		return fmt.Errorf("%v", msg[0])
	}
	return fmt.Errorf("")
}

// MissingArg returns an error stating that the name of the parameter
// for which no argument was found.
var MissingArg = func(name string) error {
	return fmt.Errorf(m_missing_arg, name)
}

// UnexpectedArg returns an error stating that the argument passed was
// unexpected in the given context.
var UnexpectedArg = func(name string) error {
	return fmt.Errorf(m_unexpected_arg, name)
}

// SyntaxErrorPanic panics with the message stating the problem.
var SyntaxErrorPanic = func(msg string) {
	panic(fmt.Sprintf(m_syntax_error, msg))
}

// SyntaxError returns an error with the message stating the problem.
var SyntaxError = func(msg string) error {
	return fmt.Errorf(m_syntax_error, msg)
}

// CallerRequired retuns an error indicating a Command was used
// incorrectly (as designed by the developer) and that is requires being
// called from something else. CmdBox command modules that cannot be
// used as standalones are examples that would have an x.Method that
// might return this error.
var CallerRequired = func() error {
	return fmt.Errorf(m_missing_caller)
}

// --------------------- resolve / call / execute ---------------------

// Resolve looks up a Command from the internal register based on the
// caller and the name. If the Name of the caller and name passed,
// joined with a space (a fully qualified entry) is found
// then that is used instead of just the name. Otherwise, just the name
// is looked up (which might itself already be fully qualified).  The
// returned Command (x) is examined further to decide which Method and Args
// to return:
//
//   * If x.Method defined, call and return it with args unaltered
//
//   * If first arg in x.Commands, recursively Call with shifted args
//
//     * First with x.Name + " " + cmd
//     * Then with just cmd
//
//   * If x.Default defined, recursively Call with shifted args
//
//     * First with x.Name + " " + x.Default
//     * Then with just x.Default
//
//   * Return nil and args
//
// By convention, passing a nil as the caller indicates the Command was
// called from something besides another Command, usually the cmdbox
// package itself. See Call, Command, ExampleResolve for more.
func Resolve(caller *Command, name string, args []string) (Method,
	[]string) {
	var x *Command

	// fully qualified, if found
	if caller != nil {
		full := Reg.Get(caller.Name + " " + name)
		if full != nil {
			x = full
		}
	}

	// plain
	if x == nil {
		x = Reg.Get(name)
	}

	// nothing at all, we're done here
	if x == nil {
		return nil, args
	}

	// so that Commands know their caller
	x.Caller = caller

	// ultimately, this is where recursion stops (successfully)
	if x.Method != nil {
		return x.Method, args
	}

	// check if the first argument is a command with Method
	if len(args) > 0 {
		first := args[0]
		if cmd := x.Commands.Get(first); cmd != "" {
			name = name + " " + cmd
			method, margs := Resolve(caller, name, args[1:])
			if method != nil {
				return method, margs
			}
			method, margs = Resolve(caller, cmd, args[1:])
			if method != nil {
				return method, margs
			}
		}
	}

	// check for default command with method
	if x.Default != "" {
		name = name + " " + x.Default
		method, margs := Resolve(caller, name, args)
		if method != nil {
			return method, margs
		}
		method, margs = Resolve(caller, x.Default, args)
		if method != nil {
			return method, margs
		}
	}

	// out of options
	return nil, args
}

// Call allows any Command in the internal register to be called
// directly by name. The first argument is an optional pointer to the
// calling Command, the second is the required name, and the third is an
// optional list of string arguments (or nil). Resolve is first called
// to get the Command from the internal registry and lookup the proper
// Method and any argument shifting required. If no Method is returned
// Call returns Unimplemented. Otherwise, Method is called with its
// arguments and error result returned.  See Resolve, Command, Execute,
// and ExampleCall as well.
func Call(caller *Command, name string, args ...string) error {
	defer TrapPanic()
	if name == "" {
		return MissingArg("name")
	}
	method, args := Resolve(caller, name, args)
	if method == nil {
		return Unimplemented(name)
	}
	return method(args...)
}

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

var executedAs = filepath.Base(os.Args[0])

// Execute is the main entrypoint into a CmdBox composite command and is
// always called from a main() function. In fact, most composite
// commands that follow the CmdBox subcommand convention of putting each
// into its own file will need nothing more than this in their main.go
// file.
//
//     package main
//     import "github.com/rwxrob/cmdbox"
//     func main() { cmdbox.Execute() }
//
// Execute first determines the name of the command to be executed
// (explicitly passed or inferred from multicall binary, see ExecutedAs)
// and assigns the command to cmdbox.Main; adds the builtin commands;
// traps all panics; and finally Calls the Command. If completion
// context is detected (see comp.Yes), Execute calls x.Complete instead
// of Calling it. Execute is gauranteed to always exit the program
// cleanly. See Call, TrapPanic, and Command.
func Execute(a ...string) {
	defer TrapPanic()
	var name string
	if len(a) > 0 {
		name = a[0]
	} else {
		name = executedAs
	}
	x := Reg.Get(name)
	if x == nil {
		ExitUnimplemented(name)
	}
	Main = x
	x.Add("help")
	x.Add("version")
	x.UpdateUsage()
	if comp.Yes() {
		x.Complete()
		Exit()
	}
	err := Call(x, name, os.Args[1:]...)
	if err != nil {
		ExitError(err)
	}
	Exit()
}
