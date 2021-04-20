package cmdbox

import (
	"encoding/json"
	_fmt "fmt"
	"strings"

	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/util"
)

// Command contains a Method or delegates to  one or more other Commands
// by name. Typically a Command is created within an init() function by
// calling cmdbox.New:
//
//     import (
//           "github.com/rwxrob/cmdbox"
//           "github.com/rwxrob/cmdbox/fmt"
//     )
//
//     func init() {
//          // use x by convention
//          x := cmdbox.New("greet","hi","hello")
//          x.Method = func(args []string) error {
//                if len(args) == 0 {
//                      args = append(args, "hi")
//                }
//                switch args[0] {
//                case "hello":
//                      fmt.Println("*Hello!*")
//                case "hi":
//                      fmt.Println("*Hello!*")
//                default:
//                      return x.UsageError()
//                }
//                return nil
//          }
//     }
//
// Providing the method, documentation, and tab completion rules in
// a single file providing a tight, clean view of the Command that is
// easy for humans and computers to quickly parse. Such Commands can
// also be placed into their own "side-effect" packages for importing
// into other code using the underscore identifier making for
// potentially very succinct commands.
//
//    import (
//        "github.com/rwxrob/cmdbox"
//        _ "github.com/rwxrob/cmdbox-greet"
//    )
//
//    func main() { cmdbox.Run() }
//
// Or, it can be combined with others and composed into an entirely new
// monolith one of its commands:
//
//    import (
//        "github.com/rwxrob/cmdbox"
//        _ "github.com/rwxrob/cmdbox-greet"
//        _ "github.com/rwxrob/cmdbox-timer"
//        _ "github.com/rwxrob/cmdbox-pomo"
//    )
//
//    func init() {
//          x := cmdbox.New("skeeziks","greet","timer","pomo")
//          x.Usage = `[greet|timer|pomo]` // default: greet
//          x.Summary = `a simple command line assistant`
//          // notice no Method
//          // ...
//    }
//
// This modularity allows a CmdBox monolith to compose in an unlimited
// number of commands. Hence the name CmdBox, a nod to BusyBox which
// does something similar but without the ability for sub-command
// composition and coding in Go.
//
// A Command also includes all the documentation for the method as well
// as rules about how to handle tab completion so there is no need for
// extra shell code to be "eval"ed. In Bash, this happens by using the
// -C option of complete telling Bash to use the same command that it is
// completing for completion itself.
//
//    complete -C skeeziks skeeziks
//
// And then the following will complete with the list passed to New():
//
//     skeeziks g<TAB>
//     skeeziks greet
//
// Additional builtin Commands are automatically composed into the
// monolith as well (unless disabled):
//
//     skeeziks h<TAB>
//     skeeziks help
//
// Tab completion rules default to the list of Commands, but can be
// overriden per Command by defining and assigning an anonymous function
// to the CompFunc field (which is passed the value of COMP_LINE in
// Bash). This allows for dynamic tab completion possibilities that have
// nothing to do with sub-Commands and can access program and system
// state for their determination, all written in Go (and not bloated,
// evaled shell).
//
// The Params list is for completion as well. For things that are
// neither Commands nor actions to be handled by Method. While this may
// include dashed options they should be avoided. See
// cmdbox/util/mapopt.go and the general documentation regarding cmdbox
// best practices and design considerations.
//
// The help documentation is multilayered and defaults to using color
// and any pager detected on the system. It is vitually
// indistinguishable from a "man" page, but does not use the man
// command.
//
// Most of the fields can be set to any interface{} which can be
// anything that returns a string, Printfs as a string, or is a string.
// This allows values to be determined when they are used during runtime
// providing dynamic documentation possibilities. See cmdbox/fmt
// package.
//
// For examples of different Command structs search on GitHub for any
// project beginning with cmdbox- such as the following:
//
// * https://github.com/rwxrob/cmdbox-greet
// * https://github.com/rwxrob/cmdbox-pomo
//
type Command struct {
	Name        string                     // <= 14 recommended
	As          string                     // only set if renamed
	Summary     interface{}                // > 65 truncated
	Version     interface{}                // semantic version (v0.1.3)
	Usage       interface{}                // following docopt syntax
	Description interface{}                // long form, see fmt.Emph()
	Examples    interface{}                // long form, see fmt.Emph()
	SeeAlso     interface{}                // links, other commands, etc.
	Author      interface{}                // email format, commas between
	Git         interface{}                // repo location, no schema ok
	Issues      interface{}                // full web URL
	Copyright   interface{}                // legal copyright statement
	License     interface{}                // released under license(s)
	Other       map[string]interface{}     // other (custom) doc sections
	Method      func(args []string) error  // optional method, see Call()
	Params      []string                   // params, completion only
	CompFunc    func(line string) []string // override completion func
	Commands    []string                   // subcmds/actions, see Call()
	Default     interface{}                // default subcmd/action
}

