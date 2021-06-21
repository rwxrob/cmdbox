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
	"os"
	"strings"

	"github.com/rwxrob/cmdbox"
)

func ExampleReg() {
	cmdbox.Init() // just for testing

	m := cmdbox.Reg()
	fmt.Println(m)
	m["foo"] = cmdbox.NewCommand("foo")
	cmdbox.PrintReg()

	// Output:
	// map[]
	// foo:
	//     name: foo

}

func ExampleJSON() {
	cmdbox.Init() // just for testing

	fmt.Println(cmdbox.JSON())
	cmdbox.Add("foo")
	fmt.Println(cmdbox.JSON())

	// Output:
	// {"commands":{},"messages":{"bad_type":"unsupported type: %T","invalid_name":"invalid name (lower case words only)","unimplemented":"unimplemented: %v"}}
	// {"commands":{"foo":{"name":"foo"}},"messages":{"bad_type":"unsupported type: %T","invalid_name":"invalid name (lower case words only)","unimplemented":"unimplemented: %v"}}

}

func ExampleYAML() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	fmt.Print(cmdbox.YAML())

	// Output:
	// commands:
	//     foo:
	//         name: foo
	// messages:
	//     bad_type: 'unsupported type: %T'
	//     invalid_name: invalid name (lower case words only)
	//     unimplemented: 'unimplemented: %v'

}

func ExampleInit() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	cmdbox.PrintReg()

	fmt.Println("-----")

	cmdbox.Init()
	cmdbox.PrintReg()

	// Output:
	// foo:
	//     name: foo
	// -----
	// {}

}

func ExampleAdd_Simple() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	cmdbox.PrintReg()

	// Output:
	// foo:
	//     name: foo

}

func ExampleAdd_With_Subcommands() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo help")
	cmdbox.PrintReg()

	// Output:
	// foo:
	//     name: foo
	//     commands:
	//         h: help
	//         help: help
	//     default: help
	// foo help:
	//     name: foo help

}

func ExampleAdd_With_Duplicates() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo", "h|help")
	fmt.Println(cmdbox.Dups())
	cmdbox.PrintReg()

	// Output:
	// [foo_]
	// foo:
	//     name: foo
	//     commands:
	//         h: help
	//         help: help
	//     default: help
	// foo_:
	//     name: foo_
	//     commands:
	//         h: help
	//         help: help
	//     default: help

}

func ExampleNames() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	cmdbox.Add("bar")

	fmt.Println(cmdbox.Names())

	// Output:
	// [bar foo]

}

func ExampleRename() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo", "h|help")

	cmdbox.Rename("foo_", "bar")
	cmdbox.PrintReg()

	// Output:
	// bar:
	//     name: bar
	//     commands:
	//         h: help
	//         help: help
	//     default: help
	// foo:
	//     name: foo
	//     commands:
	//         h: help
	//         help: help
	//     default: help

}

func ExampleLoad() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo help")

	buf := strings.NewReader(`
commands:
    foo:
        usage: '[h|help]'
        summary: some summary
        description: some description
    foo help:
        summary: display foo help
messages:
    new message: this is new
    unimplemented: nope, don't have this yet
`)

	err := cmdbox.Load(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	cmdbox.Print()

	// Output:
	// commands:
	//     foo:
	//         name: foo
	//         summary: some summary
	//         usage: '[h|help]'
	//         description: some description
	//         commands:
	//             h: help
	//             help: help
	//         default: help
	//     foo help:
	//         name: foo help
	//         summary: display foo help
	// messages:
	//     bad_type: 'unsupported type: %T'
	//     invalid_name: invalid name (lower case words only)
	//     new message: this is new
	//     unimplemented: nope, don't have this yet

}

func ExampleLoadFS() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo help")

	cmdbox.LoadFS(os.DirFS("testdata"), "loadfs.yaml")
	cmdbox.Print()

	// Output:
	// commands:
	//     foo:
	//         name: foo
	//         summary: some summary
	//         usage: '[h|help]'
	//         description: some description
	//         commands:
	//             h: help
	//             help: help
	//         default: help
	//     foo help:
	//         name: foo help
	//         summary: display foo help
	// messages:
	//     bad_type: 'unsupported type: %T'
	//     invalid_name: invalid name (lower case words only)
	//     new message: this is new
	//     unimplemented: nope, don't have this yet

}

