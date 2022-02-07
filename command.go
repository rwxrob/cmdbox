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

package cmdbox

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rwxrob/cmdbox/util"
	"github.com/rwxrob/cmdbox/valid"
)

// Command contains a Method or delegates to  one or more other Commands
// by name. Typically a Command is created within an init() function by
// calling cmdbox.New:
//
//     import "github.com/rwxrob/cmdbox"
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
// Tab completion rules default to the list of Commands and Params,
// but can be overriden per Command by defining and assigning an
// anonymous closure function to the CompFunc field (see CompFunc type).
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
// Conventions for Text Field Values
//
// When providing the text of the different fields of the the Command
// struct please keep the following conventional considerations in mind:
//
// * Imagine your text being spoken by a conversational assistant
// * Avoid the use of acronymns and excessive punctuation
// * Resist the temptation to assign const and vars outside init
//
// Fields such as Usage are obvious exceptions to these conventions.
//
// Description
//
// The Description contains a long Command description and can span
// multiple paragraphs, even pages. Use of backtick quotes (`) is
// preferred and text may be indended and wrapped anywhere. Paragraphs
// are recognized by a blank line between them. Keep Go documentation
// style guidelines in mind when writing them (headers on single line,
// four space indent for verbatim text, etc.)
//
// Params
//
// The Params list is for completion as well, specifically for things
// that are neither Commands nor actions to be handled by the Method but
// would be nice to have included in completion. Unlike Commands and
// Aliases, Params do not need to be valid names (see valid package).
// This allows them to be used for things such as default numeric values
// that may begin with a number or dash and other things that contain
// punctuation.  These should not, however, be used to bypass the core
// requirement for speakable Commands and Aliases, and whenever
// possible, arguments as well. (For more complex completion, assign
// a custom x.CompFunc function.)
//
// Examples
//
// For examples of different Command structs search on GitHub for any
// project beginning with cmdbox- such as the following:
//
// * https://github.com/rwxrob/cmdbox-greet
// * httpe://github.com/rwxrob/cmdbox-pomo
//
type Command struct {
	Name        string          `json:"name,omitempty" yaml:",omitempty"`
	Summary     string          `json:"summary,omitempty" yaml:",omitempty"`
	Usage       string          `json:"usage,omitempty" yaml:",omitempty"`
	Version     string          `json:"version,omitempty" yaml:",omitempty"`
	Copyright   string          `json:"copyright,omitempty" yaml:",omitempty"`
	License     string          `json:"license,omitempty" yaml:",omitempty"`
	Description string          `json:"description,omitempty" yaml:",omitempty"`
	Site        string          `json:"git,omitempty" yaml:",omitempty"`
	Source      string          `json:"git,omitempty" yaml:",omitempty"`
	Issues      string          `json:"issues,omitempty" yaml:",omitempty"`
	Commands    *util.StringMap `json:"commands,omitempty" yaml:",omitempty"`
	Params      []string        `json:"params,omitempty" yaml:",omitempty"`
	Default     string          `json:"default,omitempty" yaml:",omitempty"`
	// Title()
	// Legal()
	CompFunc   CompFunc `json:"-" yaml:"-"`
	Caller     *Command `json:"-" yaml:"-"`
	Method     Method   `json:"-" yaml:"-"`
	sync.Mutex `json:"-" yaml:"-"`
}

// Method defines functions for use primarily as Command.Method values.
type Method func(args ...string) error

// NewCommand returns pointer to new initialized Command. See the New
// package function instead for creating a new Command that is also
// added to the Register. Minimal validation is done on the name and all
// arguments (subcommands, actions) to ensure a consistent user
// experience for all CmdBox commands (see valid subpackage for more).
// The Usage string is set to the default with UpdateUsage.
// Since calling NewCommand involves critical syntax checks a panic is
// thrown if invalid.
func NewCommand(name string, a ...string) *Command {
	x := new(Command)

	// panic unless valid command name
	if !valid.Name(name) && name[len(name)-1] != '_' {
		SyntaxErrorPanic(fmt.Sprintf(m_invalid_name, name))
	}

	// initialize command with internal h|help and version
	x.Name = name
	x.Commands = util.NewStringMap()
	x.Default = "help"

	// add any subcommands
	if len(a) > 0 {
		x.Add(a...)
	}

	x.UpdateUsage()

	return x
}

