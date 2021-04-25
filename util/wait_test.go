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
