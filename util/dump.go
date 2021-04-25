package util

import "fmt"

// Dump prints the string form of the stuff passed to it to standard
// output. Use for debugging. Use the cmdbox/fmt package for general
// printing.
func Dump(a ...interface{}) { fmt.Printf("%v\n", a) }
