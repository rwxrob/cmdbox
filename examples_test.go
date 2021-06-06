package cmdbox_test

import (
	"time"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/fmt"
)

func ExampleNew_simple() {
	x := cmdbox.New("tstamp")
	fmt.Println(x.Name)
	// Output:
	// tstamp
}

func ExampleNew_duplicates() {
	x1 := cmdbox.New("foo")
	x2 := cmdbox.New("foo")
	fmt.Println(x1.Name)
	fmt.Println(x2.Name)
	fmt.Println(len(cmdbox.Duplicates()))
	fmt.Println(cmdbox.DuplicateKeys())
	// Output:
	// foo
	// foo
	// 1
	// [foo_]
}

func ExampleNew_two_commands() {
	x := cmdbox.New("pomo", "start", "stop")
	fmt.Println(x.Commands)
	// Output:
	// {
	//   "start": "start",
	//   "stop": "stop"
	// }
}

func ExampleCommand_Complete_ignored() {
	x := cmdbox.New("tstamp")
	x.Method = func(args []string) error {
		fmt.Println(time.Now().Format(time.RFC3339))
		return nil
	}
	comp.This = "ignored"
	x.Complete()
	// Output:
}

func ExampleCommand_Complete_actions() {

	x := cmdbox.New("pomo", "start", "stop")
	x.Method = func(a []string) error {
		if len(a) == 0 {
			return x.UsageError()
		}
		switch a[0] {
		case "start":
			fmt.Println("would start")
		case "stop":
			fmt.Println("would stop")
		}
		return nil
	}

	comp.This = "st"
	x.Complete()

	comp.This = "sto"
	x.Complete()

	comp.This = "  "
	x.Complete()

	// Output:
	// start
	// stop
	// stop
	// start
	// stop
}

func ExampleCommand_Complete_params() {

	x := cmdbox.New("pomo", "start", "stop")
	x.Params = []string{"-v", "aparam"}
	x.Method = func(a []string) error {
		if len(a) == 0 {
			return x.UsageError()
		}
		switch a[0] {
		case "start":
			fmt.Println("would start")
		case "stop":
			fmt.Println("would stop")
		}
		return nil
	}

	comp.This = "st"
	x.Complete()

	comp.This = "sto"
	x.Complete()

	comp.This = "-"
	x.Complete()

	comp.This = "  "
	x.Complete()

	// Output:
	// start
	// stop
	// stop
	// -v
	// -v
	// aparam
	// start
	// stop
}

func ExampleCommand_NameFromSig() {
	sig := "d|uncode|decode"
	fmt.Println(cmdbox.NameFromSig(sig))
	sig = "d|decode"
	fmt.Println(cmdbox.NameFromSig(sig))
	sig = "decode"
	fmt.Println(cmdbox.NameFromSig(sig))
	sig = ""
	fmt.Println(cmdbox.NameFromSig(sig))
	// Output:
	// decode
	// decode
	// decode
}

func ExampleString() {
	x := cmdbox.New("foo")
	x.Author = "Rob"
	fmt.Println(cmdbox.String())
	// Output:
	// {
	//   "foo": {
	//     "Author": "Rob",
	//     "Name": "foo",
	//     "Summary": ""
	//   }
	// }
}

func ExamplePrint() {
	x := cmdbox.New("foo")
	x.Author = "Rob"
	cmdbox.Print()
	// Output:
	// {
	//   "foo": {
	//     "Author": "Rob",
	//     "Name": "foo",
	//     "Summary": ""
	//   }
	// }
}
