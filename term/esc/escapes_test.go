package esc

// NOTE: These tests must be run with -v and verified by human eyes on
// a fully capable terminal for best test validation.

import (
	"testing"
	"time"
)

func Test_ClearScreen(t *testing.T) {
	t.Log(ClearScreen)
	t.Log("Something")
	time.Sleep(2 * time.Second)
	t.Log(CS)
}

func Test_Bold(t *testing.T) {
	t.Log(Bold + "bold" + Reset)
	t.Log(B + "bold" + X)
}

func Test_Italic(t *testing.T) {
	t.Log(Italic + "italic" + Reset)
	t.Log(I + "italic" + X)
}

func Test_BoldItalic(t *testing.T) {
	t.Log(BoldItalic + "bold_italic" + Reset)
	t.Log(BI + "bold_italic" + X)
}
