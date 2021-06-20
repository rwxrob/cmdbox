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
	"sort"

	"github.com/rwxrob/cmdbox/util"
)

// Map is a high-level type used to contain string data for both keys
// and values Developers can use the Map type in their own modules as
// a convenience since requires less declaration (Map{"foo":"bar"}). As
// with any Go map, however, reading and writing to a Map is not safe
// for concurrency without additional locking (usually contained within
// the parent structure which contains a Map). See Command.Commands as
// well.
type Map map[string]string

// Names returns a sorted list of the command names only. Also see
// Aliases and Keys
func (m Map) Names() []string {
	return util.UniqStrMapVals(m)
}

// Keys returns a sorted list of all possible commands, actions and
// aliases from the Commands. Note that this does not include any Params
// and therefore is not suitable by itself for producing a full list for
// completion. This is just a list of everything that is associated,
// directly or indirectly, with a Command in the internal register. Also
// see Names and Aliases.
func (m Map) Keys() []string {
	keys := make([]string, len(m))
	var i int
	for k, _ := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// Aliases returns only the keys that are not identical to their value.
func (m Map) Aliases() []string {
	a := []string{}
	for k, v := range m {
		if k == v {
			continue
		}
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}

// JSON is shortcut for json.Marshal(m). See util.ToJSON.
func (m Map) JSON() string { return util.ToJSON(m) }

// String fullfills fmt.Stringer interface as JSON.
func (m Map) String() string { return util.ToJSON(m) }

// YAML is shortcut for yaml.Marshal(m). See util.ToYAML.
func (m Map) YAML() string { return util.ToYAML(m) }

// Print outputs as YAML (nice when testing).
func (m Map) Print() { fmt.Println(util.ToYAML(m)) }
