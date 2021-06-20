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

package fmt_test

import (
	"strings"

	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/util"
)

func ExampleFill() {
	fmt.Println(fmt.Fill("{1} and {2}\n{2} and {1}", "one", "two"))
	// Output:
	// one and two
	// two and one
}

func ExampleFillFrom() {
	r := strings.NewReader("{1} and {2}\n{2} and {1}")
	fmt.Println(fmt.FillFrom(r, "one", "two"))
	// Output:
	// one and two
	// two and one
}

func ExampleFillIn() {
	util.MockStdin("{1} and {2}\n{2} and {1}")
	defer util.UnmockStdin()
	fmt.Println(fmt.FillIn("one", "two"))
	// Output:
	// one and two
	// two and one
}
