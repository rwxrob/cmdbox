# Each Command in Own File

This sample breaks out each command into it's own file for easier
management and modular moving around between different applications.
Each file stands alone as completely independent from the rest without
any coupling. The `main.go` composes them together in the `hello`
command.

Add bash completion with `complete -C hello hello` (or the equivalent at
`./hello` depending on how you are building/installing).

Run the program without any options to see the automatically generated
help documentation.
