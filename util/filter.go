package util

import (
	"bufio"
	"io"
	"strings"
)

// FilterF is a function that takes a string as input and outputs
// another string. See Filter.
//
type FilterF func(in string) (out string)

// Filter implements the basic idea of a UNIX filter which takes in
// a string, []byte, or io.Reader and transforms each line in some way
// resulting in a new output string. Lines are identified by
// combinations of carriage return and line return or just line return
// (as is standard in the Go scanner). Note, however, that carriage
// returns are always discarded in the resulting output string. Errors
// either for a single FilterF call or for any other reason result in
// either blank lines or a blank output string. The second argument is
// first-class function to perform the transformation. See FilterFunc,
// FilterReader.
//
func Filter(i interface{}, f FilterF) string {
	var out, in string
	switch v := i.(type) {
	case io.Reader:
		byt, err := io.ReadAll(v)
		if err != nil {
			return ""
		}
		in = string(byt)
	case string:
		in = v
	case []byte:
		in = string(v)
	}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		out += f(scanner.Text()) + "\n"
	}
	return out
}
