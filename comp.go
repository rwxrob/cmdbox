package cmdbox

import (
	"sort"
	"strings"

	"github.com/rwxrob/cmdbox/comp"
)

// CompFunc is the main completion function (comp.Func) assigned by
// default to the cmdtab.Main Command if it doesn't already have one.
// The completion algorithm is as follows:
//
// * If comp.Args is empty return an empty slice.
//
// * If comp.Args has one item assume it is the name of the executable
//   being called and return all visible command names (keys) from the
//   Register.
//
// * If comp.Args has two items and the second item is *not* in the
//   cmdbox.Register then return all Visible command names from the
//   Register that also have that item as a prefix.
//
// * If comp.Args has two or more items and the second item is a command
//   name then delegate to the Complete method on that command.
//
// * If comp.Args has more than two items and the second is not
//   a command name then delegate to comp.FileDir to complete with
//   whatever is in the immediate directory.
//
// The []string is sorted before being returned in lexographical order.
//
// Note this function must remain in cmdbox package to avoid circular
// import cycles that would happen if in the comp subpackage with other
// cmdbox.CompFunc implementations.
//
func CompFunc() []string {
	m := []string{}
	args := comp.Args()
	switch {
	case len(args) < 1:
		return m
	case len(args) == 1:
		for k, _ := range Visible() {
			m = append(m, k)
		}
	case len(args) == 2:
		if c, has := Register[args[1]]; has {
			return c.Complete()
		}
		for k, _ := range Visible() {
			if strings.HasPrefix(k, args[1]) {
				m = append(m, k)
			}
		}
	default:
		return comp.FileDir()
	}
	sort.Strings(m)
	return m
}
