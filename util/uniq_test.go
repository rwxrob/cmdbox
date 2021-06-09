package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleUniqMapVals() {
	mymap := map[string]string{
		"first":  "one",
		"second": "two",
		"third":  "two",
	}
	fmt.Println(util.UniqStrMapVals(mymap))
	// Output:
	// [one two]
}
