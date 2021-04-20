/*
Package cmdbox provides a lightweight commander focused on modern
human-computer interaction through terminal command-line interfaces
composed of modular subcommands with portable completion, and embedded,
dynamic documentation.
*/
package cmdbox

import (
	"os"
	"strings"

	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/util"
)

var Version string

// MultiName (which is true by default) indicates that the compiled
// executable may have several different names when invoked and, if so,
// to use that name as the name of the command to execute. If set to
// false the executable name will be entirely ignored. Executables are
// renamed by being copied or linked depending on the constraints of the
// host system.
var MultiName bool = true

// KeepAlive allows developers to stop Execute() from exiting. It should
// not be used for any purpose other than testing and should be kept out
// of any test examples.
var KeepAlive bool

// OmitBuiltins turns off the injection of the Builtin subcommands into
// the Main command when Execute is called. It can be assigned in any
// init() or from main() before calling Execute().
var OmitBuiltins bool = true

// OmitAllBuiltins prevents even the help and version builtins from
// being included (which is useful mostly for creating example tests).
var OmitAllBuiltins bool

// Args returns a reliable collection of arguments to the executable.
//
// WARNING: Although the first the element of os.Args is usually the binary
// of the compiled program executed it is never reliable and significantly
// differs depending on operating system and method of program execution.
// The first argument is therefore stripped completely leaving only the
// arguments to be processed. The cmd.Args package variable can also be
// set during testing to check cmd.Execute() behavior.
var Args = os.Args[1:]

// Main contains the main command passed to Execute to start the
// program. While it can be changed by Subcommands it usually should not be.
var Main *Command

// CompLineFix activates an attempt to correct the comp.Line to work best with
// completion. For example, when an executable that uses the cmd package is
// renamed, symbolicly linked, or called as a path and not just the single
// command name. True bydefault. Set to false to leave the comp.Line exactly as
// it is detected but note that depending on a specific form of
// comp.Line may not be consistent across operating systems. Also see
// util.MultiName.
var CompLineFix bool = true

// Complete calls complete on the Main command passing it comp.Line. No
// verification of Main's existence is checked. The comp.Line is always changed
// to match the actual name of the Main command even if the executable name has
// been changed or called as an alias. This ensures proper tab completion no
// matter what the actual executable is called.
func Complete() {
	if !CompLineFix {
		Main.Complete(comp.Line)
		return
	}
	i := strings.Index(comp.Line, " ")
	if i < 0 {
		Main.Complete(Main.Name)
	} else {
		Main.Complete(Main.Name + comp.Line[i:])
	}
}

// JSON returns a JSON representation of the state of the cmd package including
// the main command and all subcommands from the internal index. This can be
// useful when providing documentation in a structured data format that can be
// easily shared and rendered in different ways. The json builtin simply calls
// this and prints it. Empty values are always omitted. (See
// Command.MarshalJSON() as well.)
func JSON() string {
	s := make(map[string]interface{})
	s["PackageVersion"] = Version
	s["Main"] = Main.Name
	s["Register"] = Register
	return util.ConvertToJSON(s)
}
