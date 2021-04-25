package comp_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleLine() {
	os.Setenv("COMP_LINE", "go test -run Line")
	fmt.Println(comp.Line())
	// Output:
	// go test -run Line
}

func ExampleLineArgs() {
	os.Setenv("COMP_LINE", "go test -run Line")
	fmt.Println(comp.LineArgs())
	// Output:
	// [go test -run Line]
}
