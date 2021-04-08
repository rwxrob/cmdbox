package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdbox/util"
)

func main() {
	args := util.ArgsFromStdin(os.Args[1:]...)
	for _, r := range args {
		fmt.Println(r)
	}
}
