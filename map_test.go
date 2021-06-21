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

func ExampleToMap() {
	m := map[string]interface{}{
		"foo": "a foo",
		"bar": "a bar",
	}
	n := cmdbox.ToMap(m)
	fmt.Printf("%T != %T\n", m, n)
	// Output:
	// map[string]interface {} != cmdbox.Map
}
