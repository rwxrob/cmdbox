package cmdbox

import "github.com/rwxrob/cmdbox/fmt"

func init() {
	x := New("help")
	x.Usage = `<subcmd> ...`
	x.Summary = `provide help output similar to man page`

	x.Method = func(args []string) error {
		fmt.Println("would show help")
		return nil
	}
}
