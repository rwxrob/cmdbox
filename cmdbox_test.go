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

func ExampleReg() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 

	// by default the internal registry is empty but initialized

	fmt.Println(cmdbox.Reg.Names())

	// Output:
	// []
}

func ExampleJSON() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	fmt.Println(cmdbox.JSON())
	cmdbox.Add("foo")
	fmt.Println(cmdbox.JSON())
}

func ExampleInit() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	fmt.Println(cmdbox.Names())
	cmdbox.Add("foo")
	fmt.Println(cmdbox.Names())
	cmdbox.TestOn()
	defer cmdbox.TestOff()
	fmt.Println(cmdbox.Names())

	// Output:
	// []
	// [foo]
	// []

}

func ExampleAdd_simple() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	fmt.Println(cmdbox.Names())

	// Output:
	// [foo]
}

func ExampleAdd_with_Subcommands() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("foo bar")
	fmt.Println(cmdbox.Names())

	// Output:
	// [foo foo bar]

}

func ExampleAdd_with_Duplicates() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("foo")
	fmt.Println(cmdbox.Dups())
	fmt.Println(cmdbox.Names())

	// Output:
	// [foo_]
	// [foo foo_]

}

func ExampleNames() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("bar")
	fmt.Println(cmdbox.Names())

	// Output:
	// [bar foo]

}

func ExampleRename() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("foo")
	cmdbox.Rename("foo_", "bar")
	fmt.Println(cmdbox.Names())

	// Output:
	// [bar foo]

}

func ExampleGet() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo", "bar")
	cmdbox.Get("foo").Print()

	// Output:
	// {
	//     "name": "foo",
	//     "usage": "bar",
	//     "commands": {
	//       "bar": "bar"
	//     }
	//   }

}

func ExampleSlice() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("foo help")
	cmdbox.Add("bar")
	for _, x := range cmdbox.Slice("foo", "bar") {
		fmt.Println(x.Name)
	}

	// Output:
	// foo
	// bar

}

func ExampleSet() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	foo := cmdbox.NewCommand("foo")
	cmdbox.Set("foo", foo)
	bar := cmdbox.NewCommand("bar")
	cmdbox.Set("bar", bar)
	cmdbox.Set("bar", foo)
	newbar := cmdbox.Get("bar")
	fmt.Println(newbar.Name)

	// Output:
	// foo
}

func ExampleDelete() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.Add("foo")
	cmdbox.Add("bar")
	fmt.Println(cmdbox.Names())
	cmdbox.Delete("bar")
	fmt.Println(cmdbox.Names())

	// Output:
	// [bar foo]
	// [foo]

}

func ExampleResolve() {
	cmdbox.TestOn()
	defer cmdbox.TestOff()

	gr := cmdbox.Add("greet", "fr|french", "ru|russian")
	gr.Summary = "main greet composite, no method"

	fr := cmdbox.Add("greet french")
	fr.Method = func(args ...string) error {
		fmt.Print("Salut")
		return nil
	}

	ru := cmdbox.Add("greet russian")
	ru.Method = func(args ...string) error {
		fmt.Print("Privyet")
		return nil
	}

	tests := []struct {
		caller *cmdbox.Command
		name   string
		args   []string
	}{
		{nil, "greet", nil},            // usage
		{nil, "greet", []string{"h"}},  // usage
		{nil, "greet", []string{"hi"}}, // usage
		{nil, "greet russian", []string{"hi"}},
		{gr, "russian", []string{"hi"}},
	}

	for _, t := range tests {
		method, args := cmdbox.Resolve(t.caller, t.name, t.args)
		if method != nil {
			method(args...)
			fmt.Printf("%v %q\n", t.name, args)
			continue
		}
		fmt.Printf("failed: %v %q\n", t.name, t.args)
	}

	// Output:
	// failed: greet []
	// failed: greet ["h"]
	// failed: greet ["hi"]
	// Privyetgreet russian ["hi"]
	// Privyetrussian ["hi"]

}

func ExampleCall_nil_Caller() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 

	x := cmdbox.Add("greet")

	x.Method = func(args ...string) error {
		fmt.Println("hello")
		return nil
	}

	cmdbox.Call(nil, "greet")

	// Output:
	// hello

}

func ExampleCall_caller_Subcommand() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 

	caller := cmdbox.Add("foo", "h|help")

	x := cmdbox.Add("foo help")

	x.Method = func(args ...string) error {
		fmt.Printf("help for foo %v\n", args)
		return nil
	}

	cmdbox.Call(caller, "help")
	cmdbox.Call(nil, "foo help")
	cmdbox.Call(caller, "help", "with", "args")

	// Output:
	// help for foo []
	// help for foo []
	// help for foo [with args]

}

func ExampleExecute_no_Method() {
	cmdbox.TestOn()
	defer cmdbox.TestOff() 
	cmdbox.ExitOff()
	defer cmdbox.ExitOn()

	cmdbox.Add("foo", "some")
	cmdbox.Execute("foo")

	// Output:
	// usage: foo some

}
