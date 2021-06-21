package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleToStringMap() {
	m := map[string]interface{}{
		"foo": "a foo",
		"bar": "a bar",
	}
	n := util.ToStringMap(m)
	fmt.Printf("%T != %T\n", m, n)
	// Output:
	// map[string]interface {} != map[string]string
}
