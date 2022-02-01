package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	 x := cmdbox.Add("russian")
	x.Summary = `print "hello" in Russian`
	x.Method = func(args []string) error {
		fmt.Println("Здравствуйте")
		return nil
	}
}
