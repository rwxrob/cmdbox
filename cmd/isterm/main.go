package main

import (
	"os"

	"github.com/rwxrob/cmdbox/fmt"
	"github.com/rwxrob/cmdbox/term"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	if term.IsTerminal() {
		if verbose {
			fmt.SmartPrintln("yes")
		}
		os.Exit(0)
	}
	if verbose {
		fmt.SmartPrintln("no")
	}
	os.Exit(1)
}
