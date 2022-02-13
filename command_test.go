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
	// commands: {}

}

func ExampleNewCommand_subcommand() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo bar")
	x.Print()
	// Output:
	// name: foo bar
	// commands: {}

}

func ExampleNewCommand_invalid() {
	defer cmdbox.TrapPanicNoExit()
	cmdbox.NewCommand("Foo") // see valid/name.go for more

	// Output:
	// syntax error: invalid name (must be lowercase word): Foo
}

func ExampleNewCommand_dup() {
	x := cmdbox.NewCommand("foo_")
	x.Print()
	// Output:
	// name: foo_
	// commands: {}
	// default: help

}

func ExampleNewCommand_commands() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo", "h|help", "l|ls|list")
	x.Print()
	// Output:
	// name: foo
	// usage: '[h|help|l|list|ls]'
	// commands:
	//   h: help
	//   help: help
	//   l: list
	//   list: list
	//   ls: list
	// default: help
}

func ExampleCommand_JSON() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo", "help")
	fmt.Println(x.JSON())
	fmt.Println(x.String())
	fmt.Println(x)
	// Output:
	// {"name":"foo","usage":"[help]","commands":{"help":"help"},"default":"help"}
	// {"name":"foo","usage":"[help]","commands":{"help":"help"},"default":"help"}
	// {"name":"foo","usage":"[help]","commands":{"help":"help"},"default":"help"}
}

func ExampleCommand_YAML() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Print(x.YAML())
	x.Print()
	// Output:
	// name: foo
	// commands: {}
	// default: help
	// name: foo
	// commands: {}
	// default: help

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

func ExampleCommand_Legal() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	fmt.Println(x.Legal())
	x.Version = "v0.0.1"
	fmt.Println(x.Legal())
	x.Copyright = "Copyright 2021 Rob"
	fmt.Println(x.Legal())
	x.License = "Apache-2"
	fmt.Println(x.Legal())
	// Output:
	//
	// foo
	// Copyright 2021 Rob
	// foo v0.0.1
	// Copyright 2021 Rob
	// License Apache-2

}

func ExampleCommand_Add() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Add("l|ls|list", "h|help", "version")
	x.Print()
	// Output:
	// name: foo
	// commands:
	//   h: help
	//   help: help
	//   l: list
	//   list: list
	//   ls: list
	//   version: version
	// default: help
}

func ExampleCommand_Complete_commands() {
	cmdbox.Init() // just for testing
	x := cmdbox.NewCommand("foo")
	x.Add("l|ls|list")
	comp.This = "l"
	x.Complete()

	// Output:
	// list
}

func ExampleCommand_Complete_compFunc() {
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

	x := cmdbox.NewCommand("foo")
	x.Usage = "unique usage here"
	fmt.Println(x.UsageError())

	// Output:
	// usage: foo unique usage here

}

func ExampleCommand_Sigs() {
	cmdbox.Init() // just for testing
	x := cmdbox.Add("foo", "h|help", "version", "bar")
	x.Sigs().Print()

	// Output:
	// bar: bar
	// help: h|help
	// version: version
}

func ExampleCommand_Titles() {
	cmdbox.Init() // just for testing
	x := cmdbox.Add("foo", "bar", "other", "p|print", "c|comp|complete")

	b := cmdbox.Add("bar")
	b.Summary = `does bar stuff`

	h := cmdbox.Add("other")
	h.Summary = `other stuff`

	p := cmdbox.Add("print")
	p.Summary = `prints stuff`

	c := cmdbox.Add("complete")
	c.Summary = `complete stuff`

	fmt.Println(x.Titles(0, 0))
	fmt.Println(x.Titles(2, 0))
	fmt.Println(x.Titles(4, 7))

	// Output:
	// bar             - does bar stuff
	// c|comp|complete - complete stuff
	// other           - other stuff
	// p|print         - prints stuff
	//   bar             - does bar stuff
	//   c|comp|complete - complete stuff
	//   other           - other stuff
	//   p|print         - prints stuff
	//     bar     - does bar stuff
	//     c|comp|complete - complete stuff
	//     other   - other stuff
	//     p|print - prints stuff

}

func ExampleCommand_Resolve() {
	cmdbox.Init() // just for testing

	x := cmdbox.Add("foo", "bar")

	b := cmdbox.Add("bar")
	fmt.Println(b.Name)

	r := x.Resolve("bar")
	if r != nil {
		fmt.Println(r.Name)
	}

	// Output:
	// bar
	// bar
}

func ExampleCommand_Resolve_bork() {
	cmdbox.Init() // just for testing

	x := cmdbox.Add("foo me", "bork")

	r := x.Resolve("bork")
	fmt.Println(r)

	// Output:
	// <nil>
}

func ExampleCommand_Resolve_subcommand() {
	cmdbox.Init() // just for testing

	x := cmdbox.Add("foo me", "bar")

	b := cmdbox.Add("bar")
	fmt.Println(b.Name)

	r := x.Resolve("bar")
	if r != nil {
		fmt.Println(r.Name)
	}

	// Output:
	// bar
	// bar
}

/*
// TODO
func ExampleCommand_PrintHelp() {
	cmdbox.Init() // just for testing
	x := cmdbox.Add("foo", "bar", "other", "p|print", "c|comp|complete")

	b := cmdbox.Add("bar")
	b.Summary = `does bar stuff`

	h := cmdbox.Add("other")
	h.Summary = `other stuff`

	p := cmdbox.Add("print")
	p.Summary = `prints stuff`

	c := cmdbox.Add("complete")
	c.Summary = `complete stuff`

	x.PrintHelp()

	// Output:
	// NAME

}
*/
