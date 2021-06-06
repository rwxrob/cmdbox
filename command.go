package cmdbox

import (
	"encoding/json"
	_fmt "fmt"
	"strings"
	"sync"

	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/util"
)

// Command contains a Method or delegates to  one or more other Commands
// by name. Typically a Command is created within an init() function by
// calling cmdbox.New:
//
//     import (
//         "github.com/rwxrob/cmdbox"
//         "github.com/rwxrob/cmdbox/fmt"
//     )
//
//     func init() {
//         // use x by convention
//         x := cmdbox.New("greet","hi","hello")
//         x.Method = func(args []string) error {
//             if len(args) == 0 {
//                 args = append(args, "hi")
//             }
//             switch args[0] {
//             case "hello":
//                 fmt.Println("*Hello!*")
//             case "hi":
//                 fmt.Println("*Hello!*")
//             default:
//                 return x.UsageError()
//             }
//             return nil
//         }
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
//          "github.com/rwxrob/cmdbox"
//          _ "github.com/rwxrob/cmdbox-greet"
//    )
//
//    func main() { cmdbox.Run() }
//
// Or, it can be combined with others and composed into an entirely new
// monolith one of its commands:
//
//    import (
//          "github.com/rwxrob/cmdbox"
//          _ "github.com/rwxrob/cmdbox-greet"
//          _ "github.com/rwxrob/cmdbox-timer"
//          _ "github.com/rwxrob/cmdbox-pomo"
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
// Tab completion rules default to the list of Commands and Parameter,
// but can be overriden per Command by defining and assigning an
// anonymous closure function to the CompFunc field (see comp.Func type).
//
// If a Command.CompFunc is not assigned then Command.Complete will
// delegate to the package cmdbox.CompFunc passing it a pointer to the
// Command. In this way the default completion behavior of all Commands
// can be easily tested and changed, even at run time.
//
// This allows for dynamic tab completion possibilities that have
// nothing to do with sub-Commands and can access program and system
// state for their determination.
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
	Name        string                    // <= 14 recommended
	Summary     interface{}               // > 65 truncated
	Version     interface{}               // semantic version (v0.1.3)
	Usage       interface{}               // following docopt syntax
	Description interface{}               // long form, see fmt.Emph()
	Examples    interface{}               // long form, see fmt.Emph()
	SeeAlso     interface{}               // links, other commands, etc.
	Author      interface{}               // email format, commas between
	Git         interface{}               // repo location, no schema ok
	Issues      interface{}               // full web URL
	Copyright   interface{}               // legal copyright statement
	License     interface{}               // released under license(s)
	Other       map[string]interface{}    // other (custom) doc sections
	Method      func(args []string) error // optional method, see Call()
	Caller      *Command                  // last caller, see Call()
	CompFunc    comp.Func                 // set tab completion function
	Commands    CommandsMap               // actions and aliases to actions, see Add()
	Params      []string                  // params, completion only
	Default     interface{}               // default subcmd/action
}

// CommandsMap contains the command names and aliases used for
// completion pointing to the command or action name to be used.
type CommandsMap map[string]string

// String fulfills the fmt.Stringer interface to print as JSON.
func (c CommandsMap) String() string { return util.ConvertToJSON(c) }

