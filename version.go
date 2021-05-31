package cmdbox

import "github.com/rwxrob/cmdbox/fmt"

func init() {
	x := New("version")
	x.Summary = `provide version and legal information`

	x.Method = func(args []string) error {
		fmt.Println("would show version")
		return nil
	}
}
