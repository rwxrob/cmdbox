package util

import "testing"

func TestFound(t *testing.T) {
	if !Found("found.go") {
		t.Error("found.go")
	}
	if Found("__inoexist") {
		t.Fail()
	}
}
