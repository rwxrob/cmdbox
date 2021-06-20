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
