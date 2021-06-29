package cmdbox

func init() { addHelp() }

func addHelp() {
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
		command is added to all CmdBox commands as h|help. The help command
		always returns an error so as to not be confused with more
		legitimate command output.`

	x.Method = func(args []string) error {
		var helpFor *Command

		if len(args) > 0 {
			potential := args[0]

			// "foo help cmd" -> "foo cmd"
			if x.Caller != nil {
				qualified := x.Caller.Name + " " + potential
				helpFor = Get(qualified)
				if helpFor != nil {
					helpFor.PrintHelp()
					return Harmless()
				}
			}

			// "help cmd" -> "cmd"
			helpFor = Get(potential)
			if helpFor != nil {
				helpFor.PrintHelp()
				return Harmless()
			}

		}

		// "foo help" -> "foo"
		if x.Caller != nil {
			helpFor = Get(x.Caller.Name)
			if helpFor != nil {
				helpFor.PrintHelp()
				return Harmless()
			}
		}

		return CallerRequired()
	}

}
