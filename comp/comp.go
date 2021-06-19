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

package comp

import (
	"os"
	"strings"
)

// This can be set to force a completion context no matter what the
// shell or situation. This is useful mostly for testing and should
// usually never be modified by any Command or subcommand, which can
// create very unpredictable results for the completions of other
// Commands. (And, as a reminder, always vet any CmdBox command modules
// you import to be sure they do not do insecure things with the cmdbox
// state, such as cmdbox.This.) See Line(), Args(), Word() as well.
var This string

// Func defines a type for first-class tab completion closure functions
// that can accept a varying number of string arguments, sense the
// completion context (see Line, Args, Word, and This), and return
// a list of completion strings. By managing completion logic as first
// class functions we allow for easier completion testing and modularity.
// An empty string slice must always be returned even on failure.
type Func func(i ...string) []string

// Yes returns true if the current executable is being called in
// a completion context (usually someone tapping tab). This is detected
// only by the presence of the Bash COMP_LINE environment variable. See
// Line() and Programmable Completion in the bash man page.
func Yes() bool { return len(Line()) > 0 }

// Line returns the full current command line being evaluated for this
// executable being run in completion context (see Yes).
func Line() string {
	if This != "" {
		return This
	}
	return os.Getenv("COMP_LINE")
}

// Args returns Line but as a slice of strings. If the Line() has one or
// more spaces at the end include an space (" ") as the last item. This
// is to distinquish between users wanting to tab on prefixes versus all
// the possibilities for a command.
//
// WARNING: The first element of any arguments list is always determined
// by the underlying operating system and can be inconsistent and even
// modified from the actual executable. Use caution when relying on it
// for consistent and secure program behavior.
func Args() []string {
	args := []string{}
	line := Line()
	if line == "" {
		return args
	}
	args = strings.Split(line, " ")
	if line[len(line)-1] == ' ' {
		args = append(args, " ")
	}
	return args
}

// Word returns the last of Args or empty string. This includes the
// special single space (" ") string indicating there were trailing
// spaces when completion was invoked. This should not be confused with
// the empty string indicating Args was empty.
func Word() string {
	args := Args()
	if len(args) > 0 {
		return args[len(args)-1]
	}
	return ""
}
