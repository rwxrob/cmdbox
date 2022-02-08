package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleOmitFromSlice() {
	from := []string{"foo", "some", "hidden", "bar", "admin"}
	omit := []string{"hidden", "admin"}
	fmt.Println(util.OmitFromSlice(from, omit))
	// Output:
	// [foo some bar]
}

func ExampleInSlice() {
	slice := []string{"foo", "some", "hidden", "bar", "admin"}
	fmt.Println(util.InSlice("hidden", slice))
	// Output:
	// true
}
