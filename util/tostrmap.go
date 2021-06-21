package util

// ToStringMap converts a map[string]interface{} into a Map
// (map[string]string). See cmdbox.ToMap as well.
func ToStringMap(m map[string]interface{}) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[k] = v.(string)
	}
	return n
}
