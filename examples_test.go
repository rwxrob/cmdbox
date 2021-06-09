package cmdbox_test

import (
	"time"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/comp"
	"github.com/rwxrob/cmdbox/fmt"
)

func ExampleNew_simple() {
	cmdbox.Init()
	x := cmdbox.New("tstamp")
	fmt.Println(x.Name)
	// Output:
	// tstamp
}

func ExampleNew_missing_name() {
	cmdbox.Init()
	defer func() { recover(); fmt.Println("missing name") }()
	cmdbox.New("")
	// Output:
	// missing name
}

func ExampleNew_invalid_name() {
	cmdbox.Init()
	defer func() { recover(); fmt.Println("invalid name") }()
	cmdbox.New("-no")
	// Output:
	// invalid name
}

func ExampleNew_invalid_action_name() {
	cmdbox.Init()
	defer func() { recover(); fmt.Println("invalid name") }()
	cmdbox.New("foo", "-help")
	// Output:
	// invalid name
}

func ExampleNew_invalid_action_name_with_alias() {
	cmdbox.Init()
	defer func() { recover(); fmt.Println("invalid name") }()
	cmdbox.New("foo", "h|-help")
	// Output:
	// invalid name
}

func ExampleNew_invalid_alias() {
	cmdbox.Init()
	defer func() { recover(); fmt.Println("invalid alias") }()
	cmdbox.New("foo", "-h|help")
	// Output:
	// invalid alias
}

func ExampleNew_duplicates() {
	cmdbox.Init()
	x1 := cmdbox.New("foo")
	x2 := cmdbox.New("foo")
	x3 := cmdbox.New("foo")
	fmt.Println(x1.Name)
	fmt.Println(x2.Name)
	fmt.Println(x3.Name)
	fmt.Println(len(cmdbox.Duplicates()))
	fmt.Println(cmdbox.DuplicateKeys())
	// Output:
	// foo
	// foo_
	// foo__
	// 2
	// [foo_ foo__]
}

func ExampleNew_two_commands() {
	cmdbox.Init()
	x := cmdbox.New("pomo", "start", "stop")
	fmt.Println(x.Commands)
	// Output:
	// {
	//   "start": "start",
	//   "stop": "stop"
	// }
}

func ExampleCommand_Complete_ignored() {
	cmdbox.Init()
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
	cmdbox.Init()

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
	cmdbox.Init()

	x := cmdbox.New("pomo", "start", "stop")
	x.Params = []string{"25m", "aparam", "-10.4"}
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

	comp.This = "2"
	x.Complete()

	comp.This = "  "
	x.Complete()

	comp.This = "-"
	x.Complete()

	// Output:
	// start
	// stop
	// stop
	// 25m
	// -10.4
	// 25m
	// aparam
	// start
	// stop
	// -10.4
}

func ExampleCommand_NameFromSig() {
	sig := "foo d|uncode|decode"
	fmt.Println(cmdbox.NameFromSig(sig))
	sig = "d|uncode|decode"
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
	// decode
}

func ExampleString() {
	cmdbox.Init()
	x := cmdbox.New("foo")
	x.Author = "Rob"
	fmt.Println(cmdbox.String())
	// Output:
	// {
	//   "PackageVersion": "",
	//   "Register": {
	//     "foo": {
	//       "Author": "Rob",
	//       "Name": "foo",
	//       "Summary": ""
	//     }
	//   }
	// }
}

func ExamplePrint() {
	cmdbox.Init()
	x := cmdbox.New("foo")
	x.Author = "Rob"
	cmdbox.Print()
	// Output:
	// {
	//   "PackageVersion": "",
	//   "Register": {
	//     "foo": {
	//       "Author": "Rob",
	//       "Name": "foo",
	//       "Summary": ""
	//     }
	//   }
	// }
}
