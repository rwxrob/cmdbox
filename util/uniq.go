package util

import "sort"

// UniqStrMapVals returns a slice of sorted unique map string values. (Perhaps
// a candidate for generics one day.)
func UniqStrMapVals(m map[string]string) []string {
	seen := map[string]bool{}
	uniq := []string{}
	for _, v := range m {
		if seen[v] {
			continue
		}
		seen[v] = true
		uniq = append(uniq, v)
	}
	sort.Strings(uniq)
	return uniq
}
