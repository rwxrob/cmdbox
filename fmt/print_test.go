package fmt_test

import (
	"github.com/rwxrob/cmdbox/fmt"
)

func ExamplePrintln() {

	stringer := func() string { return "something 1" }

	fmt.Println("something 1")
	fmt.Println("some%v %v", "thing", 1)
	fmt.Println(stringer)
	fmt.Println()
	fmt.Println(nil)

	// Output:
	// something 1
	// something 1
	// something 1
	//
	//
}

func ExampleSprint() {

	stringer := func() string { return "something 1" }

	fmt.Println(fmt.Sprint("something 1"))
	fmt.Println(fmt.Sprint("some%v %v", "thing", 1))
	fmt.Println(fmt.Sprint(stringer))
	fmt.Println(fmt.Sprint())

	// Output:
	// something 1
	// something 1
	// something 1
	//
}
