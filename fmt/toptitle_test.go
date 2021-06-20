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
