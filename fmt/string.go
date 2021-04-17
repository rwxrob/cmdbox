package fmt

import _fmt "fmt"

// String returns fmt.Sprintf("%v",arg) but also accepts 'func() string'
// as a type returning its result. Consider this anywhere you might use
// a string() cast or Itoa or other string conversion method. This is
// useful for forcing the evaluation of a string at different points
// during runtime to provide dynamic values. For example, command help
// documentation that depends on some dynamic state. In addition,
// passing nil is a special case and returns an empty string.
func String(a interface{}) string {
	switch s := a.(type) {
	case string:
		return s
	case func() string:
		return s()
	case nil:
		return ""
	default:
		return _fmt.Sprintf("%v", s)
	}
}

// Stringify converts all arguments passed into their fmt.Sprintf("%v")
// equivalents but returns them as an array of interface{} in order to
// maintain compatibility with other fmt package functions. See
// FuncString and String().
func Stringify(a ...interface{}) []interface{} {
	nw := make([]interface{}, len(a))
	for i, n := range a {
		nw[i] = String(n)
	}
	return nw
}
