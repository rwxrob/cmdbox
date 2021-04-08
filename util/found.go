package util

import "os"

// Found returns true if the given path was absolutely found to exist on
// the system. A false return value means either the file does not
// exists or it was not able to determine if it exists or not. WARNING:
// do not use this function if a definitive check for the non-existence
// of a file is required since the possible indeterminate error state is
// a possibility. These checks are also not atomic on many file systems
// so avoid this usage for pseudo-semaphore designs and depend on file
// locks.
func Found(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
