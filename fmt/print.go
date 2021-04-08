package fmt

import (
	_fmt "fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rwxrob/cmdbox/term"
	"github.com/rwxrob/cmdbox/util"
)

// PagedOut sets the default use of a pager application and calling
// PrintPaged() instead of Print() for non-hidden builtin subcommands.
var PagedOut = true

// Print calls Sprint and prints it.
func Print(stuff ...interface{}) {
	_fmt.Print(Sprint(stuff...))
}

// Println calls Sprint and prints it with a line return.
func Println(stuff ...interface{}) {
	_fmt.Println(Sprint(stuff...))
}

// Sprint returns nothing if empty, acts like fmt.Sprint if it has one
// argument, or acts like Sprintf if it has more than one argument.
// Print can also print Stringers. Use Dump instead for debugging.
func Sprint(stuff ...interface{}) string {
	switch {
	case len(stuff) == 1:
		switch s := stuff[0].(type) {
		case string:
			if len(s) > 0 {
				return _fmt.Sprint(s)
			}
		case util.Stringer:
			return _fmt.Sprint(util.String(s))
		}
	case len(stuff) > 1:
		return _fmt.Sprintf(util.String(stuff[0]), stuff[1:]...)
	}
	return ""
}

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

// SmartPrintln calls Println() or Print() based on if IsTerminal()
// returns true or not.
func SmartPrintln(a ...interface{}) {
	if term.IsTerminal() {
		Println(a...)
		return
	}
	Print(a...)
}
