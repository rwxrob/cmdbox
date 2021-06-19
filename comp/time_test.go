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

package comp_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleMonth() {
	defer func() { comp.This = "" }()

	for _, ex := range []string{"j", "J", "ju", "jul", "d", ""} {
		comp.This = ex // simulate been typed and tab pressed
		fmt.Println(comp.Month())
	}

	// Output:
	// [january june july]
	// [January June July]
	// [june july]
	// [july]
	// [december]
	// [january february march april may june july august september october november december January February March April May June July August September October November December]

}
