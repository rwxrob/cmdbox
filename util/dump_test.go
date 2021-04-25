package util_test

import "github.com/rwxrob/cmdbox/util"

func ExampleDump() {
	util.Dump(
		struct {
			Some string
			One  int
			True bool
		}{"some", 1, true},
	)

	// Output:
	// [{some 1 true}]

}