// New initializes a new Command and returns a pointer to it. An
// optional list of sub-Commands can be passed as arguments and will be
// added with Command.Add(). The first argument after the name is
// assigned to Command.Default but can be overriden with a direct
// assignment later.
//
// If a Command has a Method, then the Commands are interpreted as
// keywords for actions to be handled within that Command.Method
// (usually within a switch/case block). The Method may still
// cmdbox.Call() to delegate to other other Commands in the
// cmdbox.Register (but avoid unnecessary coupling between Commands when
// possible).
//
// If the Command does not have a Method of its own, then the list of
// Commands is assumed to be the names of other Commands in the
// cmdbox.Register. However, no validation is done to check that any
// specific Command has been added to the Register. This is because
// usually cmdbox.New() is called from init() and not all Commands may
// yet have been registered with before any other cmdbox.New() is
// called. Note that this means runtime testing is required to check for
// errant calls to unregistered Commands (which otherwise produce an
// "Unimplemented" error.
//
// If any of the Commands are the words 'help' or 'version' then these
// will be fulfilled by the those named as such in packages within the
// cmdtab/prefab package. Only 'help' and 'version' are allowed to be
// overriden by a call to New("help") or New("version") in order to
// allow developers to override these defaults and provide their own
// versions of these for their own applications. These prefabs are just
// the defaults but there is some desire to keep all cmdbox tools as
// consistent as possible, but this is a decision for cmdbox users to
// decide.
//
// If a name conflicts with one that has already been added then an
// underscore (_) is appended to the duplicate name. Late this can be
// renamed with cmdbox.Rename() or removed with cmdbox.Remove() or
// changed directly by accessing it from the cmdbox.Register. If the
// Register key is changed do not forget to set the Command.As field as
// well to match. Generally the original Name should never be changed
// since it is referred to throughout the rest of the embedded Command
// documentation.
//
func New(name string, a ...string) *Command {
	x := new(Command)
	x.Name = name
	x.Commands = []string{}
	Lock()
	if _, has := Register[name]; has && name != "help" && name != "version" {
		name = name + "_"
	}
	Register[name] = x
	Unlock()
	if len(a) > 1 {
		x.Add(a...)
		x.Default = a[0]
	}
	x.Other = map[string]interface{}{}
	return x
}

// Hidden returns true if the command name begins with underscore ('_').
func (x *Command) Hidden() bool { return x.Name[0] == '_' }

// TODO clean me or remove
func (x *Command) vsubcommands() []*Command {
	cmds := []*Command{}
	for _, name := range x.Commands {
		if name[0] == '_' {
			continue
		}
		if command, has := Register[name]; has {
			cmds = append(cmds, command)
		}
	}
	return cmds
}

// SprintUsage returns a usage string (without emphasis) that includes
// a bold Command.Name for each line along with the Main.Usage and each
// individual SprintUsage of every entry in Commands that is not hidden
// (no underscore prefix). If indentation is needed this can be passed
// to fmt.Indent(). To replace emphasis with terminal escapes for
// printing to colored terminals pass to fmt.Emphasize().
func (x *Command) SprintUsage() string {
	buf := ""
	name := x.Name
	if len(x.As) > 0 {
		name = x.As
	}
	if x.Usage != nil {
		buf += "**" + name + "** " +
			strings.TrimSpace(fmt.String(x.Usage)) + "\n"
	}
	for _, subcmd := range x.vsubcommands() {
		buf += "**" + name + "** " + subcmd.SprintUsage()
	}
	if len(buf) > 0 {
		return buf
	}
	return "**" + name + "**"
}

