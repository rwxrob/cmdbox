package comp_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleYes() {
	fmt.Println(comp.Yes())
	os.Setenv("COMP_LINE", "kn jan")
	fmt.Println(comp.Yes())
	// Output:
	// false
	// true
}
