# Some Commands in Own Packages[j/Modules

This sample breaks out the Russian and Chinese language command into
their own packages for completely independent development that could be
in separate git module repos entirely. This is accomplished through the
use of standard Go "side-effect" imports.

Note the use `cmd.go` file name and `cmd` package name conventions and
the inclusion of a convenience command for each independent command
module. This allows the command to be used and tested independently from
any composite into which it is imported.

Add bash completion with `complete -C name name` (or the equivalent at
`./name` depending on how you are building/installing).

Run the program without any options to see the automatically generated
help documentation.
