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

func ExampleFilter_slice_String() {
	in := []string{"some", "thing here"}
	fn := func(in string) string { return "# " + in }
	fmt.Println(util.Filter(in, fn))
	// Output:
	// # some
	// # thing here
}

func ExampleFilter_slice_Numbers() {
	fn := func(in string) string { return "# " + in }
	fmt.Print(util.Filter([]int{1, 2}, fn))
	fmt.Print(util.Filter([]float32{3.0, 4.3}, fn))
	// Output:
	// # 1
	// # 2
	// # 3
	// # 4.3
}
