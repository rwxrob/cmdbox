package cmdbox

func init() {
	x := Add("help")
	x.Usage = `[<name>]`
	x.Summary = `display command help information`
	x.Version = `v1.0.0`
	x.Copyright = `Copyright 2021 Robert S Muhlestein`
	x.License = `Apache-2`
	x.Source = `https://github.com/rwxrob/cmdbox-help`
	x.Issues = `https://github.com/rwxrob/cmdbox-help/issues`

	x.Description = `
		Use this command to display full help information about the CmdBox
		program or any of its subcommands. If an optional command <name> is
		provided the help information for that specific command will be
		provided instead. If no argument is passed will display help
		information for the immediately preceeding command. By default this
		command is added to all CmdBox commands as h|help.`

	x.Method = func(args []string) error {

		// foo help
		_x := x.Caller
		if _x == nil {
			_x = Main
		}

		// foo help stuff
		if len(args) > 0 {
			_x = Get(args[0])
			if _x == nil {
				switch {
				case x.Caller != nil:
					_x = Get(x.Caller.Name + " " + args[0])
				default:
					_x = Get(Main.Name + " " + args[0])
				}
			}
		}

		_x.PrintHelp()
		return nil
	}

}