// SprintCommandSummaries returns a printable string that includes
// a bold Command.Name for each line along with the summary string, if
// any, for that Command. This is helpful when creating custom
// builtin help.
func (x *Command) SprintCommandSummaries() string {
	buf := ""
	for _, subcmd := range x.vsubcommands() {
		buf += _fmt.Sprintf(
			"%-14v %v\n",
			"**"+subcmd.Name+"**",
			strings.TrimSpace(fmt.String(subcmd.Summary)),
		)
	}
	return buf
}

// MarshalJSON fulfills the Go JSON marshalling requirements and is
// called by String from the fmt.Stringer interface. Traditional JSON
// struct tagging does not work here because we use interface{} and
// include func() string as a type (which the JSON reflection won't
// understand). Empty values are always omitted. Strings are trimmed and
// long strings such as Description and Examples are written as basic
// Markdown. See fmt.Emph, fmt.Plain, and util.ConvertToJSON.
func (x Command) MarshalJSON() ([]byte, error) {
	s := make(map[string]interface{})
	s["Name"] = x.Name

	// check for empties before commiting
	var buf string

	s["As"] = x.As

	if x.Summary != "" {
		s["Summary"] = strings.TrimSpace(fmt.String(x.Summary))
	}

	buf = strings.TrimSpace(fmt.String(x.Version))
	if buf != "" {
		s["Version"] = buf
	}

	buf = strings.TrimSpace(fmt.String(x.Usage))
	if buf != "" {
		s["Usage"] = buf
	}

	buf = fmt.Emph(fmt.String(x.Description), 0, -1)
	if buf != "" {
		s["Description"] = buf
	}

	buf = fmt.Emph(fmt.String(x.Examples), 0, -1)
	if buf != "" {
		s["Examples"] = buf
	}

	buf = fmt.Emph(fmt.String(x.SeeAlso), 0, -1)
	if buf != "" {
		s["SeeAlso"] = buf
	}

	buf = fmt.Emph(fmt.String(x.Author), 0, -1)
	if buf != "" {
		s["Author"] = buf
	}

	buf = strings.TrimSpace(fmt.String(x.Git))
	if buf != "" {
		s["Git"] = buf
	}

	buf = strings.TrimSpace(fmt.String(x.Issues))
	if buf != "" {
		s["Issues"] = buf
	}

	buf = strings.TrimSpace(fmt.String(x.Copyright))
	if buf != "" {
		s["Copyright"] = buf
	}

	buf = strings.TrimSpace(fmt.String(x.License))
	if buf != "" {
		s["License"] = buf
	}

	// add custom (other) sections to docs
	for k, v := range x.Other {
		s[k] = fmt.String(v)
	}

	// skip CompFunc
	// skip Method

	if len(x.Params) > 0 {
		s["Params"] = x.Params
	}

	if len(x.Commands) > 0 {
		s["Commands"] = x.Commands
	}

	buf = strings.TrimSpace(fmt.String(x.Default))
	if buf != "" {
		s["Default"] = buf
	}

	return json.Marshal(s)
}

// Fulfills the fmt.Stinger interface rendering a Command as a JSON
// string.
func (x Command) String() string {
	return util.ConvertToJSON(x)
}

// VersionLine returns a single line with the combined values of the Name,
// Version, Copyright, and License. If Version is empty or nil an empty string
// is returned instead. VersionLine() is used by the version builtin command
// to aggregate all the version information into a single output.
func (x *Command) VersionLine() string {
	version := fmt.String(x.Version)
	if version == "" || x.Name == "" {
		return ""
	}
	copyright := fmt.String(x.Copyright)
	license := fmt.String(x.License)
	buf := x.Name + " " + version
	if copyright != "" {
		buf += " " + copyright
	}
	if license != "" {
		buf += " (" + license + ")"
	}
	return buf
}

// Has returns true if name matches an of the Commands.
func (x *Command) Has(name string) bool {
	for _, sc := range x.Commands {
		if sc == name {
			return true
		}
	}
	return false
}

