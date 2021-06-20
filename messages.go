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

// Developers: Please keep this file separate from the others.

package cmdbox

// Messages contains every human readable message used by the cmdbox
// package and can be used by CmdBox module and composite developers to
// support dynamic locale output at runtime. By convention, any human
// readable natural language messages should be stored here rather than
// embedded thus preserving the ability to dynamically adapt to the
// language and locale. Since, by design, any CmdBox package can change
// anything in the Messages map module, creators should prefix their
// Message keys with the Command.Name and a single space to prevent
// unwanted accidental conflict. See JSON and Load for more.
var Messages = map[string]string{
	"invalid.name":  "invalid name (lower case words only)",
	"unimplemented": "unimplemented: %v",
}
