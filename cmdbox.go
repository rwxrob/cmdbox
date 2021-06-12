/*
Package cmdbox provides a lightweight commander focused on modern
human-computer interaction through terminal command-line interfaces
composed of modular subcommands with portable completion, and embedded,
dynamic documentation.
*/
package cmdbox

import (
	"github.com/rwxrob/cmdbox/fmt"
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

// Main contains the main command passed to Execute to start the
// program. While it can be changed by Subcommands it usually should not be.
var Main *Command

// JSON returns a JSON representation of the state of the cmd package including
// the main command and all subcommands from the internal index. This can be
// useful when providing documentation in a structured data format that can be
// easily shared and rendered in different ways. The json builtin simply calls
// this and prints it. Empty values are always omitted. (See
// Command.MarshalJSON() as well.)
func JSON() string {
	s := make(map[string]interface{})
	if Version != "" {
		s["PackageVersion"] = Version
	}
	if Main != nil {
		s["Main"] = Main.Name
	}
	s["Register"] = Register
	return util.ConvertToJSON(s)
}

// String returns Package metadata and Register as a JSON string.
func String() string { return JSON() }

// Print is shortcut for fmt.Println(cmdbox.String()).
func Print() { fmt.Println(String()) }

// Init resets the internal Register as if no CmdBox Command init
// function had been called. Package metadata is preserved.
func Init() { Register = map[string]*Command{} }

// Has returns true if Register has an entry that matches the key.
func Has(key string) bool { _, has := Register[key]; return has }
