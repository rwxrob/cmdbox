package cmdbox

// commands contains the commands that are loaded at init() by calling
// cmdbox.New(). One of them is the set to cmdbox.MainCommand.
var commands = map[string]*Command{}

// Commands returns a copy of the commands that are loaded at init() by calling
// cmdbox.New(). One of them is set to cmdbox.MainCommand. Also see
// VisibleCommands() and HiddenCommands().
func Commands() map[string]Command {
	copy := make(map[string]Command, len(commands))
	for k, v := range commands {
		copy[k] = *v
	}
	return copy
}

// VisibleCommands returns a copy of the visible commands not beginning
// with underscore ('_'). See Commands()
func VisibleCommands() map[string]Command {
	copy := make(map[string]Command, len(commands))
	for k, v := range commands {
		if k[0] != '_' {
			copy[k] = *v
		}
	}
	return copy
}

// HiddenCommands returns a copy of the hidden commands beginning
// with underscore ('_'). See Commands()
func HiddenCommands() map[string]Command {
	copy := make(map[string]Command, len(commands))
	for k, v := range commands {
		if k[0] == '_' {
			copy[k] = *v
		}
	}
	return copy
}

// Has looks for the named command in the internal commands table.
func Has(name string) bool {
	_, has := commands[name]
	return has
}