// Title returns a dynamic field of Name and Summary combined (if
// exists).
func (x *Command) Title() string {
	switch {
	case len(x.Summary) > 0 && len(x.Version) > 0:
		return x.Name + " (" + x.Version + ")" + " - " + x.Summary
	case len(x.Summary) > 0:
		return x.Name + " - " + x.Summary
	default:
		return x.Name
	}
}

// Legal returns a single line with the combined values of the
// Name, Version, Copyright, and License. If Version is empty or nil an
// empty string is returned instead. Legal() is used by the
// version builtin command to aggregate all the version information into
// a single output.
func (x *Command) Legal() string {
	switch {
	case len(x.Copyright) > 0 && len(x.License) > 0 && len(x.Version) > 0:
		return x.Name + " " + x.Version + "\n" +
			x.Copyright + "\nLicense " + x.License
	case len(x.Copyright) > 0 && len(x.License) > 0:
		return x.Name + "\n" + x.Copyright + "\nLicense " + x.License
	case len(x.Copyright) > 0:
		return x.Name + "\n" + x.Copyright
	default:
		return ""
	}
}

// UpdateUsage will set x.Usage to the default, which is all of the
// Commands joined with bar (|) and wrapped in brackets ([]).
func (x *Command) UpdateUsage() {
	names := x.Commands.Keys()
	if len(names) > 0 {
		x.Usage = "[" + strings.Join(names, "|") + "]"
	}
}

// JSON is shortcut for json.Marshal(x). See util.ToJSON.
func (x Command) JSON() string { return util.ToJSON(x) }

// YAML is shortcut for yaml.Marshal(x). See util.ToYAML.
func (x Command) YAML() string { return util.ToYAML(x) }

// Print outputs as YAML (nice when testing).
func (x Command) Print() { fmt.Println(util.ToYAML(x)) }

// String fullfills fmt.Stringer interface as JSON.
func (x Command) String() string { return util.ToJSON(x) }

// Add adds the list of Command signatures passed. A command signature
// consists of one or more more aliases separated by a bar (|) with the
// final word being the name of the actual Command.  Aliases are
// a useful way to provide shortcuts when tab completion is not
// available and should generally be considered for every Command.
// Single letter aliases are common and encouraged.
//
// Note that Add does not validate inclusion in the internal Register
// (Reg) since in many cases there may not yet be a Register entry, and
// in the case of actions handled entirely by the Command itself there
// never will be.  See Command.Commands and Command.Run.
//
func (x *Command) Add(sigs ...string) {
	defer x.Unlock()
	x.Lock()
	if x.Commands == nil {
		x.Commands = util.NewStringMap()
	}
	for _, sig := range sigs {
		aliases := strings.Split(sig, "|")
		name := aliases[len(aliases)-1]
		if !valid.Name(name) {
			panic(m_invalid_name)
		}
		x.Commands.Set(name, name)
		for _, alias := range aliases {
			if !valid.Name(alias) {
				panic(m_invalid_name)
			}
			x.Commands.Set(alias, name)
		}
	}
}

// Complete prints the possible strings based on the current Command and
// completion context. If the Commands CompFunc has been assigned (not
// nil) it is called and passed its own pointer. If CompFunc has not
// been assigned (is nil) then cmdbox.DefaultComplete is called instead.
// This allows Command authors to control their own completion or simply
// use the default. It also allows changing the default by assigning to
// the package cmdbox.DefaultComplete before calling cmdbox.Execute.
func (x *Command) Complete() {
	matches := []string{}
	switch {
	case x.CompFunc != nil:
		matches = x.CompFunc(x)
	case DefaultComplete != nil:
		matches = DefaultComplete(x)
	}
	for _, m := range matches {
		fmt.Println(m)
	}
}

// ------------------------------ errors ------------------------------

// Unimplemented is a convenience method that delegates calls to
// cmdbox.Unimplemented.
func (x *Command) Unimplemented(a string) error { return Unimplemented(a) }

// UsageError is a convenience method that delegates calls to
// cmdbox.UsageError.
func (x *Command) UsageError() error { return UsageError(x) }

