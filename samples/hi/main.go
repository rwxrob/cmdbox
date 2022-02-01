package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("hi",
		"en|english", "ch|chinese",
		"ru|russian", "fr|french",
	)
	x.Summary = `print "hi" in different languages`

	x = cmdbox.Add("english")
	x.Summary = `print "hi" in English`
	x.Method = func(args []string) error {
		fmt.Println("Hi")
		return nil
	}

	x = cmdbox.Add("chinese")
	x.Summary = `print "hi" in Chinese`
	x.Method = func(args []string) error {
		fmt.Println("你好")
		return nil
	}

	x = cmdbox.Add("russian")
	x.Summary = `print "hi" in Russian`
	x.Method = func(args []string) error {
		fmt.Println("Привет")
		return nil
	}

	x = cmdbox.Add("french")
	x.Summary = `print "hi" in French`
	x.Method = func(args []string) error {
		fmt.Println("Salut")
		return nil
	}

}

func main() {
	cmdbox.Execute() // detects main command from name of binary
}
