package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleSplitBarPop() {
	sig := "foo d|uncode|decode"
	fmt.Println(util.SplitBarPop(sig))
	sig = "d|uncode|decode"
	fmt.Println(util.SplitBarPop(sig))
	sig = "d|decode"
	fmt.Println(util.SplitBarPop(sig))
	sig = "decode"
	fmt.Println(util.SplitBarPop(sig))
	sig = ""
	fmt.Println(util.SplitBarPop(sig))
	// Output:
	// decode
	// decode
	// decode
	// decode
}
