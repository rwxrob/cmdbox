package fmt

import "bytes"

// LineCount counts the number of line returns (\n) within the string.
func LineCount(buf string) int {
	return bytes.Count([]byte(buf), []byte{'\n'})
}
