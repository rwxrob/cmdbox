package main

import (
	"github.com/rwxrob/cmdbox"
	_ "github.com/rwxrob/cmdbox/samples/name/chinese"
	_ "github.com/rwxrob/cmdbox/samples/name/russian"
)

func init() {
	x := cmdbox.Add("name",
		"en|english", "ch|chinese",
		"ru|russian", "fr|french",
	)
	x.Summary = `print "My name is ..." in different languages`
}

func main() {
	cmdbox.Execute() // detects main command from name of binary
}
