package util

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
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
