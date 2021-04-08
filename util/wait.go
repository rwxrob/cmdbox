package util

import (
	"context"
	"log"
	"os"
	"os/signal"
)

// WaitForInterrupt just blocks until an interrupt signal is received.
// It should only be called from the main goroutine. It takes a single
// context.CancelFunc that is designed to signal everything to stop
// cleanly before exiting.
func WaitForInterrupt(cancel context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer cancel()
	log.Println("waiting for interrupt")
	<-interrupt
	log.Println("received interrupt")
}
