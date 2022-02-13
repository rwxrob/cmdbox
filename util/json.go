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

import (
	"encoding/json"
)

// RawJSON converts any object to its compact JSON single-line string
// equivalent  If an error is encountered while marshalling an ERROR
// key will be created with the string value of the error as its value.
func RawJSON(a interface{}) (string, error) {
	byt, err := json.Marshal(a)
	return string(byt), err
}

// JSON converts any object to its JSON (pretty) long-form string
// equivalent. If an error is encountered while marshalling an ERROR
// key will be created with the string value of the error as its value.
func JSON(a interface{}) (string, error) {
	byt, err := json.MarshalIndent(a, "  ", "  ")
	return string(byt), err
}

// MustJSON calls JSON and logs any error with log.Printf returning an
// empty string if an error occurred.
func MustJSON(a interface{}) string {
	out, err := JSON(a)
	if err != nil {
		Log(err)
	}
	return out
}

// MustRawJSON calls RawJSON and logs any error with log.Printf
// returning an empty string if an error occurred.
func MustRawJSON(a interface{}) string {
	out, err := RawJSON(a)
	if err != nil {
		Log(err)
	}
	return out
}
