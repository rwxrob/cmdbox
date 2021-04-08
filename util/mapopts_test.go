package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleMapOpts() {

	// For a given command:
	//
	//   foo README.md meta.yml \
	//     --template=template \
	//     "--withquotes='with quote'" \
	//     --name "Mr. Rob" \
	//     - anotherarg -- wasarg \
	//     -abc -x -y -z \
	//     -t notavalue
	//

	cmdArgs := []string{
		"foo", "README.md", "meta.yml",
		"--template=template.html",
		"--withquotes='with quotes'",
		"--name", "Mr. Rob",
		"anotherarg", "--", "wasarg", "-",
		"-abc", "-x", "-y", "-z",
		"-t", "notavalue",
	}

	opts, args := util.MapOpts(cmdArgs[1:])

	fmt.Println("Options:")
	fmt.Println(util.ConvertToJSON(opts))

	fmt.Println("Arguments:")
	fmt.Println(util.ConvertToJSON(args))

	// Unordered output:
	//
	// Options:
	// {
	//   "a": "",
	//   "b": "",
	//   "c": "",
	//   "name": "Mr. Rob",
	//   "t": "",
	//   "template": "template.html",
	//   "withquotes": "'with quotes'",
	//   "x": "",
	//   "y": "",
	//   "z": ""
	// }
	// Arguments:
	// [
	//   "README.md",
	//   "meta.yml",
	//   "anotherarg",
	//   "--",
	//   "wasarg",
	//   "-",
	//   "notavalue"
	// ]
}
