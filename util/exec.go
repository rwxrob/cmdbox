package util

import (
	"os"
	"os/exec"
	"syscall"
)

// Exec (not to be confused with Execute) will check for the existance
// of the first argument as an executable on the system and then execute
// it using syscall.Exec(), which replaces the currently running program
// with the new one in all respects (stdin, stdout, stderr, process ID,
// etc).
//
// Note that although this is exceptionally faster and cleaner than
// calling any of the os/exec variations it may be less compatible with
// different operating systems.
func Exec(args ...string) error {
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	return syscall.Exec(path, args, os.Environ())
}

// ExitExec exits the currently running Go program and hands off memory
// and control to the executable passed as the first in a string of
// arguments along with the arguments to pass along to the called
// executable. This is only supported on systems that support Go's
// syscall.Exec() and underlying execve() system call.
func ExitExec(xnargs ...string) error {
	return syscall.Exec(xnargs[0], xnargs, os.Environ())
}

// Run checks for existance of first argument as an executable on the
// system and then runs it without exiting in a way that is supported
// across different operating systems. The stdin, stdout, and stderr are
// connected directly to that of the calling program. Use more specific
// exec alternatives if intercepting stdout and stderr are desired.
func Run(args ...string) error {
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
