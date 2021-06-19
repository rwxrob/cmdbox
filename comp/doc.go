/*
Package comp is the tab completion subpackage of CmdBox. It contains the
common completion functions needed by most CmdBox command modules.
As additional completion contexts are identified and implemented
they will be added here.

Bash Only

    complete -C foo foo

Only Bash completion (detected from COMP_LINE) is supported and planned.
CmdBox composite commands can be used efficiently without completion on
systems without Bash using aliases and conversational user interfaces with natural language completion.
*/
package comp
