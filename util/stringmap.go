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

package util

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

// StringMap is a high-level type used to contain string data for both
// keys and values Developers can use the StringMap type in their own
// modules as a convenience since requires less declaration and is 100%
// safe for concurrency (embeds synx.Mutex). The internal map is
// exported (as M) for when developers want to do their own locking and
// mutations rather than use the public interface methods.
//
type StringMap struct {
	M map[string]string
	sync.Mutex
}

// NewStringMap returns a new StringMap with the internal
// map[string]string initialized.
//
func NewStringMap() *StringMap {
	m := new(StringMap)
	m.Init()
	return m
}

// ToStringMap converts a map[string]interface{} into a StringMap
// creating all new map values (concurrency safe clone).
//
func ToStringMap(m map[string]interface{}) *StringMap {
	n := NewStringMap()
	defer n.Unlock()
	n.Lock()
	for k, v := range m {
		n.M[k] = v.(string)
	}
	return n
}

// Init initializes (or re-initialized) the StringMap deleting all its
// values (without changing its reference).
func (m *StringMap) Init() {
	defer m.Unlock()
	m.Lock()
	if m.M == nil {
		m.M = make(map[string]string)
		return
	}
	for k := range m.M {
		delete(m.M, k)
	}
}

// Get returns a value by key name safe for concurrency. Returns empty
// string if not found.
func (m *StringMap) Get(key string) string {
	defer m.Unlock()
	m.Lock()
	if v, has := m.M[key]; has {
		return v
	}
	return ""
}

// Set set a value by key name safe in a way that is for concurrency.
func (m *StringMap) Set(key, val string) {
	defer m.Unlock()
	m.Lock()
	m.M[key] = val
}

// Delete removes one or more entries from the map in a way that is safe
// for concurrency.
func (m *StringMap) Delete(keys ...string) {
	defer m.Unlock()
	m.Lock()
	for _, k := range keys {
		delete(m.M, k)
	}
}

// Rename changes a key name in a manner safe for concurrency by adding
// the new name with the same value and deleting the old one.
func (m *StringMap) Rename(from, to string) {
	defer m.Unlock()
	m.Lock()
	v, has := m.M[from]
	if !has {
		return
	}
	m.M[to] = v
	delete(m.M, from)
}

// ------------------------------ queries -----------------------------

// Same returns a sorted list of all keys that are also values.
func (m StringMap) Same() []string {
	defer m.Unlock()
	m.Lock()
	same := []string{}
	for k, v := range m.M {
		if k == v {
			same = append(same, v)
		}
	}
	sort.Strings(same)
	return same
}

