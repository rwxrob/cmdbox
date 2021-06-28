package esc_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/esc"
)

func ExampleBold() {
	fmt.Println(esc.Bold + "Bold" + esc.Reset)
	// Output:
	// [1mBold[0m
}

func ExampleItalic() {
	fmt.Println(esc.Italic + "Italic" + esc.Reset)
	// Output:
	// [3mItalic[0m
}

func ExampleBoldItalic() {
	fmt.Println(esc.BoldItalic + "BoldItalic" + esc.Reset)
	// Output:
	// [1;3mBoldItalic[0m
}
