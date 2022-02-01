package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("russian")
	x.Summary = `print "my name is ..." in Russian`
	x.Usage = `<name>`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		fmt.Printf("Меня зoвут %v\n", args[0])
		return nil
	}
}
