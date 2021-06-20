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

/*
Package comp is the Bash tab completion subpackage of CmdBox. Any CmdBox composite automatically supports tab completion. To enable it simply execute the following complete command from the shell or .bashrc file:

    complete -C foo foo

Only Bash completion is supported. Additional shell completion support
will not be added. This is by design and explained in the Design
Considerations of the project documentation. Note that CmdBox composite
commands are efficient without completion using aliases and
conversational modern command line user interfaces with intelligent
natural language completion.
*/
package comp
