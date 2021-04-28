package cmdbox

import (
	"sort"
	"strings"

	"github.com/rwxrob/cmdbox/comp"
)

// CompFunc is assigned CompleteCommand by default but can be assigned
// any valid comp.Func to override it. This function is called to
// perform completion for any Command that does not implement its own
// Command.CompFunc.
var CompFunc = comp.Func(CompleteCommand)

// CompleteCommand takes a pointer to a cmdbox.Command returning a list
// of lexigraphically sorted combination of strings from
// Command.Commands that are found in the cmdbox.Register and
// Command.Params that match the current completion context. Returns an
// empty list if anything fails. Note that no assertion that the
// specified command names exist in the cmdbox.Register. See the
// Command.Complete method and cmdbox/comp package.
func CompleteCommand(a ...interface{}) []string {
	rv := []string{}

	if len(a) == 0 || comp.Line() == "" {
		return rv
	}

	x := a[0].(*Command)
	word := comp.Word()

	for k, _ := range x.Commands {
		if word == " " || strings.HasPrefix(k, word) {
			rv = append(rv, k)
		}
	}

	for _, k := range x.Params {
		if word == " " || strings.HasPrefix(k, word) {
			rv = append(rv, k)
		}
	}

	sort.Strings(rv)
	return rv
}
