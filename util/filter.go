package util

import (
	"bufio"
	"fmt"
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
	switch v := i.(type) {
	case io.Reader:
		byt, err := io.ReadAll(v)
		if err != nil {
			return ""
		}
		return filterString(string(byt), f)
	case string:
		return filterString(v, f)
	case []byte:
		return filterString(string(v), f)

		// FIXME convert this shit to generics as soon as possible
	case []interface{}, []string, [][]byte, [][]rune, []bool,
		[]int, []int32, []int64, []uint32, []uint, []uint64,
		[]float32, []float64:
		return filterSlice(v, f)

	default:
		return filterString(fmt.Sprintf("%v", i), f)
	}
}

func filterString(s string, f FilterF) string {
	out := ""
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		out += f(scanner.Text()) + "\n"
	}
	return out
}

// this redundancy thing will be completely unnecessary with generics
func filterSlice(s interface{}, f FilterF) string {
	out := ""
	switch v := s.(type) {
	case []interface{}:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []string:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case [][]byte:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case [][]rune:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []bool:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []int:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []int32:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []int64:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []uint32:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []uint:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []uint64:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []float32:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	case []float64:
		for _, i := range v {
			out += f(fmt.Sprintf("%v", i)) + "\n"
		}
	}
	return out
}
