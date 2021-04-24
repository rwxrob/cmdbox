package util

import (
	"encoding/json"
	"fmt"
)

// ConvertToJSON converts any object to its JSON string equivalent with
// two spaces of human-readable indenting. If an error is encountered
// while marshalling an ERROR key will be created with the string value
// of the error as its value.
func ConvertToJSON(a interface{}) string {
	byt, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Sprintf("{\"ERROR\": \"%v\"}", err)
	}
	return string(byt)
}
