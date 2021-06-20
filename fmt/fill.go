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
