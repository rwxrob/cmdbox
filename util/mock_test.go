package util

import (
	"bufio"
	"os"
	"testing"
	"time"
)

func TestMockStdin(t *testing.T) {
	MockStdin("Here is some\nstuff.\n")
	defer UnmockStdin()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() != "Here is some" {
		t.Fail()
	}
	scanner.Scan()
	if scanner.Text() != "stuff." {
		t.Fail()
	}
}

func TestMockStdin_Locked(t *testing.T) {
	MockStdin("Here is some\nstuff.\n")
	go func() { t.Log("Sleeping 2 seconds."); time.Sleep(2 * time.Second); mockin.Unlock() }()
	MockStdin("Here is some\nstuff.\n")
	t.Log("Finally.")
}
