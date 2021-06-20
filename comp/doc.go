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
