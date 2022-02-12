package util

import "log"

// Log calls log.Println on each line of the string or []byte array or
// io.Reader passed to it.
//
func Log(out interface{}) {
	Filter(out, func(in string) string { log.Println(in); return in })
}
