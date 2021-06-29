package cmdbox

import "fmt"

func init() { addVersion() }

func addVersion() {
	x := Add("version")
	x.Usage = `[<subcmd>]`
	x.Summary = `provide version and legal information`
	x.Version = `v1.0.0`
	x.Copyright = `Copyright 2021 Robert S Muhlestein`
	x.License = `Apache-2`

	x.Description = `
		Use this command to print version and legal information. If an
		optional <subcmd> is provided the version information for that
		specific subcommand (if any) will be provided instead (since this
		CmdBox command may have been composed from one or more other
		independent command modules before being statically linked).`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			fmt.Println(x.Legal())
			return nil
		}
		subcmd := Get(args[0])
		if subcmd == nil {
			return Unimplemented(args[0])
		}
		fmt.Println(subcmd.Legal())
		return nil
	}

}
