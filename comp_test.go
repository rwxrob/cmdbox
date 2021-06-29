package cmdbox_test

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/comp"
)

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
