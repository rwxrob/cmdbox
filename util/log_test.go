package util_test

import "github.com/rwxrob/cmdbox/util"

func ExampleLog() {
	// prints to stderr with time stamps for each line
	util.Log("line one\nline two")
}
