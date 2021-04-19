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

// Remove allows an registered Command to be removed. If the name does
// not exist it is ignored. Note that if another registered name points
// to the same Command pointer that it will not be removed.
func Remove(name string) {
	delete(commands, name)
}

// Rename adds a commands registry entry for the new name and deletes
// the entry with the old name. This is useful when a name conflict
// causes Add() to add an underscore (_) to the duplicate and can be
// called from init() at any point after the duplicate has been added to
// resolve the conflict.
func Rename(from, to string) {
	commands[to] = commands[from]
	delete(commands, from)
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
	m := make(map[string]Command, len(commands))
	for k, v := range commands {
		if k[0] == '_' {
			m[k] = *v
		}
	}
	return m
}

// Duplicates returns map of duplicate commands who's names end with
// underscore (_).
func Duplicates() map[string]Command {
	m := make(map[string]Command, len(commands))
	for k, v := range commands {
		if k[len(k)-1] == '_' {
			m[k] = *v
		}
	}
	return m
}

// Has looks for the named command in the internal commands table.
func Has(name string) bool {
	_, has := commands[name]
	return has
}
