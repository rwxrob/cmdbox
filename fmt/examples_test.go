package fmt_test

import (
	_fmt "fmt"

	"github.com/rwxrob/cmdbox/fmt"
)

type ImmaFuncString struct{}

func (s ImmaFuncString) String() string {
	return "Hello"
}

func ExampleFuncString_String() {

	f := func() string { return "Hello" }
	fmt.Println(f)

	s := "Hello"
	fmt.Println(s)

	st := ImmaFuncString{} // st.String()
	fmt.Println(st)

	// Output:
	// Hello
	// Hello
	// Hello
}

type another struct {
	msg string
}

func NewStructPointer(msg string) *another {
	p := new(another)
	p.msg = msg
	return p
}

func (a *another) String() string {
	return a.msg
}

func ExampleString_string() {
	_fmt.Printf("%q\n", "hello")
	fmt.Printf("%q\n", "hello")
	// Output:
	// "hello"
	// "hello"
}

func ExampleString_int() {
	_fmt.Printf("%q\n", 42)
	fmt.Printf("%q\n", 42)
	// Output:
	// '*'
	// "42"
}

func ExampleString_float() {
	_fmt.Printf("%v\n", 2.4)
	fmt.Printf("%q\n", 2.4)
	// Output:
	// 2.4
	// "2.4"
}

func ExampleString_bool() {
	_fmt.Printf("%v\n", true)
	fmt.Printf("%q\n", true)
	// Output:
	// true
	// "true"
}

func ExampleString_nil() {
	_fmt.Printf("%v\n", nil)
	fmt.Printf("%q\n", nil)
	// Output:
	// <nil>
	// ""
}

func ExampleString_FuncString() {
	f := func() string { return "func soul brotha" }
	_fmt.Printf("%T\n", f)
	fmt.Printf("%v\n", f)
	// Output:
	// func() string
	// func soul brotha
}

func ExampleString_Stringer() {
	p := NewStructPointer("right about now")
	_fmt.Printf("%q\n", p)
	fmt.Printf("%q\n", p)
	// Output:
	// "right about now"
	// "right about now"
}
