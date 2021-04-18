package fmt

import (
	"os"
	"os/exec"
	"strings"

	"github.com/rwxrob/cmdbox/term"
)

// PagedOut sets the default use of a pager application and calling
// PrintPaged() instead of Print() for non-hidden builtin subcommands.
var PagedOut = true

// PagedDefStatus is the status line passed to `less` to provide information at
// the bottom of the screen prompting the user what to do. Helpful with
// implementing help in languages besides English.
var PagedDefStatus = `Line %lb [<space>(down), b(ack), h(elp), q(quit)]`

// PrintPaged prints a string to the system pager (usually less) using
// the second argument as the custom status string (usually at the
// bottom). Control is returned to the calling program after completion.
// If no pager application is detected the regular Print() will be
// called intead. If status string is empty PagedDefStatus will be used
// (use " " to empty). Currently only the less pager is supported.
func PrintPaged(buf, status string) {
	if status == "" {
		status = PagedDefStatus
	}
	_, err := exec.LookPath("less")
	if err != nil || term.LineCount(buf) < int(term.WinSize.Row) {
		Print(buf)
		return
	}
	less := exec.Command("less", "-r", "-Ps"+status)
	less.Stdin = strings.NewReader(buf)
	less.Stdout = os.Stdout
	less.Run()
}
