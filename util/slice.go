package util

// OmitFromSlice returns a new array created from the original with the
// each element of the omit array omitted.
//
func OmitFromSlice(orig, omit []string) []string {
	if omit == nil {
		return orig
	}
	out := []string{}
OUTER:
	for _, o := range orig {
		for _, h := range omit {
			if h == o {
				continue OUTER
			}
		}
		out = append(out, o)
	}
	return out
}

// InSlice returns true if this is in that slice of strings.
func InSlice(this string, that []string) bool {
	if that == nil {
		return false
	}
	for _, i := range that {
		if i == this {
			return true
		}
	}
	return false
}
