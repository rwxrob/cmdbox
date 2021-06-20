package cmdbox_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func ExampleMap() {
	m := cmdbox.Map{
		"f":   "foo",
		"foo": "foo",
		"b":   "bar",
		"bar": "bar",
	}
	fmt.Println(m.Names())
	fmt.Println(m.Aliases())
	fmt.Println(m.Keys())
	fmt.Println(m.JSON())
	fmt.Println(m)
	fmt.Print(m.YAML())
	m.Print()
	// Output:
	// [bar foo]
	// [b f]
	// [b bar f foo]
	// {"b":"bar","bar":"bar","f":"foo","foo":"foo"}
	// {"b":"bar","bar":"bar","f":"foo","foo":"foo"}
	// b: bar
	// bar: bar
	// f: foo
	// foo: foo
	// b: bar
	// bar: bar
	// f: foo
	// foo: foo
}
