package valid_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/valid"
)

func ExampleName() {
	fmt.Println(valid.Name("yes"))
	fmt.Println(valid.Name("no"))
	fmt.Println(valid.Name("nOpe"))
	fmt.Println(valid.Name("-no"))
	fmt.Println(valid.Name("--hell=no"))
	fmt.Println(valid.Name(" no"))
	fmt.Println(valid.Name("no_no"))
	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
	// false
}
