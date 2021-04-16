package fmt

import (
	"testing"
)

func TestEmphasize(t *testing.T) {
	want := []string{
		italic + "Italic" + reset,
		bold + "Bold" + reset,
		bolditalic + "BoldItalic" + reset,
	}
	args := []string{"*Italic*", "**Bold**", "***BoldItalic***"}
	for i, arg := range args {
		t.Logf("testing: %v\n", arg)
		got := Emphasize(arg)
		if got != want[i] {
			t.Errorf("\nwant: %v\ngot:  %v\n", want[i], got)
		}
	}
}
