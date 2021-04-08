package util

import "fmt"

// Dump simply dumps the stuff passed to it to standard output. Use for
// debugging. Use Print for general printing.
func Dump(stuff ...interface{}) {
	fmt.Printf("%v\n", stuff)
}
