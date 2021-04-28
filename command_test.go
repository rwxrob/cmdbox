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

func ExampleNew_twocommands() {
	x := cmdbox.New("pomo", "start", "stop")
	fmt.Println(x.Commands)
	// Output:
	// {
	//   "help": "help",
	//   "start": "start",
	//   "stop": "stop",
	//   "version": "version"
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
	// help
	// start
	// stop
	// version
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
	// help
	// start
	// stop
	// version
}
