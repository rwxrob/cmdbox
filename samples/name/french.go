package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("french")
	x.Summary = `print "My name is" in French`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		fmt.Printf("Je m'appelle %v\n", args[0])
		return nil
	}
}
