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
	"os"

	"github.com/rwxrob/cmdbox/comp"
)

func ExampleYes() {
	cl := os.Getenv("COMP_LINE")
	defer func() { os.Setenv("COMP_LINE", cl); comp.This = "" }()

	fmt.Println(comp.Yes())
	comp.This = "kn jan"
	fmt.Println(comp.Yes())
	comp.This = ""
	fmt.Println(comp.Yes())
	os.Setenv("COMP_LINE", "kn jan")
	fmt.Println(comp.Yes())

	// Output:
	// false
	// true
	// false
	// true
}

func ExampleLine() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Line())
	// Output:
	// go test -run Line
}

func ExampleArgs() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Args())
	// Output:
	// [go test -run Line]
}

func ExampleWord() {
	defer func() { comp.This = "" }()
	comp.This = "go test -run Line"
	fmt.Println(comp.Word())
	comp.This = ""
	fmt.Printf("%q\n", comp.Word())
	fmt.Println(comp.Word() == "")
	comp.This = "kn "
	fmt.Printf("%q\n", comp.Word())
	// Output:
	// Line
	// ""
	// true
	// " "
}
