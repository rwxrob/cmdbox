package cmdbox

import (
	"sort"

	"github.com/rwxrob/cmdbox/util"
)

// CommandMap contains the command names and aliases used for
// completion pointing to the command or action name to be used.
type CommandMap map[string]string

// Names returns a sorted list of the command names.
func (c CommandMap) Names() []string {
	return util.UniqStrMapVals(c)
}

// Aliases returns a sorted list of all possible commands, actions
// and aliases from the Commands CommandMap. Note that this does not
// include any Params and therefore is not suitable by itself for
// producing a full list for completion. This is just a list of
// everything that is associated, directly or indirectly, with a Command
// in the Register.
func (c CommandMap) Aliases() []string {
	aliases := make([]string, len(c))
	var i int
	for k, _ := range c {
		aliases[i] = k
		i++
	}
	sort.Strings(aliases)
	return aliases
}

// String fulfills the fmt.Stringer interface to print as JSON.
func (c CommandMap) String() string { return util.ConvertToJSON(c) }
