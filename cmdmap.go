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

/* Package cmdbox is a multicall, modular commander with embedded tab
completion and locale-driven documentation, that prioritizes modern,
speakable human-computer interactions from the command line.
*/
package cmdbox

import (
	"fmt"
	"sort"
	"sync"

	"github.com/rwxrob/cmdbox/util"
)

// CommandMap encapsulates a map[string]*Command and embeds a sync.Mutex
// for locking making it safe for concurrency. The internal map is
// exported (M) in the event developers want more direct control without
// using the CommandMap methods that ensure concurrency safety. See the
// Register function for more.
type CommandMap struct {
	M map[string]*Command
	sync.Mutex
}

// NewCommandMap returns a new CommandMap with the internal
// map[string]*Command initialized.
func NewCommandMap() *CommandMap {
	m := new(CommandMap)
	m.Init()
	return m
}

// Init initializes (or re-initialized) the CommandMap deleting all its
// values (without changing its reference).
func (m *CommandMap) Init() {
	m.Lock()
	defer m.Unlock()
	if m.M == nil {
		m.M = make(map[string]*Command)
		return
	}
	for k := range m.M {
		delete(m.M, k)
	}
}

// Get returns a Command pointer by key name safe for concurrency.
// Returns nil if not found.
func (m *CommandMap) Get(key string) *Command {
	m.Lock()
	defer m.Unlock()
	if v, has := m.M[key]; has {
		return v
	}
	return nil
}

// Set set a value by key name safe in a way that is for concurrency.
func (m *CommandMap) Set(key string, val *Command) {
	defer m.Unlock()
	m.Lock()
	m.M[key] = val
}

// Delete removes one or more entries from the map in a way that is safe
// for concurrency.
func (m *CommandMap) Delete(keys ...string) {
	defer m.Unlock()
	m.Lock()
	for _, k := range keys {
		delete(m.M, k)
	}
}

// Rename renames a Command in the Register by adding the
// new name with the same *Command and deleting the old one. This is
// useful when a name conflict causes New to append and underscore (_)
// to the duplicate's name. Rename can be called from init() at any
// point after the duplicate has been added to resolve the conflict.
// Note the order of init() execution --- while predictable --- is not
// always apparent.  When in doubt do Rename from main() to be sure.
// Rename is safe for concurrency.
func (m CommandMap) Rename(from, to string) {
	defer m.Unlock()
	m.Lock()
	x, has := m.M[from]
	if !has {
		return
	}
	x.Name = to
	m.M[to] = x
	delete(m.M, from)
}

// ------------------------------ queries -----------------------------

// Names returns a sorted list of all Command names.
func (m CommandMap) Names() []string {
	m.Lock()
	defer m.Unlock()
	keys := make([]string, len(m.M))
	var i int
	for k, _ := range m.M {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// Dups returns key strings of duplicates (which can then be easily
// renamed). Keys are sorted in lexicographic order. See Rename.
func (m CommandMap) Dups() []string {
	defer m.Unlock()
	m.Lock()
	var keys []string
	for k, _ := range m.M {
		if k[len(k)-1] == '_' {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// Slice returns a slice of *Command pointers and fetched from the
// internal register that match the key names passed.  If an entry is
// not found it is simply skipped. Will return an empty slice if none
// found.
func (m CommandMap) Slice(names ...string) []*Command {
	defer m.Unlock()
	m.Lock()
	cmds := []*Command{}
	for _, name := range names {
		if x, has := m.M[name]; has {
			cmds = append(cmds, x)
		}
	}
	return cmds
}

// ---------------------------- marshaling ----------------------------

// JSON is shortcut for json.Marshal(m). See ToJSON.
func (m CommandMap) JSON() string { return util.ToJSON(m.M) }

// String fullfills fmt.Stringer interface as JSON.
func (m CommandMap) String() string { return util.ToJSON(m.M) }

// YAML is shortcut for yaml.Marshal(m). See ToYAML.
func (m CommandMap) YAML() string { return util.ToYAML(m.M) }

// Print outputs as YAML (nice when testing).
func (m CommandMap) Print() { fmt.Print(util.ToYAML(m.M)) }
