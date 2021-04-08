package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

type ImmaStringer struct{}

func (s ImmaStringer) String() string {
	return "Hello"
}

func ExampleStringer_String() {

	f := func() string { return "Hello" }
	fmt.Println(cmdbox.String(f))

	s := "Hello"
	fmt.Println(cmdbox.String(s))

	st := ImmaStringer{} // st.String()
	fmt.Println(cmdbox.String(st))

	// Output:
	// Hello
	// Hello
	// Hello
}