// Add adds the list of Command (or action) names passed skipping any it
// already has.  If any name contains a bar (|) then it will be split
// with the last item assumed to be the actual name and the first
// elements considered aliases (all of which are added individually to
// Commands.) See Command and Command.Run.
func (x *Command) Add(names ...string) {
	for _, name := range names {
		if strings.ContainsRune(name, '|') {
			for _, n := range strings.Split(name, "|") {
				x.Add(n)
			}
			return
		}
		if !x.Has(name) {
			x.Commands = append(x.Commands, name)
		}
	}
}

// CommandUsage returns the Usage strings for each Command in list of
// Command.Commands. This is useful when creating usages that have
// additional notes or formatting when it is desirable to loop through
// the Command.Usage strings. The order is gauranteed to match the order
// of Command.Commands even if the Usage For a particular Command is
// empty.
func (x *Command) CommandUsage() []string {
	usages := []string{}
	for _, name := range x.Commands {
		usage := Register[name].Usage
		usages = append(usages, fmt.String(usage))
	}
	return usages
}

// UsageError is frequently returned from within Command.Method
// definitions when something about the arguments to the Command or its
// input is wrong.
func (x *Command) UsageError() error {
	return fmt.Errorf(fmt.Emphasize(strings.TrimSpace("**usage:** " + x.SprintUsage())))
}

// Complete prints the tab completion replies for the current context.
// The current command line (COMP_LINE) is passed exactly as detected
// leaving maximum flexibility for parsing and matching from the
// optional Command.CompFunc function if defined. Otherwise, the names
// of the Commands and any Params are used to provide the list for
// completion. See Command.CompFunc and Programmable Completion in the
// Bash man page.
func (x *Command) Complete(line string) {
	if x.CompFunc != nil {
		for _, name := range x.CompFunc(line) {
			fmt.Println(name)
		}
		return
	}
	words := strings.Split(strings.TrimSpace(line), " ")
	if len(words) >= 2 {
		name := words[len(words)-2]
		complete := words[len(words)-1]
		if x.Name != name {
			if cmd, has := Register[name]; has {
				line = strings.Join(words[len(words)-2:], " ")
				cmd.Complete(line)
			}
			return
		}
		for _, cmdname := range x.Commands {
			if cmdname == complete {
				if cmd, has := Register[cmdname]; has {
					line = strings.Join(words[len(words)-1:], " ")
					cmd.Complete(line)
				}
				return
			}
			if strings.HasPrefix(cmdname, complete) {
				fmt.Println(cmdname)
			}
		}
		if x.Params != nil {
			for _, param := range strings.Split(fmt.String(x.Params), " ") {
				if complete != param && strings.HasPrefix(param, complete) {
					fmt.Println(param)
				}
			}
		}
		return
	}

	// do not include hidden commands in completion
	for _, cmdname := range x.Commands {
		if cmdname[0] != '_' {
			fmt.Println(cmdname)
		}
	}

	// always include params in completion
	for _, param := range strings.Split(fmt.String(x.Params), " ") {
		fmt.Println(param)
	}

}

func (x *Command) Title() string {
	summary := fmt.String(x.Summary)
	if len(summary) > 0 {
		return x.Name + " - " + summary
	}
	return x.Name
}

// Call invokes Method if it has one, or assumes first argument is
// from Commands and attempts to delegate to it. If that fails,
// assumes arguments are meant for the Default Command. Otherwise,
// returns a UsageError.
func (x *Command) Call(args []string) error {

	// if has own method assume it can handle itself and args
	if x.Method != nil {
		return x.Method(args)
	}

	// if no method and no args try default
	if args == nil || len(args) == 0 {
		if x.Default != nil {
			return Call(fmt.String(x.Default), args)
		}
		return fmt.Errorf("empty Call() arguments with no Default subcommand")
	}

	// if first arg is in list of subcommands call it
	first := args[0]
	for _, name := range x.Commands {
		if name == first {
			return Call(name, args[1:])
		}
	}

	// assume arguments are for default
	if len(args) > 0 && x.Default != nil {
		return Call(fmt.String(x.Default), args)
	}

	return x.UsageError()
}

// Unimplemented calls Unimplemented passing the name of the command.
// Useful for temporarily notifying users of commands in beta that
// something has not yet been implemented.
func (x *Command) Unimplemented() error { return Unimplemented(x.Name) }
