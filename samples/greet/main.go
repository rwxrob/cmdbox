package main

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("greet",
		"en|english", "ch|chinese",
		"ru|russian", "fr|french",
	)
	x.Summary = `print a greeting in different languages`

	x = cmdbox.Add("english")
	x.Summary = `print hello in English`
	x.Method = func(args []string) error {
		fmt.Println("Hello")
		return nil
	}

	x = cmdbox.Add("chinese")
	x.Summary = `print hello in Chinese`
	x.Method = func(args []string) error {
		fmt.Println("你好")
		return nil
	}

	x = cmdbox.Add("russian")
	x.Summary = `print hello in Russian`
	x.Method = func(args []string) error {
		fmt.Println("Здравствуйте")
		return nil
	}

	x = cmdbox.Add("french")
	x.Summary = `print hello in French`
	x.Method = func(args []string) error {
		fmt.Println("Bonjour")
		return nil
	}

}

func main() {
	cmdbox.Execute()
}
