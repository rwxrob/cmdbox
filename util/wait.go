package util

import (
	"context"
	"log"
	"os"
	"os/signal"
)

// WaitForInterrupt blocks until an interrupt signal is received.  It
// should only be called from the main goroutine. It takes a single
// context.CancelFunc that is designed to signal everything to stop
// cleanly before exiting if wanted before the interrupt is received.
// This can be used for timeouts.
//
func WaitForInterrupt(prep context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer prep()
	<-interrupt
}

// WaitForInterruptAndLog is same as WaitForInterrupt but logs "waiting
// to interrupt" and "received interrupt".
//
func WaitForInterruptAndLog(prep context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer prep()
	log.Println("waiting for interrupt")
	<-interrupt
	log.Println("received interrupt")
}
