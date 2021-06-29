package term_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/term"
)

func ExampleIsTerminal() {
	term.ForceYes = true
	fmt.Println(term.IsTerminal())
	term.ForceNo = true
	fmt.Println(term.IsTerminal())
	term.ForceYes = false
	fmt.Println(term.IsTerminal())

	// Output:
	// true
	// false
	// false
}