func ExampleGet() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Get("foo").Print()

	// Output:
	// name: foo
	// commands:
	//     h: help
	//     help: help
	// default: help
}

func ExampleSlice() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo help")
	cmdbox.Add("bar")

	cmds := cmdbox.Slice("foo", "bar")
	fmt.Println(cmds)

	// Output:
	// [{"name":"foo","commands":{"h":"help","help":"help"},"default":"help"} {"name":"bar"}]

}

func ExampleSet() {
	cmdbox.Init() // just for testing

	foo := cmdbox.NewCommand("foo")
	cmdbox.Set("foo", foo)
	cmdbox.PrintReg()

	fmt.Println("-----")

	bar := cmdbox.NewCommand("bar")
	cmdbox.Set("bar", bar)
	cmdbox.PrintReg()

	fmt.Println("-----")

	cmdbox.Set("bar", foo)
	cmdbox.PrintReg()

	// Output:
	// foo:
	//     name: foo
	// -----
	// bar:
	//     name: bar
	// foo:
	//     name: foo
	// -----
	// bar:
	//     name: foo
	// foo:
	//     name: foo
}

func ExampleDelete() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	cmdbox.Add("bar")
	cmdbox.PrintReg()
	fmt.Println("-----")

	cmdbox.Delete("bar")
	cmdbox.PrintReg()

	// Output:
	// bar:
	//     name: bar
	// foo:
	//     name: foo
	// -----
	// foo:
	//     name: foo

}

func ExampleCall_nil_Caller() {
	cmdbox.Init() // just for testing

	x := cmdbox.Add("greet")

	x.Method = func(args []string) error {
		fmt.Println("hello")
		return nil
	}

	cmdbox.Call(nil, "greet", nil)

	// Output:
	// hello

}

func ExampleCall_Caller_Subcommand() {
	cmdbox.Init() // just for testing

	caller := cmdbox.Add("foo", "h|help")

	x := cmdbox.Add("foo help")

	x.Method = func(args []string) error {
		fmt.Printf("help for foo %v\n", args)
		return nil
	}

	cmdbox.Call(caller, "help", nil)
	cmdbox.Call(nil, "foo help", nil)
	cmdbox.Call(caller, "help", []string{"with", "args"})

	// Output:
	// help for foo []
	// help for foo []
	// help for foo [with args]

}

func ExampleExecute() {
	cmdbox.Init() // just for testing

	// add foo with default (first) help subcommand
	x := cmdbox.Add("foo", "h|help")
	x.Summary = "foo the things"

	// will fail since no method and help is unimplemented
	// (could also leave "foo" out if exe named foo)
	cmdbox.Execute("foo")

	// add foo's help subcommand (without a method)
	h := cmdbox.Add("foo help")
	h.Summary = "help for foo"

	// will still fail since help method unimplemented
	cmdbox.Execute("foo")

	h.Method = func(args []string) error {
		fmt.Println("would print foo help and exit")
		return nil
	}

	// will call help method since foo has no method, yet
	cmdbox.Execute("foo")

	x.Method = func(args []string) error {
		fmt.Println("would foo and then exit")
		return nil
	}

	// will call foo method
	cmdbox.Execute("foo")

	// completion context, will list all possible completions
	os.Setenv("COMP_LINE", "foo ")
	cmdbox.Execute("foo")
	os.Unsetenv("COMP_LINE")

	// Output:
	// unimplemented: help
	// unexpected call to os.Exit(0) during test
	// unimplemented: help
	// unexpected call to os.Exit(0) during test
	// unimplemented: help
	// unexpected call to os.Exit(0) during test
	// would foo and then exit
	// unexpected call to os.Exit(0) during test
	// h
	// help
	// unexpected call to os.Exit(0) during test

}
