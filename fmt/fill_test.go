package fmt_test

import (
	"strings"

	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/util"
)

func ExampleFill() {
	fmt.Println(fmt.Fill("{1} and {2}\n{2} and {1}", "one", "two"))
	// Output:
	// one and two
	// two and one
}

func ExampleFillFrom() {
	r := strings.NewReader("{1} and {2}\n{2} and {1}")
	fmt.Println(fmt.FillFrom(r, "one", "two"))
	// Output:
	// one and two
	// two and one
}

func ExampleFillIn() {
	util.MockStdin("{1} and {2}\n{2} and {1}")
	defer util.UnmockStdin()
	fmt.Println(fmt.FillIn("one", "two"))
	// Output:
	// one and two
	// two and one
}
