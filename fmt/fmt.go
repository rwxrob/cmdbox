package fmt

import (
	_fmt "fmt"
	"io"

	"github.com/rwxrob/cmdbox/term"
)

// ------------------------------- types ------------------------------

// GoStringer is same as fmt.GoStringer.
type GoStringer interface {
	String() string
}

// ScanState is same as fmt.ScanState.
type ScanState interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	SkipSpace()
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	Width() (wid int, ok bool)
	Read(buf []byte) (n int, err error)
}

// Scanner is same as fmt.Scanner.
type Scanner interface {
	Scan(state ScanState, verb rune) error
}

// State is same as fmt.State.
type State interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}

// Stringer is same as fmt.Stringer.
type Stringer interface {
	String() string
}

// ----------------------------- functions ----------------------------

// Errorf with Stringify().
func Errorf(format interface{}, a ...interface{}) error {
	return _fmt.Errorf(String(format), Stringify(a...)...)
}

// Fprint with Stringify().
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return _fmt.Fprint(w, Stringify(a...)...)
}

// Fprintf with Stringify().
func Fprintf(w io.Writer, format interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Fprintf(w, String(format), Stringify(a...)...)
}

// Fprintln with Stringify().
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return _fmt.Fprintln(w, Stringify(a...)...)
}

// Fscan delegates directly to fmt.Fscan().
func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	return _fmt.Fscan(r, a...)
}

// Fscanf with any interface{} for format.
func Fscanf(r io.Reader, format interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Fscanf(r, String(format), a...)
}

// Fscanln delegates directly to fmt.Fscanln().
func Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
	return _fmt.Fscanln(r, a...)
}

// Print with Stringify().
func Print(a ...interface{}) (n int, err error) {
	return _fmt.Print(Stringify(a...)...)
}

// Printf with Stringify().
func Printf(format interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Printf(String(format), Stringify(a...)...)
}

// Println with Stringify().
func Println(a ...interface{}) (n int, err error) {
	return _fmt.Println(Stringify(a...)...)
}

// Scan just delegates to fmt.
func Scan(a ...interface{}) (n int, err error) {
	return _fmt.Scan(a...)
}

// Scanf with Stringify() for the format.
func Scanf(format interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Scanf(String(format), a...)
}

// Scanln just delegates to fmt.
func Scanln(a ...interface{}) (n int, err error) {
	return _fmt.Scanln(a...)
}

// Sprint with Stringify().
func Sprint(a ...interface{}) string {
	return _fmt.Sprint(Stringify(a...)...)
}

// Sprintf with Stringify().
func Sprintf(format interface{}, a ...interface{}) string {
	return _fmt.Sprintf(String(format), Stringify(a...)...)
}

// Sprintf with Stringify().
func Sprintln(a ...interface{}) string {
	return _fmt.Sprintln(Stringify(a...)...)
}

// Sscan with Stringify().
func Sscan(str interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Sscan(String(str), a...)
}

// Sscanf with Stringify().
func Sscanf(str interface{}, format interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Sscanf(String(str), String(format), a...)
}

// Sscanln with Stringify().
func Sscanln(str interface{}, a ...interface{}) (n int, err error) {
	return _fmt.Sscanln(String(str), a...)
}

// ------------------ end fmt compatibility functions -----------------

// SmartPrintln calls Println() or Print() based on if IsTerminal()
// returns true or not.
func SmartPrintln(a ...interface{}) {
	if term.IsTerminal() {
		Println(a...)
		return
	}
	Print(a...)
}
