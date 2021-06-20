/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package term

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// see ../cmd/isterm for an example of this test that is
// much easier to understand

// also coverage will never catch this test

func TestIsTerminal_false(t *testing.T) {
	if os.Getenv("TEST_ISNOTTERM") == "1" {
		fmt.Println("out")
		if !IsTerminal() {
			os.Exit(20)
		}
		os.Exit(1)
	}
	exe := os.Args[0]
	cmd := exec.Command(exe, "-test.run=TestIsTerminal_false")
	cmd.Env = append(os.Environ(), "TEST_ISNOTTERM=1")
	cmd.StdoutPipe() // just enough to push into background
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok {
		t.Log(e.ExitCode())
		if e.ExitCode() != 20 {
			t.Errorf("exit %v: still a terminal", e.ExitCode())
		}
	}
}

func TestIsTerminal_true(t *testing.T) {
	if !IsTerminal() {
		t.Error("terminal not connected")
	}
}
