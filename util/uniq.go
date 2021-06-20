/*
Copyright 2021 Robert S. Muhlestein.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
