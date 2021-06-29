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
	"sort"
	"strings"

	"github.com/rwxrob/cmdbox/comp"
)

// CompFunc defines a first-class function type for tab completion
// closures that sense the completion context, and return a list of
// completion strings. By managing completion logic as first-class
// functions we allow for easier completion testing and modularity. An
// empty string slice must always be returned even on failure.
type CompFunc func(x *Command) []string

// DefaultComplete is assigned CompleteCommand by default but can be
// assigned any valid CompFunc to override it. This function is called
// to perform completion for any Command that does not implement its own
// Command.CompFunc.
var DefaultComplete = CompFunc(CompleteCommand)

// CompleteCommand takes a pointer to a Command (x) returning a list of
// lexigraphically sorted combination of strings from x.Commands that
// are found in the internal register and x.Params that match the
// current completion context. Returns an empty list if anything fails.
// Note that no assertion validating that the specified command names
// exist in the register. See the Command.Complete method and comp
// subpackage.
func CompleteCommand(x *Command) []string {
	rv := []string{}
	if comp.Line() == "" {
		return rv
	}
	word := comp.Word()
	if word == " " {
		rv = append(rv, x.Commands.Keys()...)
	} else {
		rv = append(rv, x.Commands.HasPrefix(word).Keys()...)
	}
	for _, k := range x.Params {
		if word == " " || strings.HasPrefix(k, word) {
			rv = append(rv, k)
		}
	}
	sort.Strings(rv)
	return rv
}
