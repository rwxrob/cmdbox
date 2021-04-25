package cmdbox_test

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/comp"
)

func ExampleCompFunc() {

	comp.This = ""
	fmt.Println(cmdbox.CompFunc()) // [] (empty)

	// first create a few commands in the register

	x := cmdbox.New("tstamp")
	x.Method = func(args []string) error {
		fmt.Println(time.Now().Format(time.RFC3339))
		return nil
	}

	y := cmdbox.New("pomo", "start")
	y.Method = func(args []string) error {
		fmt.Println("would print a Pomodoro timer")
		return nil
	}

	comp.This = "kn"
	fmt.Println(cmdbox.CompFunc())
	comp.This = "kn po"
	fmt.Println(cmdbox.CompFunc())

	// Output:
	// []
	// [pomo tstamp]
	// [pomo]
}
