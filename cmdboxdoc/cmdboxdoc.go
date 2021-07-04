package cmdboxdoc

import "fmt"

// Format calls Format on an internal Formatter using the defaults. For
// more precise control of formatting see NewFormatter.
func Format(markup string) string { return NewFormatter().Format(markup) }

// Print calls Print on an internal Formatter.
func Print(markup string) { fmt.Print(Format(markup)) }

// NewFormatter returns a new internal formatter that fulfills the Formatter interface.
func NewFormatter() Formatter { return newFormatter() }
