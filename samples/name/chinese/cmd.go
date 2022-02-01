package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("chinese")
	x.Summary = `print "my name is ..." in Chinese`
	x.Usage = `<name>`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		fmt.Printf("我的名字是 %v\n", args[0])
		return nil
	}
}
