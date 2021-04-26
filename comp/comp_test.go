package comp_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleYes() {
	cl := os.Getenv("COMP_LINE")
	defer func() { os.Setenv("COMP_LINE", cl); comp.This = "" }()

	fmt.Println(comp.Yes())
	comp.This = "kn jan"
	fmt.Println(comp.Yes())
	comp.This = ""
	fmt.Println(comp.Yes())
	os.Setenv("COMP_LINE", "kn jan")
	fmt.Println(comp.Yes())

	// Output:
	// false
	// true
	// false
	// true
}

func ExampleLine() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Line())
	// Output:
	// go test -run Line
}

func ExampleArgs() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Args())
	// Output:
	// [go test -run Line]
}

func ExampleWord() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Word())
	comp.This = ""
	fmt.Printf("%q\n", comp.Word())
	fmt.Println(comp.Word() == "")
	comp.This = "kn "
	fmt.Printf("%q\n", comp.Word())
	// Output:
	// Line
	// ""
	// true
	// " "
}
