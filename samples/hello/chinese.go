package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	 x := cmdbox.Add("chinese")
	x.Summary = `print "hello" in Chinese`
	x.Method = func(args []string) error {
		fmt.Println("你好")
		return nil
	}
}
