package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleConvertToPrettyJSON() {

	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}

	fmt.Println(util.ConvertToPrettyJSON(sample))

	// Unordered output:
	//
	// {
	//   "array": [
	//     "blah",
	//     "another"
	//   ],
	//   "float": 1,
	//   "int": 1,
	//   "map": {
	//     "blah": "another"
	//   },
	//   "string": "some thing"
	// }
}

func ExampleConvertToJSON() {

	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}

	fmt.Println(util.ConvertToJSON(sample))

	// Output:
	// {"array":["blah","another"],"float":1,"int":1,"map":{"blah":"another"},"string":"some thing"}
}
