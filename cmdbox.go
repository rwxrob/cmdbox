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
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/util"
	"gopkg.in/yaml.v3"
)

var state = map[string]interface{}{
	"commands": reg.reg,
	"messages": Messages,
}

var reg = func() *register {
	r := new(register)
	r.init()
	return r
}()

var mut sync.Mutex

type register struct {
	reg map[string]*Command
	sync.Mutex
}

// Main is always set to the main command that was used for Execute.
// This can be useful from certain subcommands to query or call directly.
var Main *Command

// Reg provides direct access to the otherwise encapsulated and internal
// register. This is to allow CmdBox composite command authors to more
// fully control and manipulate the register directly but should be done
// with caution. Always provide some safety (mutex, etc.) for
// concurrency when reading or writing to the returned map.
func Reg() map[string]*Command { return reg.reg }

// JSON serializes the current state of the cmdbox package with its
// internal message set, register and every Command in it as JSON which
// can then be used to present documentation of the composite command in
// different forms. Also see YAML and Load for a way to overwrite this
// data in such a way as to provide dynamic language and locale
// detection and adjustments. Empty values are always omitted.
func JSON() string { return util.ToJSON(state) }

// YAML serializes the current state of the cmdbox package with its
// internal message set, register and every Command in it as YAML which
// can then be used to present documentation of the composite command in
// different forms. See Load for a way to overwrite this data in such
// a way as to provide dynamic language and locale detection and
// adjustments. Empty values are always omitted.
func YAML() string { return util.ToYAML(state) }

// Print is shortcut for fmt.Println(cmdbox.YAML()) which is mostly only
// useful during testing.
func Print() { fmt.Print(YAML()) }

// PrintReg is shortcut for util.PrintYAML(cmdbox.Reg()) which is mostly
// only useful during testing.
func PrintReg() { util.PrintYAML(reg.reg) }

// Init initializes (or re-initialized) the package status and empties
// the internal commands register (without changing its reference) and
// sets the standard Messages back to defaults (but does not override
// any newly added Messages). Init is primarily intended for testing to
// reset the cmdbox package.
func Init() {
	defer mut.Unlock()
	mut.Lock()
	reg.init()
	Messages["invalid_name"] = m_invalid_name
	Messages["unimplemented"] = m_unimplemented
	Messages["syntax_error"] = m_syntax_error
	Messages["missing_arg"] = m_missing_arg
	Messages["bad_type"] = m_bad_type
}

