package valid_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/valid"
)

func ExampleName() {
	want := map[string]bool{
		"yes":           true,
		"no":            true,
		"nOpe":          false,
		"-no":           false,
		"--hell=no":     false,
		" no":           false,
		"no_no":         false,
		"foo help":      true,
		"foo help here": true,
		"foo  help":     false,
	}
	for val, expected := range want {
		if valid.Name(val) != expected {
			fmt.Printf("Expected %q to be %v\n", val, expected)
			break
		}
	}
	fmt.Println("pass")
	// Output:
	// pass
}
