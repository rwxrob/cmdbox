package cmdbox

import (
	"sort"
	"strings"

	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/term"
	"github.com/rwxrob/cmdbox/term/esc"
)

// WARNING: The use of the internal commands directly is because of the
// builtin nature of these specific commands. Commands should generally
// use a read-only call to Index() instead to do similar things.

// Builtins are subcommands that are added to every Main command when
// Execute is called. This can be prevented by setting OmitBuiltins to false.
//
// Most of the builtins are hidden (beginning with underscore '_') but the
// following are so standardized they are included by default:
//
// * help [subcmd]  - very long, formatted documentation
// * version - version, copyright, license, authors, git source
//
// Any of these can be overridden by command authors simply by naming their
// own version the same. This may be desirable when creating commands in
// other languages although keeping these standard English names is
// strongly recommended due to their ubiquitous usage.
//
// The following are hidden but can be promoted by encapsulating them in
// other subcommands each with its own file, name, title, and documentation:
func Builtins() []string {
	return builtins
}

var builtins []string

func init() {

	allnames := func() []string {
		names := []string{}
		for name, _ := range commands {
			names = append(names, name)
		}
		sort.Strings(names)
		return names
	}

	// seq scan is just fine for this scale
	notbuiltin := func(name string) bool {
		for _, n := range builtins {
			if n == name {
				return false
			}
		}
		return true
	}

	names := func() []string {
		alln := allnames()
		justnames := []string{}
		for _, name := range alln {
			if notbuiltin(name) {
				justnames = append(justnames, name)
			}
		}
		return justnames
	}

	// ----------------------------------------------------------------

	_version := New("_cmdversion")
	_version.Summary = `print the cmd package version`
	_version.Method = func(ignored []string) error { fmt.Println(Version); return nil }
	builtins = append(builtins, "_cmdversion")

	_builtins := New("_builtins")
	_builtins.Summary = `list all cmd package builtins names and summaries`
	_builtins.Method = func(ignored []string) error {
		sort.Strings(builtins)
		for _, name := range builtins {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(commands[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_builtins")

	// ----------------------------------------------------------------

	_complete := New("_complete")
	_complete.Summary = `force completion context`
	_complete.Method = func(args []string) error {
		words := []string{Main.Name}
		words = append(words, args...)
		comp.Line = strings.Join(words, " ")
		Complete()
		return nil
	}
	builtins = append(builtins, "_complete")

	// ----------------------------------------------------------------

	_index := New("_index")
	_index.Summary = `list all names and summaries from cmd package index`
	_index.Method = func(ignored []string) error {
		for _, name := range allnames() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(commands[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_index")

	// ----------------------------------------------------------------

	_names := New("_names")
	_names.Summary = `list names, main first`
	_names.Method = func(ignored []string) error {
		fmt.Println(Main.Name)
		for _, name := range names() {
			if name != Main.Name {
				fmt.Println(name)
			}
		}
		return nil
	}
	builtins = append(builtins, "_names")

	// ----------------------------------------------------------------

	_summaries := New("_summaries")
	_summaries.Summary = `list names and summaries`
	_summaries.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(commands[name].Summary))
		}
		return nil
	}
	builtins = append(builtins, "_summaries")

	// ----------------------------------------------------------------

	_versions := New("_versions")
	_versions.Summary = `list names and versions`
	_versions.Method = func(args []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(commands[name].Version)))
		}
		return nil
	}
	builtins = append(builtins, "_versions")

	// ----------------------------------------------------------------

	_copyrights := New("_copyrights")
	_copyrights.Summary = `list names and copyrights`
	_copyrights.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(string(fmt.String(commands[name].Copyright))))
		}
		return nil
	}
	builtins = append(builtins, "_copyrights")

	// ----------------------------------------------------------------

	_licenses := New("_licenses")
	_licenses.Summary = `list names and licenses`
	_licenses.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(commands[name].License)))
		}
		return nil
	}
	builtins = append(builtins, "_licenses")

	// ----------------------------------------------------------------

	_authors := New("_authors")
	_authors.Summary = `list names and authors`
	_authors.Method = func(ignored []string) error {
		for _, name := range names() {
			author := commands[name].Author
			if author == nil {
				author = ""
			}
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(author)))
		}
		return nil
	}
	builtins = append(builtins, "_authors")

	// ----------------------------------------------------------------

	_gits := New("_gits")
	_gits.Summary = `list names and git source repos`
	_gits.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(commands[name].Git)))
		}
		return nil
	}
	builtins = append(builtins, "_gits")

	// ----------------------------------------------------------------

	_issues := New("_issues")
	_issues.Summary = `list names and issue reporting URLs`
	_issues.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(commands[name].Issues)))
		}
		return nil
	}
	builtins = append(builtins, "_issues")

	// ----------------------------------------------------------------

	_usages := New("_usages")
	_usages.Summary = `list names and usages`
	_usages.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("%-14v %v\n", name, strings.TrimSpace(fmt.String(commands[name].Usage)))
		}
		return nil
	}
	builtins = append(builtins, "_usages")

	// ----------------------------------------------------------------

	_desc := New("_descriptions")
	_desc.Summary = `list names and descriptions`
	_desc.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("DESCRIPTION %v\n\n", name)
			fmt.Println(fmt.Plain(fmt.String(commands[name].Description), 4, int(term.WinSize.Col)))
			fmt.Println()
		}
		return nil
	}
	builtins = append(builtins, "_descriptions")

	// ----------------------------------------------------------------

	_examples := New("_examples")
	_examples.Summary = `list names and examples`
	_examples.Method = func(ignored []string) error {
		for _, name := range names() {
			fmt.Print("EXAMPLE %v\n\n", name)
			fmt.Print(fmt.Sprint(fmt.Plain(fmt.String(commands[name].Examples), 4, int(term.WinSize.Col))))
			fmt.Print("\n\n")
		}
		return nil
	}
	builtins = append(builtins, "_examples")

	// ----------------------------------------------------------------

	_help_json := New("_help_json")
	_help_json.Summary = `dump help documentation as JSON`
	_help_json.Method = func(args []string) error {
		fmt.Println(JSON())
		return nil
	}
	builtins = append(builtins, "_help_json")

	// ----------------------------------------------------------------

	help := New("help")
	help.Summary = "display detailed help documentation"
	help.Method = func(args []string) error {
		c := Main
		if len(args) > 0 && Has(args[0]) {
			c = commands[args[0]]
		}
		output := fmt.Sprint(
			fmt.TopTitle(
				esc.B+c.Name, esc.X+"DOCUMENTATION", esc.B+c.Name,
				int(term.WinSize.Col),
			)+"\n\n",
			esc.B+"NAME"+esc.X+"\n",
		)
		fmt.Println(output)
		return nil

		/*
				esc.B + "NAME" + esc.X + "\n" +
				fmt.Emph(c.Title(), 4, int(term.WinSize.Col)) + "\n\n" +
				esc.B + "USAGE" + esc.X + "\n" +
				fmt.Indent(fmt.Emphasize(c.SprintUsage()), 4) + "\n\n")
			if len(c.vsubcommands()) > 0 {
				output += esc.B + "COMMANDS" + esc.X + "\n" +
					fmt.Indent(fmt.Emphasize(c.SprintCommandSummaries()), 4) + "\n\n"
			}
			output += esc.B + "DESCRIPTION" + esc.X + "\n" +
				fmt.Emph(fmt.String(c.Description), 4, int(term.WinSize.Col)) + "\n\n"

			// TODO finish output
			if fmt.PagedOut {
				fmt.PrintPaged(output, "")
			} else {
				fmt.Print(output)
			}
			return nil
		*/
	}
	builtins = append(builtins, "help")

	// ----------------------------------------------------------------

	version := New("version")
	version.Summary = `display version, author, and legal information`
	version.Method = func(args []string) error {
		if len(args) > 0 {
			v := commands[args[0]].VersionLine()
			if v != "" {
				fmt.Println(v)
			}
			return nil
		}
		vl := Main.VersionLine()
		if vl == "" {
			return nil
		}
		fmt.Println(vl)
		for _, name := range names() {
			if name == Main.Name {
				continue
			}
			line := commands[name].VersionLine()
			if line == "" {
				continue
			}
			fmt.Println(line)
		}
		return nil
	}
	builtins = append(builtins, "version")

	// ----------------------------------------------------------------

	_usage := New("_usage")
	_usage.Summary = "display usage summaries"
	_usage.Method = func(ignored []string) error {
		fmt.Print(Main.SprintUsage())
		return nil
	}
	builtins = append(builtins, "_usage")

}