func (r *register) init() {
	defer r.Unlock()
	r.Lock()
	if r.reg == nil {
		r.reg = map[string]*Command{}
		return
	}
	for k := range r.reg {
		delete(r.reg, k)
	}
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
// called. The first in the list will be assigned to x.Default, but can
// be overridden with a direct assignment later.
//
//    x := cmdbox.New("foo", "help")
//    // x.Default == "help"
//
// Subcommands May Have Aliases
//
// Each argument passed in the list may be in signature form (rather
// than just a name) meaning it may have one or more aliases prefixed
// and bar-delimited which are added to the x.Commands Map:
//
//    x := cmdbox.New("foo", "h|help")
//    // x.Default == "help"
//    // x.Commands == {"h":"help","help":"help"}
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
// specific time any init if called. Not all Commands may yet have been
// registered before any other Add is called. This means runtime testing
// is required to check for errant calls to unregistered Commands (which
// otherwise produce a relatively harmless "unimplemented" error.)
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
// Commands h|help and version Automatically Injected
//
// If h|help or version commands are not already included in the list of
// Commands then they will be automatically appended to the very end of
// the list and set to call the internal help (see help.go) and version
// Commands (see version.go). These can still be overriden explicitly.
// Note that the means any CmdBox module with any of its own Commands
// will have a default automatically set to the first of the declared
// Commands in the argument list. To make help the default, be sure to
// explicitly add it as `h|help` for the first argument in the list.
func Add(name string, a ...string) *Command {
	var x *Command
	for {
		x = reg.get(name)
		if x == nil {
			break
		}
		name = name + "_"
	}
	x = NewCommand(name, a...)
	reg.set(name, x)
	if _, has := x.Commands["version"]; !has {
		x.Add("version")
	}
	if _, has := x.Commands["help"]; !has {
		x.Add("h|help")
	}
	return x
}

// Names returns a sorted list of all Command names in the internal
// register.
func Names() []string {
	names := []string{}
	for name, _ := range reg.reg {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Dups returns key strings of duplicates (which can then be easily
// renamed). Keys are sorted in lexicographic order. See Rename.
func Dups() []string { return reg.dups() }

func (r *register) dups() []string {
	defer r.Unlock()
	r.Lock()
	var keys []string
	for k, _ := range r.reg {
		if k[len(k)-1] == '_' {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// Rename renames a Command in the Register by adding the
// new name with the same *Command and deleting the old one. This is
// useful when a name conflict causes New to append and underscore (_)
// to the duplicate's name. Rename can be called from init() at any
// point after the duplicate has been added to resolve the conflict.
// Note the order of init() execution --- while predictable --- is not
// always apparent.  When in doubt do Rename from main() to be sure.
// Rename is safe for concurrency.
func Rename(from, to string) { reg.rename(from, to) }

func (r *register) rename(from, to string) {
	if to == "" || from == "" {
		return
	}
	x, has := r.reg[from]
	if !has {
		return
	}
	r.Lock()
	x.Name = to
	r.reg[to] = x
	delete(r.reg, from)
	r.Unlock()
}

// Load buffers the input and unmarshals it into the package scope over
// any existing matches for the package state including the internal
// Messages and Commands. Input must be valid YAML data (which includes
// JSON). This allows CmdBox composites to dynamically adapt to language
// and locale. Attempts to load Command data that has not already been
// registered (with Add or Set) will silently fail. Messages will always
// be unmarshaled as is so that authors can add their own. See JSON,
// YAML, Command.Update, and String as well.
func Load(in interface{}) error {
	var buf []byte
	var err error
	switch v := in.(type) {
	case io.Reader:
		buf, err = io.ReadAll(v)
		if err != nil {
			return err
		}
	case []byte:
		buf = v
	default:
		return BadType(v)
	}
	m := map[string]interface{}{}
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return err
	}
	if v, has := m["messages"]; has {
		mut.Lock()
		for k, v := range v.(map[string]interface{}) {
			Messages[k] = v.(string)
		}
		mut.Unlock()
	}
	if v, has := m["commands"]; has {
		for name, cmd := range reg.reg {
			if i, has := v.(map[string]interface{})[name]; has {
				cmd.Update(i)
			}
		}
	}
	return nil
}

// LoadFS loads the specified file from the filesystem passed. Any
// filesystem that satisfies the FS interface from io/fs will work. This
// includes local files, files transferred over HTTP, and embeds. For
// example, multiple language/locale files could be embedded into the
// binary at compilation time to be used after locale detection. The
// method of detection, however, is left to the caller rather than
// implied. Also see the go:embed compiler directive for ways to ship
// single executables with multi-lingual/-locale support.  The util
// subpackage of this module may also contain helpful tools for
// determining locale and such to help identify which file to pass. See
// Load for more.
func LoadFS(fsys fs.FS, file string) error {
	buf, err := fs.ReadFile(fsys, file)
	if err != nil {
		return err
	}
	return Load(buf)
}

// Get returns the *Command for key name if found.
func Get(name string) *Command { return reg.get(name) }

func (r *register) get(name string) *Command {
	defer r.Unlock()
	r.Lock()
	if x, has := r.reg[name]; has {
		return x
	}
	return nil
}

// Slice returns a slice of *Command pointers and fetched from the
// internal register that match the key names passed.  If an entry is
// not found it is simply skipped. Will return an empty slice if none
// found.
func Slice(names ...string) []*Command { return reg.slice(names) }

func (r *register) slice(names []string) []*Command {
	defer r.Unlock()
	r.Lock()
	cmds := []*Command{}
	for _, name := range names {
		if x, has := r.reg[name]; has {
			cmds = append(cmds, x)
		}
	}
	return cmds
}

// Set the internal register for the given key to the given *Command
// pointer in a way that is safe for concurrency. Replaces entries that
// already exist. Note that although this allows register key names to
// refer to commands that have an actual x.Name that differs from the
// key this is discouraged, which is why Add and Rename should generally
// be used instead. Also see Add and Get.
func Set(name string, x *Command) { reg.set(name, x) }

func (r *register) set(name string, x *Command) {
	defer r.Unlock()
	r.Lock()
	r.reg[name] = x
}

// Delete deletes one or more commands from the internal register.
func Delete(names ...string) { reg.del(names) }

func (r *register) del(names []string) {
	defer r.Unlock()
	r.Lock()
	for _, k := range names {
		delete(r.reg, k)
	}
}

// ---------------------- exit and error handling ---------------------

// Exit calls os.Exit(0).
func Exit() { os.Exit(0) }

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
	return fmt.Errorf(Messages["unimplemented"], a)
}

// UsageError returns an error containing the usage string suitable for
// printing directly.  This function may be overriden by CmdBox command
// modules from their init and main methods to change behavior for
// everthing in the composite command.
var UsageError = func(x *Command) error { return fmt.Errorf(x.Usage) }

// BadType returns an error containing the bad type attempted.
var BadType = func(v interface{}) error {
	return fmt.Errorf(Messages["bad_type"], v)
}

// MissingArg returns an error stating that the name of the parameter
// for which no argument was found.
var MissingArg = func(name string) error {
	return fmt.Errorf(Messages["missing_arg"], name)
}

// SyntaxErrorPanic panics with the message stating the problem.
var SyntaxErrorPanic = func(msg string) {
	panic(fmt.Sprintf(Messages["syntax_error"], msg))
}

// SyntaxError returns an error with the message stating the problem.
var SyntaxError = func(msg string) error {
	return fmt.Errorf(Messages["syntax_error"], msg)
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
	x := reg.get(name)

	// override with fully qualified, if found
	if caller != nil {
		full := reg.get(caller.Name + " " + name)
		if full != nil {
			x = full
		}
	}

	// nothing to even start from, bailing
	if x == nil {
		return nil, args
	}

	// ultimately, this is where recursion stops (successfully)
	if x.Method != nil {
		x.Caller = caller
		return x.Method, args
	}

	// check if the first argument is a command with Method
	if len(args) > 0 {
		first := args[0]
		if cmd, has := x.Commands[first]; has {
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
func Call(caller *Command, name string, args []string) error {
	defer TrapPanic()
	if name == "" {
		return MissingArg("name")
	}
	method, args := Resolve(caller, name, args)
	if method == nil {
		return Unimplemented(name)
	}
	return method(args)
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
// and assigns the command to cmdbox.Main. It traps all panics and Calls
// the Command. If completion context is detected (see comp.Yes),
// Execute calls x.Complete instead of Calling it. Execute is gauranteed
// to always exit the program cleanly. See Call, TrapPanic, and Command.
func Execute(a ...string) {
	defer TrapPanic()
	var name string
	if len(a) > 0 {
		name = a[0]
	} else {
		name = executedAs
	}
	x := reg.get(name)
	if x == nil {
		ExitUnimplemented(name)
	}
	Main = x
	if comp.Yes() {
		x.Complete()
		Exit()
	}
	err := Call(nil, name, os.Args[1:])
	if err != nil {
		ExitError(err)
	}
	Exit()
}
