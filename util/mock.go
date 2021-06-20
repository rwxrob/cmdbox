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
	"fmt"
	"os"
	"sync"
)

var mockin = struct {
	sync.Mutex
	f    *os.File
	orig *os.File
}{}

// MockStdin replaces os.Stdin (temporarily) with the content of the
// buffer passed to it by swapping it out internally and locking any
// further calls to MockStdin until UnmockStdin() is called. Note that
// anything that uses os.Stdin in any way, concurrently or otherwise,
// will share the same altered os.Stdin from that point until
// UnmockStdin() is called.
//
//     MockStdin(mystr)
//     defer UnmockStdin()
//
// Panics if cannot os.CreateTemp for any reason and therfore should
// only really be used for testing.
//
func MockStdin(buf string) {
	mockin.Lock()
	f, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(f, buf)
	f.Seek(0, 0)
	mockin.orig = os.Stdin
	mockin.f = f
	os.Stdin = f
}

// UnmockStdin restores os.Stdin to its original value, cleans up the
// buffer file, and removes the internal local preventing conflicting
// MockStdin calls.
//
func UnmockStdin() {
	os.Stdin = mockin.orig
	os.Remove(mockin.f.Name())
	mockin.f = nil
	mockin.orig = nil
	mockin.Unlock()
}
