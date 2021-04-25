package comp_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleMonth() {
	for _, ex := range []string{"j", "J", "ju", "jul", "d", ""} {
		comp.This = ex // simulate been typed and tab pressed
		fmt.Println(comp.Month())
	}

	// Output:
	// [january june july]
	// [January June July]
	// [june july]
	// [july]
	// [december]
	// [january february march april may june july august september october november december January February March April May June July August September October November December]

}
