package fmt

import "testing"

func TestPeekWord(t *testing.T) {
	var buf []rune
	var word string
	buf = []rune(`some thing`)
	word = string(peekWord(buf, 0))
	t.Logf("%q", word)
	if word != "some" {
		t.Fail()
	}
	word = string(peekWord(buf, 5))
	t.Logf("%q", word)
	if word != "thing" {
		t.Fail()
	}
	word = string(peekWord(buf, 4))
	t.Logf("%q", word)
	if word != "" {
		t.Fail()
	}
}

func TestWrap(t *testing.T) {
	buf := "Here's a string that's not long."
	want := "Here's a\nstring\nthat's not\nlong."
	got := Wrap(buf, 10)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestWrap_none(t *testing.T) {
	if Wrap("some thing", 0) != "some thing" {
		t.Fail()
	}
}
