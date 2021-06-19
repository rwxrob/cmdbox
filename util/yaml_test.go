package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleToYAML() {

	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}

	fmt.Println(util.ToYAML(sample))

	// Unordered output:
	// array:
	//     - blah
	//     - another
	// float: 1
	// int: 1
	// map:
	//     blah: another
	// string: some thing

}
