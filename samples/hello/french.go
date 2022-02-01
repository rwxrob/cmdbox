package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	 x := cmdbox.Add("french")
	x.Summary = `print "hello" in French`
	x.Method = func(args []string) error {
		fmt.Println("Bonjour")
		return nil
	}
}
