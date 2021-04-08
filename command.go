package cmdbox

import (
	"encoding/json"
	_fmt "fmt"
	"strings"

	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/util"
)

// Command structs encapsulate the documentation and method of the
// command. This is the most important structure of CmdBox. Most of the
// fields can be set to util.Stringer which can be anything that returns
// a string, or is a string. This allows values to be determined at
// init() (runtime).  The Method contains the code to implement the
// command. And Completion, if assigned, till override the default tab
// completion behavior.
type Command struct {
	Name        string                   // <= 14 recommended
	Summary     string                   // < 65 recommended
	Version     util.Stringer            // semantic version (v0.1.3)
	Usage       util.Stringer            // following docopt syntax
	Description util.Stringer            // long form
	Examples    util.Stringer            // long form
	SeeAlso     util.Stringer            // links, other commands, etc.
	Author      util.Stringer            // email format, commas between
	Git         util.Stringer            // same as Go
	Issues      util.Stringer            // full web URL
	Copyright   util.Stringer            // legal copyright statement
	License     util.Stringer            // released under license(s)
	Other       map[string]util.Stringer // other (custom) doc sections

	Method     func(args []string) error
	Parameters util.Stringer
	Completion func(compline string) []string

	subcommands []string // Subcommands()
	Default     string   // default subcommand
}

// New initializes a new Command with subcommands (adding them to the
// internal subcommand index) and returns a pointer to the Command. Note
// that the subcommands are *not* added to the internal Command Index.
// They are saved as a list within the Command as Subcommands.
func New(name string, subcmds ...string) *Command {
	c := new(Command)
	c.Name = name
	c.Other = map[string]util.Stringer{}
	c.subcommands = []string{}
	if len(subcmds) > 0 {
		c.Add(subcmds...)
	}
	commands[name] = c
	return c
}

// Hidden returns true if the command name begins with underscore ('_').
func (c *Command) Hidden() bool {
	return c.Name[0] == '_'
}

func (c *Command) vsubcommands() []*Command {
	cmds := []*Command{}
	for _, name := range c.subcommands {
		if name[0] == '_' {
			continue
		}
		if command, has := commands[name]; has {
			cmds = append(cmds, command)
		}
	}
	return cmds
}

// SprintUsage returns a usage string (without emphasis) that includes a bold
// Command.Name for each line along with the main command usage and each individual
// SprintUsage of every subcommand that is not hidden. If indentation is needed this
// can be passed to Indent(). To replace emphasis with terminal escapes for printing
// to colored terminals pass to fmt.Emphasize().
func (c *Command) SprintUsage() string {
	buf := ""
	if c.Usage != nil {
		buf += "**" + c.Name + "** " + strings.TrimSpace(util.String(c.Usage)) + "\n"
	}
	for _, subcmd := range c.vsubcommands() {
		buf += "**" + c.Name + "** " + subcmd.SprintUsage()
	}
	if len(buf) > 0 {
		return buf
	}
	return "**" + c.Name + "**"
}

// SprintCommandSummaries returns a printable string that includes a bold
// Command.Name for each line along with the summary string, if any, for that
// subcommand. This is helpful when creating custom builtin help commands.
func (c *Command) SprintCommandSummaries() string {
	buf := ""
	for _, subcmd := range c.vsubcommands() {
		buf += _fmt.Sprintf("%-14v %v\n", "**"+subcmd.Name+"**", strings.TrimSpace(subcmd.Summary))
	}
	return buf
}

// MarshalJSON fulfills the Go JSON marshalling requirements and is called by
// String. Empty values are always omitted. Strings are trimmed and long
// strings such as Description and Examples are written as basic Markdown
// (which any Markdown engine will able to render. See fmt.Emph() and fmt.Plain() for
// specifics).
func (c *Command) MarshalJSON() ([]byte, error) {
	s := make(map[string]interface{})
	s["Name"] = c.Name
	var buf string

	if c.Summary != "" {
		s["Summary"] = strings.TrimSpace(c.Summary)
	}

	buf = strings.TrimSpace(util.String(c.Version))
	if buf != "" {
		s["Version"] = buf
	}

	buf = strings.TrimSpace(util.String(c.Usage))
	if buf != "" {
		s["Usage"] = buf
	}

	buf = fmt.Emph(util.String(c.Description), 0, -1)
	if buf != "" {
		s["Description"] = buf
	}

	buf = fmt.Emph(util.String(c.Examples), 0, -1)
	if buf != "" {
		s["Examples"] = buf
	}

	buf = fmt.Emph(util.String(c.SeeAlso), 0, -1)
	if buf != "" {
		s["SeeAlso"] = buf
	}

	buf = fmt.Emph(util.String(c.Author), 0, -1)
	if buf != "" {
		s["Author"] = buf
	}

	buf = strings.TrimSpace(util.String(c.Git))
	if buf != "" {
		s["Git"] = buf
	}

	buf = strings.TrimSpace(util.String(c.Issues))
	if buf != "" {
		s["Issues"] = buf
	}

	buf = strings.TrimSpace(util.String(c.Copyright))
	if buf != "" {
		s["Copyright"] = buf
	}

	buf = strings.TrimSpace(util.String(c.License))
	if buf != "" {
		s["License"] = buf
	}

	buf = strings.TrimSpace(util.String(c.Default))
	if buf != "" {
		s["Default"] = buf
	}

	if len(c.subcommands) > 0 {
		s["Subcommands"] = c.subcommands
	}

	// add custom (other) sections to docs
	for k, v := range c.Other {
		s[k] = util.String(v)
	}

	return json.Marshal(s)
}

