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
	Name        string                 // <= 14 recommended
	Summary     string                 // < 65 recommended
	Version     interface{}            // semantic version (v0.1.3)
	Usage       interface{}            // following docopt syntax
	Description interface{}            // long form
	Examples    interface{}            // long form
	SeeAlso     interface{}            // links, other commands, etc.
	Author      interface{}            // email format, commas between
	Git         interface{}            // same as Go
	Issues      interface{}            // full web URL
	Copyright   interface{}            // legal copyright statement
	License     interface{}            // released under license(s)
	Other       map[string]interface{} // other (custom) doc sections

	Method     func(args []string) error
	Parameters []string
	Completion func(compline string) []string

	subcommands []string    // Subcommands()
	Default     interface{} // default subcommand
}

// New initializes a new Command and returns a pointer to it. An
// optional list of subcommand name strings can be passed as arguments
// and will be added with c.Add(), the first being assigned to
// c.Default. Note that calling New is the *only* way to add a new
// Command to the protected commands registry (which can be read with
// the cmdbox.Commands() function). If a name conflicts with one that
// has already been added an underscore (_) has added.  This can be
// renamed with cmdbox.Rename() or removed with cmdbox.Remove().
func New(name string, a ...string) *Command {
	c := new(Command)
	c.Name = name
	c.subcommands = []string{} // used by add
	commands[c.Name] = c
	if len(a) > 1 {
		c.Add(a...)
		c.Default = a[0]
	}
	c.Other = map[string]interface{}{}
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
		buf += "**" + c.Name + "** " + strings.TrimSpace(fmt.String(c.Usage)) + "\n"
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

	buf = strings.TrimSpace(fmt.String(c.Version))
	if buf != "" {
		s["Version"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.Usage))
	if buf != "" {
		s["Usage"] = buf
	}

	buf = fmt.Emph(fmt.String(c.Description), 0, -1)
	if buf != "" {
		s["Description"] = buf
	}

	buf = fmt.Emph(fmt.String(c.Examples), 0, -1)
	if buf != "" {
		s["Examples"] = buf
	}

	buf = fmt.Emph(fmt.String(c.SeeAlso), 0, -1)
	if buf != "" {
		s["SeeAlso"] = buf
	}

	buf = fmt.Emph(fmt.String(c.Author), 0, -1)
	if buf != "" {
		s["Author"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.Git))
	if buf != "" {
		s["Git"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.Issues))
	if buf != "" {
		s["Issues"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.Copyright))
	if buf != "" {
		s["Copyright"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.License))
	if buf != "" {
		s["License"] = buf
	}

	buf = strings.TrimSpace(fmt.String(c.Default))
	if buf != "" {
		s["Default"] = buf
	}

	if len(c.subcommands) > 0 {
		s["Subcommands"] = c.subcommands
	}

	// add custom (other) sections to docs
	for k, v := range c.Other {
		s[k] = fmt.String(v)
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
	version := fmt.String(c.Version)
	if version == "" || c.Name == "" {
		return ""
	}
	copyright := fmt.String(c.Copyright)
	license := fmt.String(c.License)
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
		usages = append(usages, fmt.String(usage))
	}
	return usages
}

func (c *Command) UsageError() error {
	return _fmt.Errorf(fmt.Emphasize(strings.TrimSpace("**usage:** " + c.SprintUsage())))
}

// Complete prints the completion replies for the current context (See
// Programmble Completion in the bash man page.) The line is passed
// exactly as detected leaving the maximum flexibility for parsing and
// matching up to the Completion function.  The Completion method will
// be delegated if defined.  Otherwise, the subcommands (provided when
// New() was called) are used to provide traditional prefix completion
// recursively.
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
			for _, param := range strings.Split(fmt.String(c.Parameters), " ") {
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
	for _, param := range strings.Split(fmt.String(c.Parameters), " ") {
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
	d := fmt.String(c.Default)
	if d != "" {
		if command, has := commands[d]; has {
			return command.Call(args)
		}
		return Unimplemented(d)
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
