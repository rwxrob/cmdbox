package util

import "strings"

// SplitBarPop returns the last slice item from a string split on bar
// (|) ("h|help" -> "help").
func SplitBarPop(a string) string {
	all := strings.Split(a, "|")
	return all[len(all)-1]
}
