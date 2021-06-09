package valid_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/valid"
)

func ExampleName() {
	fmt.Println(valid.Name("-no"))
	fmt.Println(valid.Name("yes"))
	// Output:
	// false
	// true
}
