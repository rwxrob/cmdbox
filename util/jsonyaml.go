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
	"fmt"

	"gopkg.in/yaml.v2"
)

// ToJSON converts any object to its compact JSON string
// equivalent.  If an error is encountered while marshalling an ERROR
// key will be created with the string value of the error as its value.
func ToJSON(a interface{}) string {
	byt, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\": \"%v\"}", err)
	}
	return string(byt)
}

// PrintJSON prints any object as its compact JSON string equivalent.
// See ToJSON.
func PrintJSON(a interface{}) { fmt.Println(ToJSON(a)) }

// ToYAML converts any object to its compact YAML string equivalent.  If
// an error is encountered while marshalling an ERROR key will be
// created with the string value of the error as its value.
func ToYAML(a interface{}) string {
	byt, err := yaml.Marshal(a)
	if err != nil {
		return fmt.Sprintf("ERROR: \"%v\"", err)
	}
	return string(byt)
}

// PrintYAML prints any object as its YAML string equivalent. Commonly
// used for testing, debugging, and visualizing complex structures. See
// ToYAML.
func PrintYAML(a interface{}) { fmt.Print(ToYAML(a)) }
