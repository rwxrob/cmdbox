package util

import (
	"os"
	"os/exec"
	"testing"
)

// go coverage detection is fucked for this sort of stuff, oh well, we
// did the test even if coverage falsely reports 50%
func TestExec(t *testing.T) {
	if os.Getenv("TESTING_EXEC") == "1" {
		err := Exec("go", "version")
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExec")
	cmd.Env = append(os.Environ(), "TESTING_EXEC=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("process exited with %v", err)
	}
}

func TestExec_noargs(t *testing.T) {
	err := Exec()
	if err == nil {
		t.Error("should have failed since no command")
	}
}

func TestRun_noargs(t *testing.T) {
	err := Run()
	if err == nil {
		t.Error("should have failed since no command")
	}
}

func TestRun_nocmd(t *testing.T) {
	err := Run("__inoexist")
	if err == nil {
		t.Error("should have failed since no command")
	}
}

func TestExec_nocmd(t *testing.T) {
	err := Exec("__inoexist")
	if err == nil {
		t.Error("should have failed since no command")
	}
}

func TestRun(t *testing.T) {
	err := Run("go", "version")
	if err != nil {
		t.Error(err)
	}
}