// MissingArg returns cmdbox.MissingArg
func (x *Command) MissingArg(a string) error {
	return MissingArg(a)
}

// UnexpectedArg returns cmdbox.UnexpectedArg
func (x *Command) UnexpectedArg(a string) error {
	return UnexpectedArg(a)
}

// ------------------------------- help -------------------------------

// Help returns a formatted string suitable for printing either to
// a file or to an interactive terminal. For a more structured form of
// the same information see YAML, JSON, Print, and PrintHelp.
func (x Command) Help() string {
	var buf string

	buf += util.Emph("**NAME**", 0, -1) + "\n       " + x.Title() + "\n\n"
	buf += util.Emph("**SYNOPSIS**", 0, -1) + "\n       " + x.Name + " " + x.Usage + "\n\n"

	if len(x.Commands.M) > 0 {
		buf += util.Emph("**COMMANDS**", 0, -1) + "\n" + x.Titles(7, 20) + "\n\n"
	}

	if len(x.Description) > 0 {
		buf +=
			util.Emph("**DESCRIPTION**", 0, -1) + "\n" +
				util.Emph(x.Description, 7, 65) + "\n\n"
	}

	if x.Source != "" || x.Issues != "" || x.Site != "" {

		buf += util.Emph("**LINKS**", 0, -1) + "\n"

		if x.Site != "" {
			buf += "       Site:   " + x.Site + "\n"
		}

		if x.Source != "" {
			buf += "       Source: " + x.Source + "\n"
		}

		if x.Issues != "" {
			buf += "       Issues: " + x.Issues + "\n"
		}

		buf += "\n"

	}

	if x.Copyright != "" {
		buf += util.Emph("**LEGAL**", 0, -1) + "\n" + util.Indent(x.Legal(), 7) + "\n\n"
	}

	return buf

}

// PrintHelp simply prints what Help returns.
func (x Command) PrintHelp() { fmt.Print(x.Help()) }

// ResolveDelegate returns a Command pointer looked up from the internal
// register based on the following priority from more
//
//   1. x.Caller.Name + " " + arg[0] as name
//   2. cmdbox.Main.Name + " " + arg[0] as name
//   3. arg[0] as name
//
// This is a specialized lookup for Commands that are designed to
// operate on other Commands in the registry. See Help and Legal for
// examples. See Resolve for simpler resolution of Commands.
//
func (x *Command) ResolveDelegate(args []string) *Command {
	var xx *Command

	if x.Caller != nil {
		xx = Get(x.Caller.Name + " " + args[0])
		if xx != nil {
			return xx
		}
	}

	xx = Get(Main.Name + " " + args[0])
	if xx != nil {
		return xx
	}

	return Get(args[0])

}

// Sigs returns a StringMap keyed to the Command.Names with
// signatures as values suitable for printing usage information.
//
func (x Command) Sigs() *util.StringMap {
	return x.Commands.KeysCombined("|")
}

// Titles returns a single string with the titles of each subcommand
// indented and with a maximum title signature length for justification.
//
func (x Command) Titles(indent, max int) string {
	buf := ""
	sigs := x.Sigs()
	_, lval := sigs.LongestValue()
	limit := len(lval)
	if max != 0 && max < limit {
		limit = max
	}
	for _, name := range x.Commands.Values() {
		summary := "<not yet implemented>"
		c := x.Resolve(name)
		if c != nil {
			summary = c.Summary
		}
		buf += fmt.Sprintf("%-"+fmt.Sprintf("%v", limit)+"v - %v\n",
			sigs.Get(name), summary)
	}
	return util.Indent(buf, indent)
}

// Resolve looks up the Command from the register based on the name
// passed. First it looks for a fully qualified entry in the register
// (x.Name + " " + name), then it just looks for the name alone. Returns
// nil if nothing found in the register. This method is particularly
// important because at init time when Commands are Added to the register
// their subcommands may not yet have been registered. Resolve allows
// this lookup to happen reliably later in runtime.
func (x Command) Resolve(name string) *Command {
	n := Reg.Get(x.Name + " " + name)
	if n == nil {
		n = Reg.Get(name)
	}
	return n
}
