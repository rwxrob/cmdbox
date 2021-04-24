package util

import (
	"testing"

	"github.com/rwxrob/cmdbox/fmt"
)

func TestArgsFromStdin(t *testing.T) {
	MockStdin("one\ntwo\nthree")
	defer UnmockStdin()
	args := ArgsFromStdin()
	if fmt.String(args) != "[one two three]" {
		t.Error(args)
	}
}

func TestArgsFromStdin_Positional(t *testing.T) {
	MockStdin("{1} and {2}\n{2} and {1}")
	defer UnmockStdin()
	args := ArgsFromStdin("one", "two")
	if fmt.String(args) != "[one and two two and one]" {
		t.Error(args)
	}
}
