package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleFilter_comments() {
	in := "some\nthing here"
	fn := func(in string) string { return "# " + in }
	fmt.Println(util.Filter(in, fn))
	// Output:
	// # some
	// # thing here
}
