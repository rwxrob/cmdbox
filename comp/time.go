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

package comp

import "strings"

// Lower case month names.
var MonthNamesLower = []string{
	"january",
	"february",
	"march",
	"april",
	"may",
	"june",
	"july",
	"august",
	"september",
	"october",
	"november",
	"december",
}

// Initial upper case month names.
var MonthNamesUpper = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// Names of months for Month completion.
var MonthNames = []string{}

func init() {
	MonthNames = append(MonthNames, MonthNamesLower...)
	MonthNames = append(MonthNames, MonthNamesUpper...)
}

// Month fulfills comp.Func by completing comp.Word with the English
// month names. Upper or lower case will be completed. If no Word is
// detected will return all possible MonthNames, both lower and upper
// case. This can be changed by assigning MonthNames to something else.
func Month() []string {
	word := Word()
	if word == "" {
		return MonthNames
	}
	m := []string{}
	for _, name := range MonthNames {
		if strings.HasPrefix(name, word) {
			m = append(m, name)
		}
	}
	return m
}
