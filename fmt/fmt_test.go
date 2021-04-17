package fmt

import (
	"bytes"
	_fmt "fmt"
	"os"
	"strings"
	"testing"
)

var fdoh = func() string { return "doh" }
var fv = func() string { return "%v" }

func TestErrorf(t *testing.T) {
	if String(Errorf("%v", "doh")) != "doh" {
		t.Fail()
	}
	if String(Errorf("%v", fdoh)) != "doh" {
		t.Fail()
	}
	if String(Errorf(fdoh)) != "doh" {
		t.Fail()
	}
}

func TestFprint(t *testing.T) {
	buf := bytes.NewBufferString("")
	Fprint(buf, "doh")
	if buf.String() != "doh" {
		t.Fail()
	}
	Fprint(buf, fdoh)
	if buf.String() != "dohdoh" {
		t.Fail()
	}
}

func TestFprintf(t *testing.T) {
	buf := bytes.NewBufferString("")
	Fprintf(buf, "%v", "doh")
	if buf.String() != "doh" {
		t.Fail()
	}
	Fprintf(buf, "%v", fdoh)
	if buf.String() != "dohdoh" {
		t.Fail()
	}
	Fprintf(buf, fv, fdoh)
	if buf.String() != "dohdohdoh" {
		t.Fail()
	}
}

func TestFprintln(t *testing.T) {
	buf := bytes.NewBufferString("")
	Fprintln(buf, "doh")
	if buf.String() != "doh\n" {
		t.Fail()
	}
	Fprintln(buf, fdoh)
	if buf.String() != "doh\ndoh\n" {
		t.Fail()
	}
}

func TestFscan(t *testing.T) {
	buf := strings.NewReader("5 true gophers")
	var i int
	var b bool
	var s string
	Fscan(buf, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestFscanf(t *testing.T) {
	buf := strings.NewReader("5 true gophers")
	var i int
	var b bool
	var s string
	Fscanf(buf, "%d %t %s", &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	buf = strings.NewReader("6 true stringers")
	Fscanf(buf, func() string { return "%d %t %s" }, &i, &b, &s)
	if i != 6 || !b || s != "stringers" {
		t.Fail()
	}
}

func TestFscanln(t *testing.T) {
	buf := strings.NewReader("5 true gophers\n")
	var i int
	var b bool
	var s string
	Fscanln(buf, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func ExamplePrint() {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	_fmt.Print(i, "\n")
	Print(i, "\n")
	_fmt.Print(b, "\n")
	Print(b, "\n")
	_fmt.Print(s, "\n")
	Print(s, "\n")
	_fmt.Print(fl, "\n")
	Print(fl, "\n")
	Print(fdoh, "\n")
	// Output:
	// 42
	// 42
	// true
	// true
	// doh
	// doh
	// 2.4
	// 2.4
	// doh
}

func ExamplePrintln() {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	_fmt.Println(i)
	Println(i)
	_fmt.Println(b)
	Println(b)
	_fmt.Println(s)
	Println(s)
	_fmt.Println(fl)
	Println(fl)
	Println(fdoh)
	// Output:
	// 42
	// 42
	// true
	// true
	// doh
	// doh
	// 2.4
	// 2.4
	// doh
}

func ExamplePrintf() {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	_fmt.Printf("%v%v%v%v\n", i, b, s, fl)
	Printf("%v%v%v%v%v\n", i, b, s, fl, fdoh)
	Printf(func() string { return "%v%v%v%v%v\n" }, i, b, s, fl, fdoh)
	// Output:
	// 42truedoh2.4
	// 42truedoh2.4doh
	// 42truedoh2.4doh
}

func TestScan(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
	}
	_fmt.Fprint(f, "5 true gophers")
	f.Seek(0, 0)
	orig := os.Stdin
	defer func() { os.Stdin = orig }()
	os.Stdin = f
	var i int
	var b bool
	var s string
	Scan(&i, &b, &s)
	t.Logf("%v %v %v\n", i, b, s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestScanf(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
	}
	_fmt.Fprint(f, "5 true gophers")
	f.Seek(0, 0)
	orig := os.Stdin
	defer func() { os.Stdin = orig }()
	os.Stdin = f
	var i int
	var b bool
	var s string
	Scanf("%d %t %s", &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	f.Seek(0, 0)
	Scanf(func() string { return "%d %t %s" }, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestScanln(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
	}
	_fmt.Fprint(f, "5 true gophers\n")
	f.Seek(0, 0)
	orig := os.Stdin
	defer func() { os.Stdin = orig }()
	os.Stdin = f
	var i int
	var b bool
	var s string
	Scanln(&i, &b, &s)
	t.Logf("%v %v %v\n", i, b, s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestSprint(t *testing.T) {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	var want, got string
	want += _fmt.Sprint(i)
	want += _fmt.Sprint(b)
	want += _fmt.Sprint(s)
	want += _fmt.Sprint(fl)
	want += "doh"
	got += Sprint(i)
	got += Sprint(b)
	got += Sprint(s)
	got += Sprint(fl)
	got += Sprint(fdoh)
	if want != got {
		t.Errorf("\nwant: %v\ngot:  %v\n", want, got)
	}
}

func TestSprintf(t *testing.T) {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	want := _fmt.Sprintf("%v%v%v%v%v", i, b, s, fl, "doh")
	got := Sprintf("%v%v%v%v%v", i, b, s, fl, fdoh)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
	got = Sprintf(func() string { return "%v%v%v%v%v" }, i, b, s, fl, fdoh)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestSprintln(t *testing.T) {
	i := 42
	b := true
	s := "doh"
	fl := 2.4
	var want, got string
	want += _fmt.Sprintln(i)
	want += _fmt.Sprintln(b)
	want += _fmt.Sprintln(s)
	want += _fmt.Sprintln(fl)
	want += "doh\n"
	got += Sprintln(i)
	got += Sprintln(b)
	got += Sprintln(s)
	got += Sprintln(fl)
	got += Sprintln(fdoh)
	if want != got {
		t.Errorf("\nwant: %q\ngot:  %q\n", want, got)
	}
}

func TestSscan(t *testing.T) {
	buf := "5 true gophers"
	var i int
	var b bool
	var s string
	Sscan(buf, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	Sscan(func() string { return "5 true gophers" }, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestSscanf(t *testing.T) {
	buf := "5 true gophers"
	var i int
	var b bool
	var s string
	Sscanf(buf, "%d %t %s", &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	Sscanf(func() string { return "5 true gophers" }, "%d %t %s", &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	Sscanf(func() string { return "5 true gophers" },
		func() string { return "%d %t %s" }, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestSscanln(t *testing.T) {
	buf := "5 true gophers\n"
	var i int
	var b bool
	var s string
	Sscanln(buf, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
	Sscanln(func() string { return "5 true gophers" }, &i, &b, &s)
	if i != 5 || !b || s != "gophers" {
		t.Fail()
	}
}

func TestSmartPrintln(t *testing.T) {
	t.Skip("tests for live terminal")
}
