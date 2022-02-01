package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	 x := cmdbox.Add("english")
	x.Summary = `print "hello" in English`
	x.Method = func(args []string) error {
		fmt.Println("Hello")
		return nil
	}
}
