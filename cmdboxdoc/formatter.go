package cmdboxdoc

// Formatter provides a common interface that CmdBoxDoc formatter
// authors can choose to implement to maintain extensibility. See the
// private formatter struct for an implementation example.
type Formatter interface {

	// Set the maximum number of columns before word wrap will kick in. If
	// not greater than zero then do not wrap.
	SetCols(cols int)

	// Set the number of columns to indent.
	SetIndent(cols int)

	// Mutators to set the beginning and ending escape string for each
	// detected markup token.
	SetEmphEsc(beg, end string)       // *emph*
	SetStrongEsc(beg, end string)     // **strong**
	SetStrongEmphEsc(beg, end string) // ***strongemph**
	SetIdentEsc(beg, end string)      // <ident>

	// Escape accessors
	EmphEsc() (string, string)
	StrongEsc() (string, string)
	StrongEmphEsc() (string, string)
	IdentEsc() (string, string)

	// Detects the terminal escapes to use by inspecting the system for
	// common terminal configurations including (but, not limited to)
	// ~/.config/cmdbox/config.yaml and LESS_TERMCAP_*). This includes
	// detection of an interactive terminal and setting the emphasis
	// values accordingly, off by default unless Always is true..
	Detect()

	// Format processes a CmdBoxDoc string and returns one that has been
	// formatted with terminal escapes and wrapping.
	Format(markup string) string

	// Print simply calls Format and prints to standard output.
	Print(markup string)
}

type formatter struct {
}

func newFormatter() *formatter {
	f := new(formatter)
	return f
}

func (f *formatter) SetCols(cols int) {
	// TODO
}

func (f *formatter) SetIndent(cols int) {
	// TODO
}

func (f *formatter) SetEmphEsc(beg, end string) {
	// TODO
}

func (f *formatter) SetStrongEsc(beg, end string) {
	// TODO
}

func (f *formatter) SetStrongEmphEsc(beg, end string) {
	// TODO
}

func (f *formatter) SetIdentEsc(beg, end string) {
	// TODO
}

func (f *formatter) EmphEsc() (string, string) {
	// TODO
	return "", ""
}

func (f *formatter) StrongEsc() (string, string) {
	// TODO
	return "", ""
}

func (f *formatter) StrongEmphEsc() (string, string) {
	// TODO
	return "", ""
}

func (f *formatter) IdentEsc() (string, string) {
	// TODO
	return "", ""
}

func (f *formatter) Detect() {
	// TODO
}

func (f *formatter) Format(markup string) string {
	// TODO
	return ""
}

func (f *formatter) Print(markup string) {
	// TODO
}
