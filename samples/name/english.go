package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("english")
	x.Summary = `print "My name is " in English`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		fmt.Printf("My name is %v\n", args[0])
		return nil
	}
}
