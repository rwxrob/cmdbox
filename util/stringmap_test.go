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
	"gopkg.in/yaml.v3"
)

func ExampleStringMap() {
	m := util.ToStringMap(map[string]interface{}{
		"f":   "foo",
		"foo": "foo",
		"b":   "bar",
		"bar": "bar",
	})
	fmt.Println(m.Same())
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

func ExampleNewStringMap() {
	m := util.NewStringMap()
	m.Print()
	m.Set("foo", "bar")
	m.Print()

	// Output:
	// {}
	// foo: bar
}

func ExampleToStringMap() {
	m := map[string]interface{}{
		"foo": "a foo",
		"bar": "a bar",
	}
	n := util.ToStringMap(m)
	fmt.Printf("%T != %T\n", m, n)
	n.Print()
	// Output:
	// map[string]interface {} != *util.StringMap
	// bar: a bar
	// foo: a foo
}

func ExampleStringMap_Init() {
	m := util.NewStringMap()
	m.Set("foo", "bar")
	m.Print()
	m.Init()
	m.Print()

	// Output:
	// foo: bar
	// {}
}

func ExampleStringMap_Get() {
	m := util.NewStringMap()
	fmt.Println(m.Get("foo"))
	m.Set("foo", "fooval")
	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("bar"))

	// Output:
	//
	// fooval

}

func ExampleStringMap_Set() {
	m := util.NewStringMap()
	fmt.Println(m.Get("foo"))
	m.Set("foo", "bar")
	fmt.Println(m.Get("foo"))

	// Output:
	//
	// bar
}

func ExampleStringMap_Delete() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "2")
	m.Set("you", "3")
	m.Print()
	m.Delete("foo", "bar")
	m.Print()

	// Output:
	// bar: "2"
	// foo: "1"
	// you: "3"
	// you: "3"
}

func ExampleStringMap_Same() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	m.Set("you", "you")
	fmt.Println(m.Same())

	// Output:
	// [bar you]
}

func ExampleStringMap_Keys() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	m.Set("you", "you")
	fmt.Println(m.Keys())

	// Output:
	// [bar foo you]
}

func ExampleStringMap_Values() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	m.Set("you", "1")
	fmt.Println(m.Values())

	// Output:
	// [1 bar]
}

func ExampleStringMap_Aliases() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("f", "1")
	m.Set("bar", "bar")
	m.Set("b", "bar")
	m.Set("you", "1")
	fmt.Println(m.Aliases())

	// Output:
	// [b f foo you]
}

func ExampleStringMap_AliasesFor() {
	m := util.NewStringMap()
	m.Set("foo", "foo")
	m.Set("f", "foo")
	m.Set("bar", "bar")
	m.Set("b", "bar")
	m.Set("you", "foo")
	fmt.Println(m.AliasesFor("foo"))

	// Output:
	// [f you]
}

func ExampleStringMap_KeysFor() {
	m := util.NewStringMap()
	m.Set("foo", "foo")
	m.Set("f", "foo")
	m.Set("bar", "bar")
	m.Set("b", "bar")
	m.Set("you", "foo")
	fmt.Println(m.KeysFor("foo"))

	// Output:
	// [f you foo]
}

func ExampleStringMap_AliasesCombined() {
	m := util.NewStringMap()
	m.Set("foo", "foo")
	m.Set("f", "foo")
	m.Set("bar", "bar")
	m.Set("b", "bar")
	m.Set("you", "foo")
	m.AliasesCombined("|").Print()

	// Output:
	// bar: b
	// foo: f|you

}

func ExampleStringMap_KeysCombined() {
	m := util.NewStringMap()
	m.Set("foo", "foo")
	m.Set("f", "foo")
	m.Set("bar", "bar")
	m.Set("b", "bar")
	m.Set("you", "foo")
	m.KeysCombined("|").Print()

	// Output:
	// bar: b|bar
	// foo: f|you|foo

}

func ExampleStringMap_Slice() {
	m := util.NewStringMap()
	m.Set("foo", "fooval")
	m.Set("f", "fval")
	m.Set("bar", "barval")
	m.Set("b", "bval")
	m.Set("you", "youval")
	fmt.Println(m.Slice("f", "b", "none", "you"))

	// Output:
	// [fval bval  youval]
}

func ExampleStringMap_HasSuffix() {
	m := util.NewStringMap()
	m.Set("foo", "fooval")
	m.Set("f", "fval")
	m.Set("bar", "barval")
	m.Set("b", "b")
	m.Set("you", "you")
	m.HasSuffix("val").Print()

	// Output:
	// bar: barval
	// f: fval
	// foo: fooval

}

func ExampleStringMap_HasPrefix() {
	m := util.NewStringMap()
	m.Set("foo", "fooval")
	m.Set("f", "fval")
	m.Set("bar", "barval")
	m.Set("b", "b")
	m.Set("you", "you")
	m.HasPrefix("f").Print()

	// Output:
	// f: fval
	// foo: fooval

}
func ExampleStringMap_JSON() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	fmt.Println(m.JSON())

	// Output:
	// {"bar":"bar","foo":"1"}
}

func ExampleStringMap_String() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	fmt.Println(m)

	// Output:
	// {"bar":"bar","foo":"1"}
}

func ExampleStringMap_YAML() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	fmt.Println(m.YAML())

	// Output:
	// bar: bar
	// foo: "1"
}

func ExampleStringMap_Print() {
	m := util.NewStringMap()
	m.Set("foo", "1")
	m.Set("bar", "bar")
	m.Print()

	// Output:
	// bar: bar
	// foo: "1"
}

func ExampleStringMap_Rename() {
	m := util.NewStringMap()
	m.Set("foo", "val")
	m.Print()
	m.Rename("foo", "bar")
	m.Print()

	// Output:
	// foo: val
	// bar: val
}

func ExampleStringMap_MarshalYAML() {
	m := util.NewStringMap()
	m.Set("foo", "val")
	byt, _ := yaml.Marshal(m)
	fmt.Println(string(byt))
	// Output:
	// foo: val
}

func ExampleStringMap_LongestKey() {
	m := util.NewStringMap()
	m.Set("foo", "fooval")
	m.Set("ohboy", "yeah")
	fmt.Println(m.LongestKey())
	// Output:
	// ohboy yeah
}
