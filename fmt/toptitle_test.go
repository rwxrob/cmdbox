package fmt

import "testing"

func TestTopTitle_20(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "left   center  right"
	got = TopTitle(l, c, r, 20)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_21(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "left   center  right"
	got = TopTitle(l, c, r, 21)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_30(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "left        center       right"
	got = TopTitle(l, c, r, 30)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_15(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "leftcenterright"
	got = TopTitle(l, c, r, 15)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_14(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "center   right"
	got = TopTitle(l, c, r, 14)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_10(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "  center  "
	got = TopTitle(l, c, r, 10)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestTopTitle_4(t *testing.T) {
	l := "left"
	c := "center"
	r := "right"
	var want, got string
	want = "cent"
	got = TopTitle(l, c, r, 4)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}
