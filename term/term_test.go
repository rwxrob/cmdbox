package term

import "testing"

func TestLineCount(t *testing.T) {
	txt := `
  some
  thing
  here
  `
	if LineCount(txt) != 4 {
		t.Error("LineCount failed")
	}
}