// Keys returns a sorted list of all possible string keys.
func (m StringMap) Keys() []string {
	defer m.Unlock()
	m.Lock()
	keys := make([]string, len(m.M))
	var i int
	for k, _ := range m.M {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// KeysWithout is same as Keys but omits a list of strings.
func (m StringMap) KeysWithout(omit []string) []string {
	return OmitFromSlice(m.Keys(), omit)
}

// Values returns a sorted list of all unique values safe for
// concurrency.
func (m StringMap) Values() []string {
	defer m.Unlock()
	m.Lock()
	vals := []string{}
	seen := map[string]bool{}
	for _, v := range m.M {
		if _, saw := seen[v]; saw {
			continue
		}
		vals = append(vals, v)
		seen[v] = true
	}
	sort.Strings(vals)
	return vals
}

// Aliases returns only the keys that are not identical to their value.
func (m StringMap) Aliases() []string {
	defer m.Unlock()
	m.Lock()
	a := []string{}
	for k, v := range m.M {
		if k == v {
			continue
		}
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}

// AliasesFor returns only the Aliases that point to a specific value.
// This does not include keys that are identical to their value.
func (m StringMap) AliasesFor(val string) []string {
	defer m.Unlock()
	m.Lock()
	aliases := []string{}
	for k, v := range m.M {
		if k != v && v == val {
			aliases = append(aliases, k)
		}
	}
	sort.Strings(aliases)
	return aliases
}

// KeysFor returns only the Keys that point to a specific value.
// This includes keys that are identical to their value, which will
// always be last. The rest will be sorted.
func (m StringMap) KeysFor(val string) []string {
	defer m.Unlock()
	m.Lock()
	keys := []string{}
	hasSelf := false
	for k, v := range m.M {
		if v == val {
			if k == v {
				hasSelf = true
				continue
			}
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	if hasSelf {
		keys = append(keys, val)
	}
	return keys
}

// KeysForWithout is same as KeysFor but omits a list of strings.
func (m StringMap) KeysForWithout(val string, omit []string) []string {
	return OmitFromSlice(m.KeysFor(val), omit)
}

// AliasesCombined returns a new StringMap pointer with the keys that
// point to the same value sorted, combined, and delimited into a single
// value per unique value as the key. This is useful for creating
// alternative option strings. Also see AliasesFor and KeysCombined.
func (m StringMap) AliasesCombined(delim string) *StringMap {
	n := NewStringMap()
	n.Lock()
	for _, name := range m.Values() {
		n.M[name] = strings.Join(m.AliasesFor(name), delim)
	}
	n.Unlock()
	return n
}

// KeysCombined returns a new StringMap pointer with the keys that point
// to the same value sorted, combined, and delimited into a single value
// per unique value set to the key. If any key equals the value it will
// automatically appear last in the delimited list. This is useful for
// creating alternative option strings. Also see KeysFor.
func (m StringMap) KeysCombined(delim string) *StringMap {
	n := NewStringMap()
	n.Lock()
	for _, name := range m.Values() {
		n.M[name] = strings.Join(m.KeysFor(name), delim)
	}
	n.Unlock()
	return n
}

// KeysWithoutCombined is same as KeysCombined but omits a list of
// strings.
func (m StringMap) KeysCombinedWithout(delim string, omit []string) *StringMap {
	n := NewStringMap()
	n.Lock()
	for _, name := range m.Values() {
		n.M[name] = strings.Join(m.KeysForWithout(name, omit), delim)
	}
	n.Unlock()
	return n
}

// Slice returns a slice of values fetched from the StringMap in order
// that match the key names passed. If a name is not found its value
// will be an empty string. Slice is safe for concurrency.
func (m StringMap) Slice(keys ...string) []string {
	defer m.Unlock()
	m.Lock()
	vals := make([]string, len(keys))
	for i, k := range keys {
		vals[i] = m.M[k]
	}
	return vals
}

// HasSuffix returns a new StringMap containing only those entries that
// have values with the specified suffix. HasSuffix is safe for
// concurrency. See strings.HasSuffix.
func (m StringMap) HasSuffix(s string) *StringMap {
	n := NewStringMap()
	defer n.Unlock()
	n.Lock()
	for k, v := range m.M {
		if strings.HasSuffix(v, s) {
			n.M[k] = v
		}
	}
	return n
}

// HasPrefix returns a new StringMap containing only those entries that
// have values with the specified prefis. HasPrefix is safe for
// concurrency. See strings.HasPrefix.
func (m StringMap) HasPrefix(s string) *StringMap {
	n := NewStringMap()
	defer n.Unlock()
	n.Lock()
	for k, v := range m.M {
		if strings.HasPrefix(v, s) {
			n.M[k] = v
		}
	}
	return n
}

// LongestKey returns the key and value with the longest key. The first
// longest key will win and Go maps to not promise any specific order.
func (m StringMap) LongestKey() (string, string) {
	longest := ""
	longestv := ""
	for k, v := range m.M {
		if len(k) > len(longest) {
			longest = k
			longestv = v
		}
	}
	return longest, longestv
}

// LongestValue returns the key and value with the longest value. The
// first longest value will win and Go maps to not promise any specific
// order.
func (m StringMap) LongestValue() (string, string) {
	longest := ""
	longestv := ""
	for k, v := range m.M {
		if len(v) > len(longestv) {
			longest = k
			longestv = v
		}
	}
	return longest, longestv
}

// ---------------------------- marshaling ----------------------------

// RawJSON calls MustRawJSON on the internal map.
func (m StringMap) RawJSON() string { return MustRawJSON(m.M) }

// JSON calls MustJSON on the internal map. It is often more convenient
// to simply print/Print instead since the String (from fmt.Stringer
// interface) does the same thing.
func (m StringMap) JSON() string { return MustJSON(m.M) }

// String fulfills fmt.Stringer interface as JSON.
func (m StringMap) String() string { return MustJSON(m.M) }

// Print outputs as JSON (nice when testing).
func (m StringMap) Print() { fmt.Print(MustJSON(m.M)) }

// MarshalJSON implements the json.Marshaler interface using the
// internal (M) map.
func (m StringMap) MarshalJSON() ([]byte, error) { return json.MarshalIndent(m.M, "  ", "  ") }

// UnmarshalJSON implements the json.Unmarshaler interface using the
// internal (M) map.
func (m *StringMap) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.M)
}
