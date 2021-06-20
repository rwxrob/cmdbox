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
	"github.com/rwxrob/cmdbox/comp"
)

func ExampleNewCommand_simple() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Print()
	// Output:
	// name: foo
}

func ExampleNewCommand_subcommand() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo bar")
	x.Print()
	// Output:
	// name: foo bar
}

func ExampleNewCommand_invalid() {
	defer cmdbox.TrapPanic()
	cmdbox.NewCommand("Foo") // see valid/name.go for more
	// Output:
	// invalid name (lower case words only)
}

func ExampleNewCommand_dup() {
	x := cmdbox.NewCommand("foo_")
	x.Print()
	// Output:
	// name: foo_
}

func ExampleNewCommand_commands() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo", "h|help", "version", "l|ls|list")
	x.Print()
	// Output:
	// name: foo
	// commands:
	//     h: help
	//     help: help
	//     l: list
	//     list: list
	//     ls: list
	//     version: version
	// default: help
}

func ExampleCommand_JSON() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Println(x.JSON())
	fmt.Println(x.String())
	fmt.Println(x)
	// Output:
	// {"name":"foo"}
	// {"name":"foo"}
	// {"name":"foo"}
}

func ExampleCommand_YAML() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Print(x.YAML())
	x.Print()
	// Output:
	// name: foo
	// name: foo
}

func ExampleCommand_Title() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Println(x.Title())
	x.Summary = "just a foo"
	fmt.Println(x.Title())
	// Output:
	// foo
	// foo - just a foo
}

func ExampleCommand_VersionLine() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Println(x.VersionLine())
	x.Version = "v0.0.1"
	fmt.Println(x.VersionLine())
	x.Copyright = "Copyright 2021 Rob"
	fmt.Println(x.VersionLine())
	x.License = "Apache-2"
	fmt.Println(x.VersionLine())
	// Output:
	//
	// foo v0.0.1
	// foo v0.0.1 Copyright 2021 Rob
	// foo v0.0.1 Copyright 2021 Rob (Apache-2)
}

func ExampleCommand_Add() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Add("l|ls|list", "h|help", "version")
	x.Print()
	// Output:
	// name: foo
	// commands:
	//     h: help
	//     help: help
	//     l: list
	//     list: list
	//     ls: list
	//     version: version
}

func ExampleComplete_commands() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Add("l|ls|list", "h|help", "version")
	comp.This = "l"
	x.Complete()
	comp.This = "he"
	x.Complete()
	comp.This = "z"
	x.Complete()
	// Output:
	// l
	// list
	// ls
	// help
}

func ExampleComplete_params() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Params = []string{"-1", "25m", "0.2", "FULL"}
	comp.This = "-"
	x.Complete()
	comp.This = "2"
	x.Complete()
	comp.This = "0"
	x.Complete()
	comp.This = "F"
	x.Complete()
	// Output:
	// -1
	// 25m
	// 0.2
	// FULL
}

func ExampleComplete_CompFunc() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.CompFunc = func(x *cmdbox.Command) []string {
		// deliberatly always return the same thing
		// could add filter logic here
		return []string{"bar", "this"}
	}
	x.Complete()
	comp.This = "b"
	x.Complete()
	// Output:
	// bar
	// this
	// bar
	// this
}

func ExampleCommand_Unimplemented() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Println(x.Unimplemented("help"))
	// Output:
	// unimplemented: help
}

func ExampleCommand_UsageError() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo", "h|help")
	fmt.Println(x.UsageError())
	x.Usage = "[h|help]"
	fmt.Println(x.UsageError())
	// Output:
	//
	// [h|help]
}