// New initializes a new Command and returns a pointer to it (assigned
// to 'x' by convention) allowing further initialization through direct
// assignments (ex: x = ...) thereafter within the init() function. By
// convention only one init() function with a single New call is allowed
// per file to maintain command modularity:
//
//     package mycmd
//
//     import "github.com/rwxrob/cmdbox"
//
//     func init() {
//         x := New("foo","subcmd")
//         x.Summary = `foo is a thing`
//         // ...
//     }
//
// See the package testable example for more.
//
// Up to three command names may be passed with a single space between
// them to avoid collisions and facilitate default tab command
// completion. The last item is always taken as the Command.Name.
//
//         x := New("foo bar help")
//         x.Summary = `output help information for foo`
//
// An optional list or arguments can be passed and each will be directly
// passed to Command.Add(). The first in the list of Commands will be
// assigned to Command.Default but can be overridden with a direct
// assignment later.
//
//         x := New("foo", "help")
//         // x.Default == "help"
//
// Each in the list of Commands may be in signature form (rather than
// just a name) meaning it may have one or more command aliases prefixed
// and bar-delimited which are added to the Command.Commands map:
//
//         x := New("foo", "h|help")
//         // x.Default == "help"
//         // x.Commands == {"h":"help","help":"help"}
//
// All but top-level Commands will usually assign a Command.Method to
// handle the work of the Command. By convention the arguments should be
// named "args" and no name given to the error returned:
//
//         x.Method = func(args []string) error {
//             fmt.Println("would do something")
//             return nil
//         }
//
// If a Command has a Method, then the Commands passed as a list to New
// are interpreted as keywords for actions to be handled within that
// Command.Method (usually within a switch/case block). The Method may
// still cmdbox.Call() to delegate to other other Commands in the
// cmdbox.Register (but avoid unnecessary coupling between Commands when
// possible). (See Call for more.)
//
// If the Command does not have a Method of its own, then the list of
// arguments passed to New is assumed to be the signatures for other
// Commands in the cmdbox.Register.
//
// A Command signature may also include an explicit parent command for
// disambiguation (ex: "foo h|help"). Usually this is not needed since
// New automatically adds the name of the current command when first
// checking the Register (try cmdbox.Register.Print() to see) and only
// after no such qualified name/key is found will look for unqualified
// (global) entries.
//
// Note that New does no validation of any potential command in the
// Register because the state of the Register cannot be predicted at
// init() time. Not all Commands may yet have been registered before any
// other cmdbox.New is called. This means runtime testing is required to
// check for errant calls to unregistered Commands (which otherwise
// produce a relatively harmless "Unimplemented" error.)
//
// The New function is guaranteed to never return nil.
//
func New(name string, a ...string) *Command {
	defer Unlock()
	Lock()

	// TODO detect single spaces in the name

	x := new(Command)
	x.Name = name
	Register[name] = x

	if len(a) > 0 {
		x.Add(a...)
		x.Default = NameFromSig(a[0])
	}

	x.Other = map[string]interface{}{}

	return x
}

// Hidden returns true if the command name begins with underscore ('_').
func (x *Command) Hidden() bool { return x.Name[0] == '_' }

// VisibleCommands returns an array of visual Commands (not beginning
// with underscore). These are used in usage and descriptions and do not
// include any command aliases.
func (x *Command) VisibleCommands() []*Command {
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
	if x.Usage != nil {
		buf += "**" + name + "** " +
			strings.TrimSpace(fmt.String(x.Usage)) + "\n"
	}
	for _, subcmd := range x.VisibleCommands() {
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
	for _, subcmd := range x.VisibleCommands() {
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

var addmut = new(sync.Mutex)

// Add adds the list of Command signatures passed. A command signature
// consists of one or more aliases separated by a bar (|) with the final
// word being the name of the actual Command.  Aliases are a useful way
// to provide shortcuts when tab completion is not available and should
// generally be considered for every Command. Single letter aliases are
// common and encouraged.
//
// Note that Add does not validate inclusion in the Register since in
// many cases there may not yet be a Register entry, and in the case of
// actions handled entirely by the Command itself there never will be.
// See Command.Commands and Command.Run.
//
func (x *Command) Add(sigs ...string) {
	defer func() { addmut.Unlock() }()
	addmut.Lock()
	if x.Commands == nil {
		x.Commands = map[string]string{}
	}
	for _, name := range sigs {
		aliases := strings.Split(name, "|")
		name = aliases[len(aliases)-1]
		for _, alias := range aliases {
			x.Commands[alias] = name
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

// Complete prints the possible strings based on the current Command and
// completion context. If the Commands CompFunc has been assigned (not
// nil) it is called and passed a its own pointer.  If CompFunc has not
// been assigned (is nil) then cmdbox.CompFunc is called instead. This
// allows Command authors to control their own completion or simply use
// the default. It also allows them to change the default by assigning
// to the package cmdbox.CompFunc before execution.
func (x *Command) Complete() {
	matches := []string{}
	switch {
	case x.CompFunc != nil:
		matches = x.CompFunc(x)
	case CompFunc != nil:
		matches = CompFunc(x)
	}
	for _, m := range matches {
		fmt.Println(m)
	}
}

func (x *Command) Title() string {
	summary := fmt.String(x.Summary)
	if len(summary) > 0 {
		return x.Name + " - " + summary
	}
	return x.Name
}

// Unimplemented calls Unimplemented passing the name of the command.
// Useful for temporarily notifying users of commands in beta that
// something has not yet been implemented.
func (x *Command) Unimplemented() error { return Unimplemented(x.Name) }

// NameFromSig returns the name from a Command signature. See
// Command.New.
func NameFromSig(sig string) string {
	all := strings.Split(sig, "|")
	return all[len(all)-1]
}
