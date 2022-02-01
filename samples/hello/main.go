package main

import (
	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("hello",
		"en|english", "ch|chinese",
		"ru|russian", "fr|french",
	)
	x.Summary = `print "hello" in different languages`
}

func main() {
	cmdbox.Execute() // detects main command from name of binary
}
