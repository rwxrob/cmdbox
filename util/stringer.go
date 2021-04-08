/*
Package util contains utilities that are used by CmdBox itself or that
are frequently needed and useful when writing command Method
implementations (and are not already covered by other standard stuff).
*/
package util

import (
	"fmt"
)

// Stringer is anything that can be coerced into a string (see
// fmt.Sprintf) plus any function with no arguments that returns
// a string (func() string).
type Stringer interface{} // the one time i miss generics

// String adds 'func() string' to normal Go string coercion as well as
// converting nil to the empty string "".
func String(thing Stringer) string {
	switch s := thing.(type) {
	case func() string:
		return s()
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", s)
	}
}
