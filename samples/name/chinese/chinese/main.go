package main

import (
	"github.com/rwxrob/cmdbox"
	_ "github.com/rwxrob/cmdbox/samples/name/chinese"
)

func main() {
	cmdbox.Execute() // command implied from binary name
}
