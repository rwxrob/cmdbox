/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleRawJSON() {
	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}
	fmt.Println(util.RawJSON(sample))

	// Output:
	// {"array":["blah","another"],"float":1,"int":1,"map":{"blah":"another"},"string":"some thing"} <nil>
}

func ExampleJSON() {
	sample := map[string]interface{}{}
	sample["int"] = 1
	sample["float"] = 1
	sample["string"] = "some thing"
	sample["map"] = map[string]interface{}{"blah": "another"}
	sample["array"] = []string{"blah", "another"}
	out, err := util.JSON(sample)
	fmt.Println(`JSON: ` + out)
	fmt.Println(err)
	// Output:
	// JSON: {
	//     "array": [
	//       "blah",
	//       "another"
	//     ],
	//     "float": 1,
	//     "int": 1,
	//     "map": {
	//       "blah": "another"
	//     },
	//     "string": "some thing"
	//   }
	// <nil>

}