// Fulfills the fmt.Stinger interface rendering a Command as a JSON
// string.
func (c Command) String() string {
	return util.ConvertToJSON(c)
}

// VersionLine returns a single line with the combined values of the Name,
// Version, Copyright, and License. If Version is empty or nil an empty string
// is returned instead. VersionLine() is used by the version builtin command
// to aggregate all the version information into a single output.
func (c *Command) VersionLine() string {
	version := util.String(c.Version)
	if version == "" || c.Name == "" {
		return ""
	}
	copyright := util.String(c.Copyright)
	license := util.String(c.License)
	buf := c.Name + " " + version
	if copyright != "" {
		buf += " " + copyright
	}
	if license != "" {
		buf += " (" + license + ")"
	}
	return buf
}

// Has looks for the named Subcommand.
func (c *Command) Has(name string) bool {
	for _, sc := range c.subcommands {
		if sc == name {
			return true
		}
	}
	return false
}

// Add adds new subcommands by name skipping any it already has. It is
// up to developers to ensure that the named subcommand has been added
// to the internal package index with New(). If any name contains a bar
// (|) then it will be split with the last item assumed to be the actual
// name and the first elements considered subcommand aliases (which are
// also added to the internal Subcommands).
func (c *Command) Add(names ...string) {
	for _, name := range names {
		if strings.ContainsRune(name, '|') {
			for _, n := range strings.Split(name, "|") {
				c.Add(n)
			}
			return
		}
		if !c.Has(name) {
			c.subcommands = append(c.subcommands, name)
		}
	}
}

// Subcommands returns the subcommands added with Add().
func (c *Command) Subcommands() []string {
	return c.subcommands
}

// SubcommandUsage returns the Usage strings for each Subcommand. This
// is useful when creating usages that have additional notes or formatting
// when it is desirable to loop through the subcommand usage strings. The
// order of usage strings is gauranteed to match the order of Subcommands()
// even if the usage for a particular subcommand is empty.
func (c *Command) SubcommandUsage() []string {
	usages := []string{}
	for _, name := range c.subcommands {
		usage := commands[name].Usage
		usages = append(usages, util.String(usage))
	}
	return usages
}

func (c *Command) UsageError() error {
	return _fmt.Errorf(fmt.Emphasize(strings.TrimSpace("**usage:** " + c.SprintUsage())))
}

// Complete prints the completion replies for the current context (See
// Programmble Completion in the bash man page.) The line is passed exactly as
// detected leaving the maximum flexibility for parsing and matching up to the
// Completion function.  The Completion method will be delegated if defined.
// Otherwise, the Subcommands are used to provide traditional prefix
// completion recursively.
func (c *Command) Complete(compline string) {
	if c.Completion != nil {
		for _, name := range c.Completion(compline) {
			fmt.Println(name)
		}
		return
	}
	words := strings.Split(strings.TrimSpace(compline), " ")
	if len(words) >= 2 {
		name := words[len(words)-2]
		complete := words[len(words)-1]
		if c.Name != name {
			if subcmd, has := commands[name]; has {
				compline = strings.Join(words[len(words)-2:], " ")
				subcmd.Complete(compline)
			}
			return
		}
		for _, subname := range c.subcommands {
			if subname == complete {
				if subcmd, has := commands[subname]; has {
					compline = strings.Join(words[len(words)-1:], " ")
					subcmd.Complete(compline)
				}
				return
			}
			if strings.HasPrefix(subname, complete) {
				fmt.Println(subname)
			}
		}
		if c.Parameters != nil {
			for _, param := range strings.Split(util.String(c.Parameters), " ") {
				if complete != param && strings.HasPrefix(param, complete) {
					fmt.Println(param)
				}
			}
		}
		return
	}
	for _, subname := range c.subcommands {
		if subname[0] != '_' {
			fmt.Println(subname)
		}
	}
	for _, param := range strings.Split(util.String(c.Parameters), " ") {
		fmt.Println(param)
	}

}

func (c *Command) Title() string {
	if len(c.Summary) > 0 {
		return c.Name + " - " + c.Summary
	}
	return c.Name
}

// Call calls its own Method or delegates to one of the Command's
// subcommands. If a Default has been set and the first argument does
// not appear to be a subcommand then delegate to Default subcommand by
// name.
func (c *Command) Call(args []string) error {
	if c.Method == nil && len(args) > 0 {
		subcmd := args[0]
		for _, name := range c.subcommands {
			if name == subcmd {
				if command, has := commands[name]; has {
					return command.Call(args[1:])
				}
				return Unimplemented(name)
			}
		}
	}
	if c.Default != "" {
		if command, has := commands[c.Default]; has {
			return command.Call(args)
		}
		return Unimplemented(c.Default)
	}
	if c.Method == nil {
		return c.UsageError()
	}
	return c.Method(args)
}

// Unimplemented calls Unimplemented passing the name of the command. Useful
// for temporarily notifying users of commands in beta that something has not
// yet been implemented.
func (c *Command) Unimplemented() error {
	return Unimplemented(c.Name)
}
