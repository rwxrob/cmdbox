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

package util

import (
	"os"
	"testing"
	"time"
)

func TestWaitForInterrupt_interrupted(t *testing.T) {
	t.Log("Waiting...")
	go func() {
		time.Sleep(2 * time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(os.Interrupt)
	}()
	WaitForInterrupt(func() {})
	t.Log("Interrupted")
}

func TestWaitForInterruptAndLog_interrupted(t *testing.T) {
	go func() {
		time.Sleep(2 * time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(os.Interrupt)
	}()
	WaitForInterruptAndLog(func() {})
}

func TestWaitForInterrupt_prep(t *testing.T) {
	t.Log("Waiting...")
	prep := func() { t.Log("would prepare for interrupt") }
	go func() {
		time.Sleep(2 * time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(os.Interrupt)
	}()
	WaitForInterrupt(prep)
	t.Log("Interrupted")
}

func TestWaitForInterruptAndLog_prep(t *testing.T) {
	prep := func() { t.Log("would prepare for interrupt") }
	go func() {
		time.Sleep(2 * time.Second)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(os.Interrupt)
	}()
	WaitForInterruptAndLog(prep)
}
