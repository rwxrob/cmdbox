package fmt

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Fill in form string containing fields wrapped with brackets and
// numbered. Simplier than text and html template alternatives. Good for
// use with snippet libraries. (See github.com/rwxrob/cmdbox-fill for
// a utility that uses these functions.)
func Fill(form string, fields ...string) string {
	for n, v := range fields {
		lk := fmt.Sprintf("{%v}", n+1)
		form = strings.ReplaceAll(form, lk, v)
	}
	return form
}

// FillFrom calls Fill after a io.ReadAll on the reader passed. Returns
// an empty string if any error occurs.
func FillFrom(r io.Reader, fields ...string) string {
	byt, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	form := string(byt)
	for n, v := range fields {
		lk := fmt.Sprintf("{%v}", n+1)
		form = strings.ReplaceAll(form, lk, v)
	}
	return form
}

// FillIn calls Fill after reading all standard input and passing it as
// the form.
func FillIn(fields ...string) string {
	return FillFrom(os.Stdin, fields...)
}
