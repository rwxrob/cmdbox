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
	// {"commands":{},"messages":{"bad_type":"unsupported type: %T","invalid_name":"invalid name (must be lowercase word): %v","missing_arg":"missing argument for %v","syntax_error":"syntax error: %v","unimplemented":"unimplemented: %v"}}
	// {"commands":{"foo":{"name":"foo"}},"messages":{"bad_type":"unsupported type: %T","invalid_name":"invalid name (must be lowercase word): %v","missing_arg":"missing argument for %v","syntax_error":"syntax error: %v","unimplemented":"unimplemented: %v"}}

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
	//     invalid_name: 'invalid name (must be lowercase word): %v'
	//     missing_arg: missing argument for %v
	//     syntax_error: 'syntax error: %v'
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

func ExampleAdd_simple() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo")
	cmdbox.PrintReg()

	// Output:
	// foo:
	//     name: foo

}

func ExampleAdd_with_subcommands() {
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

func ExampleAdd_with_duplicates() {
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
	//     invalid_name: 'invalid name (must be lowercase word): %v'
	//     missing_arg: missing argument for %v
	//     new message: this is new
	//     syntax_error: 'syntax error: %v'
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
	//     invalid_name: 'invalid name (must be lowercase word): %v'
	//     missing_arg: missing argument for %v
	//     new message: this is new
	//     syntax_error: 'syntax error: %v'
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

func ExampleResolve() {
	cmdbox.Init() // just for testing

	gr := cmdbox.Add("greet", "h|help", "fr|french", "ru|russian")
	gr.Summary = "main greet composite, no method"

	fr := cmdbox.Add("greet french")
	fr.Method = func(args []string) error {
		fmt.Print("Salut")
		return nil
	}

	ru := cmdbox.Add("greet russian")
	ru.Method = func(args []string) error {
		fmt.Print("Privyet")
		return nil
	}

	h := cmdbox.Add("help")
	h.Summary = "lonely, useless, help"
	h.Method = func(args []string) error {
		fmt.Printf("help")
		return nil
	}

	tests := []struct {
		caller *cmdbox.Command
		name   string
		args   []string
	}{
		{nil, "greet", nil},
		{nil, "greet", []string{"h"}},
		{nil, "greet", []string{"hi"}},
		{nil, "greet russian", []string{"hi"}},
		{gr, "russian", []string{"hi"}},
	}

	for _, t := range tests {
		method, args := cmdbox.Resolve(t.caller, t.name, t.args)
		if method != nil {
			method(args)
			fmt.Printf(" %v %q\n", t.name, args)
			continue
		}
		fmt.Printf("failed: %v %q\n", t.name, t.args)
	}

	// Output:
	// help greet []
	// help greet []
	// help greet ["hi"]
	// Privyet greet russian ["hi"]
	// Privyet russian ["hi"]

}

func ExampleCall_nil_caller() {
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

func ExampleCall_caller_subcommand() {
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

func ExampleExecute_no_method() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Execute("foo")

	// Output:
	// unimplemented: foo
	// unexpected call to os.Exit(0) during test

}

func ExampleExecute_unimplemented_default() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")
	cmdbox.Add("foo help")
	cmdbox.Execute("foo")

	// Output:
	// unimplemented: foo
	// unexpected call to os.Exit(0) during test

}

func ExampleExecute_first_is_default() {
	cmdbox.Init() // just for testing

	cmdbox.Add("foo", "h|help")

	h := cmdbox.Add("foo help")
	h.Method = func(args []string) error {
		fmt.Println("would show foo help")
		return nil
	}

	cmdbox.Execute("foo")

	// Output:
	// would show foo help
	// unexpected call to os.Exit(0) during test

}

func ExampleExecute_completion_context() {

	os.Setenv("COMP_LINE", "foo ")
	cmdbox.Execute("foo")
	os.Setenv("COMP_LINE", "foo he")
	cmdbox.Execute("foo")
	os.Unsetenv("COMP_LINE")

	// Output:
	// h
	// help
	// unexpected call to os.Exit(0) during test
	// help
	// unexpected call to os.Exit(0) during test

}
